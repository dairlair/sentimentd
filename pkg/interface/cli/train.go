package cli

import (
	"fmt"
	"github.com/dairlair/sentimentd/pkg/interface/cli/util"
	"github.com/spf13/cobra"
	"io"
)

func (f *CmdFactory) NewCmdTrain () *cobra.Command {
	return &cobra.Command{
		Use:   "train",
		Short: "Train a specified brain",
		Long: `
      train provides ability to train brain with categorized sentences.
      Find more information at:
            https://google.com/`,
		Run: func(cmd *cobra.Command, args []string) {
			li := util.NewLineIterator(f.in)
			for {
				line, err := li.Next()
				if err != nil {
					if err == io.EOF {
						break
					} else {
						_, _ = fmt.Fprintf(f.err, "error: %s\n", err)
					}
				}
				_, _ = fmt.Fprintf(f.out, "line: %s\n", line)
			}
		},
	}
}
//func runTrain()