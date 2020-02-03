package cli

import (
	"github.com/spf13/cobra"
)

// NewCmdAnalyse returns command which allows to run command''
func (runner *CommandsRunner) NewCmdAnalyse() *cobra.Command {
	return &cobra.Command{
		Use:     "analyse",
		Short:   "Analyse sentiment of specified text using certain brain",
		Args:    cobra.MinimumNArgs(2),
		Aliases: []string{"analyze", "analyse", "anal"},
		Run: func(cmd *cobra.Command, args []string) {
			brain, err := runner.app.GetBrainByReference(args[0])
			if err != nil {
				runner.Err(err)

				return
			}

			text := args[1]
			analyseText(runner, brain.GetID(), text)
		},
	}
}

func analyseText(runner *CommandsRunner, brainID int64, text string) {
	analysis, err := runner.app.HumanizedPredict(brainID, text)
	if err != nil {
		runner.Err(err)

		return
	}

	json, err := analysis.JSON()
	if err != nil {
		runner.Err(err)

		return
	}

	runner.Out(json)
}
