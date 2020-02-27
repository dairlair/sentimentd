package cli

import (
	"github.com/dairlair/sentimentd/pkg/interface/cli/util"
	stan "github.com/nats-io/go-nats-streaming"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// QueueCreator describes dependency which is used to run-time getting connection to NATS Streaming
type QueueCreator func() (stan.Conn, string, string)

type producer interface {
}

type processor func(input string) []Prediction

type consumer func([]byte)

// NewCmdListen returns command for a queue listening.
// @TODO This command should be completely refactored.
func (runner *CommandsRunner) NewCmdListen(queueCreator QueueCreator) *cobra.Command {
	return &cobra.Command{
		Use:   "listen",
		Short: "Listen a configured queue and push analysed message to queue",
		Run: func(cmd *cobra.Command, args []string) {
			var brainReferences []string
			if len(args) == 0 {
				// Try to gen brain references from config
				brainReferences = viper.GetStringSlice("listen.brains")
			} else {
				brainReferences = args
			}

			log.Infof("Listen brains: %v", brainReferences)
			brainsMap := getBrainsMap(runner, brainReferences)
			if len(brainsMap) == 0 {
				log.Fatalf("No brains found, exit...")
			}

			queue, source, target := queueCreator()

			var pr processor = func(input string) []Prediction {
				return analyse(runner, brainsMap, input)
			}

			readFromReaderAndWriteToWrite(queue, source, target, pr)
		},
	}
}

func readFromReaderAndWriteToWrite(conn stan.Conn, source string, target string, pr processor) {
	var cb consumer = func(data []byte) {
		err := conn.Publish(target, data)
		if err != nil {
			log.Errorf("Prediction results publish failed: %s", err)
		}
		log.Infof("Message published: %s", data)
	}

	subscription, err := conn.Subscribe(source, func(msg *stan.Msg) {
		processJSONAndPushBack(cb, string(msg.Data), func(text string) []Prediction {
			return pr(text)
		})
		msg.Ack()
	})
	if err != nil {
		log.Fatalf("Can not subscribe to channel [%s]: %s", source, err)
	}

	util.WaitInterruption(func() {
		_ = subscription.Close()
	})
}

func processJSONAndPushBack(cb consumer, json string, analyser func(text string) []Prediction) {
	text := gjson.Get(json, "originText")
	sentiment := analyser(text.Str)
	data, err := sjson.Set(json, "sentiment", sentiment)
	if err != nil {
		log.Errorf("JSON Modification failed: %s", err)
		return
	}
	cb([]byte(data))
}

func analyse(runner *CommandsRunner, brainsMap map[int64]string, text string) []Prediction {
	predictions := make([]Prediction, 0)
	for brainID, brainReference := range brainsMap {
		prediction, err := runner.app.HumanizedPredict(brainID, text)
		if err != nil {
			log.Errorf("Prediction failed. %s", err)
			continue
		}
		for class, probability := range prediction.Probabilities {
			predictions = append(predictions, Prediction{
				Brain:       brainReference,
				Class:       class,
				Probability: probability,
			})
		}
	}
	return predictions
}

func getBrainsMap(runner *CommandsRunner, references []string) map[int64]string {
	brainsMap := make(map[int64]string)
	for _, reference := range references {
		brain, err := runner.app.GetBrainByReference(reference)
		if err != nil {
			log.Error(err)
		} else {
			brainsMap[brain.GetID()] = brain.GetName()
		}

	}
	return brainsMap
}

type Prediction struct {
	Brain string `json:"brain"`
	Class string `json:"class"`
	Probability float64 `json:"probability"`
}