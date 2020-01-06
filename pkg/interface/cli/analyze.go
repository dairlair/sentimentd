package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (runner *CommandsRunner) NewCmdAnalyze() *cobra.Command {
	return &cobra.Command{
		Use:   "analyze",
		Short: "Analyze specified text using certain brain",
		Args:  cobra.MinimumNArgs(2),
		Aliases: []string{"anal", "score"},
		Run: func(cmd *cobra.Command, args []string) {
			brain, err := runner.app.GetBrainByReference(args[0])
			if err != nil {
				runner.Err(err)

				return
			}

			text := args[1]

			analyzeText(runner, brain.GetID(), text)
		},
	}
}

func analyzeText(runner *CommandsRunner, brainID int64, text string) {
	score, err := runner.app.Analyze(brainID, text)
	if err != nil {
		runner.Err(err)

		return
	}

	runner.Out(fmt.Sprintf("Score: %v", score))
}