package cli

import (
	"fmt"
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	stan "github.com/nats-io/go-nats-streaming"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"os"
	"os/signal"
	"github.com/cheggaaa/pb/v3"
)

func (runner *CommandsRunner) NewCmdListen() *cobra.Command {
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

			runner.Out(fmt.Sprintf("Brain for analyse: %s", brain.GetName()))
			listenFromQueue(runner, brain)
		},
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
	bar := pb.StartNew(10000)
	natsClient.Subscribe("TweetWatch.TweetSaved", func(msg *stan.Msg) {
		//fmt.Printf("JSON received: %s\n", msg.Data)
		text := gjson.Get(string(msg.Data), "fullText")
		//fmt.Printf("Message retrieved: %s\n", text.Str)
		_, err := runner.app.Predict(brain.GetID(), text.Str)
		if err != nil {
			runner.Err(err)
		}
		bar.Increment()
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
	bar.Finish()
}
