package cli

import (
	"bufio"
	"errors"
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/interface/cli/util"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strings"
)

func (runner *CommandsRunner) NewCmdTrain() *cobra.Command {
	return &cobra.Command{
		Use:   "train",
		Short: "Train a specified brain",
		Long: `train provides ability to train brain with categorized sentences.
Find more information at: https://google.com`,
		Run: func(cmd *cobra.Command, args []string) {
			brainID, err := util.ParseInt64(args[0])
			if err != nil {
				runner.Err(err)

				return
			}

			if len(args) == 1 {
				trainFromStream(runner, brainID, runner.in)
			} else {
				for _, filename := range args[1:] {
					trainFromFile(runner, brainID, filename)
				}
			}
		},
	}
}

func trainFromFile(runner *CommandsRunner, brainID int64, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	runner.Out(fmt.Sprintf("Training with dataset %s", filename))
	trainFromStream(runner, brainID, file)
}

func trainFromStream(runner *CommandsRunner, brainID int64, in io.Reader) {
	scanner := bufio.NewScanner(in)
	var samples []Sample
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) < 2 {
			runner.Err(errors.New("to few arguments to train"))
			continue
		}
		sample := Sample{
			Sentence: parts[0],
			Classes:  parts[1:],
		}
		samples = append(samples, sample)
	}
	result, err := runner.app.Train(brainID, samples)
	if err != nil {
		log.Fatalf("training error: %s\n", err)
	}
	runner.Out("the training is finished")
	runner.Out(fmt.Sprintf("result: %v", result))
}
