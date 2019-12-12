package cli

import (
	"errors"
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/interface/cli/util"
	"github.com/spf13/cobra"
	"io"
	"strings"
)

func (runner *CommandsRunner) NewCmdTrain () *cobra.Command {
	return &cobra.Command{
		Use:   "train",
		Short: "Train a specified brain",
		Long: `train provides ability to train brain with categorized sentences.
Find more information at: https://google.com`,
		Run: func(cmd *cobra.Command, args []string) {
			id, err := util.ParseInt64(args[0])
			if err != nil {
				_, _ = fmt.Fprintf(runner.err, "Error: %s\n", err)
				return
			}

			li := util.NewLineIterator(runner.in)
			// @TODO Implement runner.Iterate method for similar loops with reading from input stream
			var samples []Sample
			for {
				line, err := li.Next()
				if err != nil {
					if err == io.EOF {
						break
					}
					runner.Err(err)
				}

				parts := strings.Split(string(line), " ")
				if len(parts) < 2 {
					runner.Err(errors.New("to few arguments to train"))
					continue
				}

				sample := Sample{
					Sentence: parts[0],
					Classes:  parts[1:],
				}
				samples = append(samples, sample)
				if err != nil {
					runner.Err(err)
				}
			}
			err = runner.app.Train(id, samples)
			runner.Out("the training is finished")
		},
	}
}