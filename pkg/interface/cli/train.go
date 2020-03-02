package cli

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/spf13/cobra"
	"time"
)

// NewCmdTrain returns command for brain training
func (runner *CommandsRunner) NewCmdTrain() *cobra.Command {
	return &cobra.Command{
		Use:   "train",
		Short: "Train a specified brain",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			brain, err := runner.app.GetBrainByReference(args[0])
			if err != nil {
				runner.Err(err)

				return
			}

			if len(args) == 1 {
				trainFromStream(runner, brain.GetID(), runner.in)
			} else {
				for _, filename := range args[1:] {
					trainFromFile(runner, brain.GetID(), filename)
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
	samples := readSamples(runner, in)
	err := runner.app.Train(brainID, samples)
	if err != nil {
		runner.Err(err)
	}
	runner.Out("the training is finished")
}

func readSamples(runner *CommandsRunner, in io.Reader) []entity.Sample {
	reader := csv.NewReader(in)
	var samples []entity.Sample
	t := time.Now()
	for {
		columns, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			runner.Err(err)
			continue
		}
		if len(columns) != 2 {
			runner.Err(errors.New("wrong columns number in row: " + strings.Join(columns, ", ")))
			continue
		}
		sample := entity.Sample{
			Sentence: columns[1],
			Classes:  strings.Split(columns[0], ","),
		}
		samples = append(samples, sample)
	}
	duration := time.Since(t)
	samplesPerSecond := float64(len(samples)) / duration.Seconds()
	runner.Out(fmt.Sprintf("Dataset read for %f seconds (%f samples in second)", duration.Seconds(), samplesPerSecond))

	return samples
}
