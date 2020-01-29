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
	"time"
)

// QueueCreator describes dependency which is used to run-time getting connection to NATS Streaming
type QueueCreator func() stan.Conn

// NewCmdListen returns command for a queue listening.
// The command accepts two parameters which are providers of input and output data sources for the command.
// Actually command does not know anything about providers realisation.
func (runner *CommandsRunner) NewCmdListen(queueCreator QueueCreator) *cobra.Command {
	return &cobra.Command{
		Use:   "listen",
		Short: "Listen a configured queue and push analysed message to queue",
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			brain, err := runner.app.GetBrainByReference(args[0])
			if err != nil {
				runner.Err(err)
				return
			}

			sourceChannel := args[1]
			targetChannel := args[2]

			queue := queueCreator()
			readFromReaderAndWriteToWrite(runner, queue, brain.GetID(), sourceChannel, targetChannel)
		},
	}
}

func readFromReaderAndWriteToWrite(runner *CommandsRunner, conn stan.Conn, brainID int64, source string, target string) {
	subscription, err := conn.Subscribe(source, func(msg *stan.Msg) {
		processJSONAndPushBack(conn, target, string(msg.Data), func(text string) (string, error) {
			return analyse(runner, brainID, text)
		})
	})
	if err != nil {
		log.Fatalf("Can not subscribe to channel [%s]: %s", source, err)
	}

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			runner.Out("Received an interrupt, unsubscribing and closing connection...")
			_ = subscription.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}

func processJSONAndPushBack(conn stan.Conn, target string, json string, analyser func(text string) (string, error)) {
	text := gjson.Get(json, "fullText")
	fmt.Printf("Message retrieved: %s\n\n", text.Str)
	t := time.Now()
	prediction, err := analyser(text.Str)
	d := time.Since(t)
	log.Infof("Message analysed for %d μs", d.Nanoseconds()/1000)
	if err != nil {
		log.Errorf("Prediction failed: %s", err)
		return
	}

	data, err := sjson.Set(json, "prediction.probabilities", prediction)
	if err != nil {
		log.Errorf("JSON Modification failed: %s", err)
		return
	}

	data, err = sjson.Set(data, "prediction.duration", d.Seconds())
	if err != nil {
		log.Errorf("JSON Modification failed: %s", err)
		return
	}

	err = conn.Publish("TweetAnalysed", []byte(data))
	if err != nil {
		log.Errorf("Prediction results publish failed: %s", err)
		return
	}
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
