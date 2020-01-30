package cli

import (
	"fmt"
	"github.com/dairlair/sentimentd/pkg/interface/cli/util"
	stan "github.com/nats-io/go-nats-streaming"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// QueueCreator describes dependency which is used to run-time getting connection to NATS Streaming
type QueueCreator func() (stan.Conn, string, string)

type producer interface {

}

type processor func (input string) (output string, err error)

type consumer func ([]byte)

// NewCmdListen returns command for a queue listening.
// @TODO This command should be completely refactored.
func (runner *CommandsRunner) NewCmdListen(queueCreator QueueCreator) *cobra.Command {
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

			queue, source, target := queueCreator()

			var pr processor = func (input string) (output string, err error) {
				return analyse(runner, brain.GetID(), input)
			}

			readFromReaderAndWriteToWrite(queue, source, target, pr)
		},
	}
}

func readFromReaderAndWriteToWrite(conn stan.Conn, source string, target string, pr processor) {
	var cb consumer = func (data []byte) {
		err := conn.Publish(target, data)
		if err != nil {
			log.Errorf("Prediction results publish failed: %s", err)
		}
	}

	subscription, err := conn.Subscribe(source, func(msg *stan.Msg) {
		processJSONAndPushBack(cb, string(msg.Data), func(text string) (string, error) {
			return pr(text)
		})
	})
	if err != nil {
		log.Fatalf("Can not subscribe to channel [%s]: %s", source, err)
	}

	util.WaitInterruption(func() {
		_ = subscription.Close()
	})
}

func processJSONAndPushBack(cb consumer, json string, analyser func(text string) (string, error)) {
	text := gjson.Get(json, "fullText")
	fmt.Printf("Message retrieved: %s\n\n", text.Str)

	prediction, err := analyser(text.Str)
	if err != nil {
		log.Errorf("Prediction failed: %s", err)
		return
	}

	data, err := sjson.Set(json, "prediction", prediction)
	if err != nil {
		log.Errorf("JSON Modification failed: %s", err)
		return
	}

	cb([]byte(data))
}

func analyse(runner *CommandsRunner, brainID int64, text string) (json string, err error) {
	prediction, err := runner.app.HumanizedPredict(brainID, text)
	if err != nil {
		return "", err
	}

	return prediction.JSON()
}