package cli

import (
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
			for {
				line, err := li.Next()
				if err != nil {
					if err == io.EOF {
						break
					} else {
						// @TODO Move this annoying calls to Out() and Error() methods of command runner.
						_, _ = fmt.Fprintf(runner.err, "error: %s\n", err)
					}
				}

				parts := strings.Split(string(line), " ")
				if len(parts) < 2 {
					_, _ = fmt.Fprintf(runner.err, "error: to few arguments to train\n")
					continue
				}

				err = runner.app.Train(id, Sample{
					Sentence: parts[0],
					Classes:  parts[1:],
				})
				if err != nil {
					_, _ = fmt.Fprintf(runner.err, "error: %s\n", err)
				}
			}
			_, _ = fmt.Fprintf(runner.out, "the training is finished\n")
		},
	}
}