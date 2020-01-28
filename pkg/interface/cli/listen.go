package cli

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/dairlair/sentimentd/pkg/domain/entity"
	stan "github.com/nats-io/go-nats-streaming"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

// StringReader defines dependency which is used to read input messages
type StringReader interface {
	ReadString() string
}

// StringWriter defines dependency which is used to write output messages
type StringWriter interface {
	WriteString(string)
}

type ReaderProvider func() StringReader
type WriterProvider func() StringWriter

// NewCmdListen returns command for a queue listening.
// The command accepts two parameters which are providers of input and output datasources for the command.
// Actually command does not know anything about providers realisation.
func (runner *CommandsRunner) NewCmdListen(rProvider ReaderProvider, wProvider WriterProvider) *cobra.Command {
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

			readFromReaderAndWriteToWrite(rProvider(), wProvider())

			if true {
				return
			}

			listenFromQueue(runner, brain)
		},
	}
}

func readFromReaderAndWriteToWrite(r StringReader, w StringWriter) {
	for {
		msg := r.ReadString()
		fmt.Printf("Message read from reader: %s\n", msg)
	}
}

func listenFromQueue(runner *CommandsRunner, brain entity.BrainInterface) {
	natsSrvAddr := "nats://127.0.0.1:4222"
	natsClient, err := stan.Connect("test-cluster", "sentimentd", func(options *stan.Options) error {
		options.NatsURL = natsSrvAddr
		return nil
	})
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, natsSrvAddr)
	}
	defer natsClient.Close()
	fmt.Println("Connected to NATS Streaming server")
	natsClient.Subscribe("TweetWatch.TweetSaved", func(msg *stan.Msg) {
		//fmt.Printf("JSON received: %s\n", msg.Data)
		text := gjson.Get(string(msg.Data), "fullText")
		fmt.Printf("Message retrieved: %s\n", text.Str)
		_, err := runner.app.Predict(brain.GetID(), text.Str)
		if err != nil {
			runner.Err(err)
		}
	})

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
