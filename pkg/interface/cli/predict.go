package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (runner *CommandsRunner) NewCmdPredict() *cobra.Command {
	return &cobra.Command{
		Use:   "predict",
		Short: "Predict specified text using certain brain",
		Args:  cobra.MinimumNArgs(2),
		Aliases: []string{"analyze", "analyse", "anal"},
		Run: func(cmd *cobra.Command, args []string) {
			brain, err := runner.app.GetBrainByReference(args[0])
			if err != nil {
				runner.Err(err)

				return
			}

			text := args[1]

			predictForText(runner, brain.GetID(), text)
		},
	}
}

func predictForText(runner *CommandsRunner, brainID int64, text string) {
	score, err := runner.app.Predict(brainID, text)
	if err != nil {
		runner.Err(err)

		return
	}

	for _, classID := range score.GetClassIDs() {
		class, err := runner.app.GetClassByID(classID)
		if err != nil {
			runner.Err(err)
		}
		runner.Out(fmt.Sprintf("%s: %f", class.GetLabel(), score.GetClassProbability(classID)))
	}
}