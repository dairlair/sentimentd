package cli

import (
	"fmt"
	stan "github.com/nats-io/go-nats-streaming"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"os"
	"os/signal"
)

type QueueCreator func () stan.Conn

// NewCmdListen returns command for a queue listening.
// The command accepts two parameters which are providers of input and output data sources for the command.
// Actually command does not know anything about providers realisation.
func (runner *CommandsRunner) NewCmdListen(creator QueueCreator) *cobra.Command {
	return &cobra.Command{
		Use:   "listen",
		Short: "Listen a configured queue and push analysed message to queue",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			brain, err := runner.app.GetBrainByReference(args[0])
			if err != nil {
				runner.Err(err)
				return
			}

			readFromReaderAndWriteToWrite(runner, creator(), brain.GetID())
		},
	}
}

func readFromReaderAndWriteToWrite(runner *CommandsRunner, conn stan.Conn, brainId int64) {
	s, err := conn.Subscribe("TweetSaved", func(msg *stan.Msg) {
		json := string(msg.Data)
		text := gjson.Get(json, "fullText")
		fmt.Printf("Message retrieved: %s\n\n", text.Str)
		prediction, err := analyse(runner, brainId, text.Str)
		if err != nil {
			log.Errorf("Prediction failed: %s", err)
			return
		}

		data, err := sjson.Set(json, "prediction", prediction)
		if err != nil {
			log.Errorf("JSON Modification failed: %s", err)
			return
		}

		err = conn.Publish("TweetAnalysed", []byte(data))
		if err != nil {
			log.Errorf("Prediction results publish failed: %s", err)
			return
		}
	})
	if err != nil {
		log.Fatalf("Can not subscribe: %s", err)
	}

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			_ = s.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}

func analyse(runner *CommandsRunner, brainID int64, text string) (json string, err error) {
	prediction, err := runner.app.Predict(brainID, text)
	if err != nil {
		return "", err
	}

	json = "{}"
	for _, classID := range prediction.GetClassIDs() {
		class, _ := runner.app.GetClassByID(classID)
		probability := prediction.GetClassProbability(classID)
		json, err = sjson.Set(json, class.GetLabel(), probability)
		if err != nil {

			return "", err
		}
	}

	return json, nil
}