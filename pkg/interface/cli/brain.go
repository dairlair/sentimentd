package cli

import (
	"fmt"
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/interface/cli/util"
	"github.com/spf13/cobra"
)

func (runner *CommandsRunner) NewCmdBrain() *cobra.Command {
	var brainCmd = &cobra.Command{
		Use:   "brain",
		Short: "Brains operations",
		Long:  `Brains is the predictive models of sentimentd service`,
	}

	brainCmd.AddCommand(newCmdBrainCreate(runner))
	brainCmd.AddCommand(newCmdBrainList(runner))
	brainCmd.AddCommand(newCmdBrainInspect(runner))
	brainCmd.AddCommand(newCmdBrainDelete(runner))

	return brainCmd
}

func newCmdBrainCreate(runner *CommandsRunner) *cobra.Command {
	return &cobra.Command{
		Use:   "create <name> [description]",
		Short: "Create a brain",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var name = args[0]
			var description = ""
			if len(args) > 1 {
				description = args[1]
			}
			brain, err := runner.app.CreateBrain(name, description)
			if err != nil {
				runner.Err(err)

				return
			}
			util.PrintBrains(runner.out, []entity.BrainInterface{brain})
		},
	}
}

func newCmdBrainList(runner *CommandsRunner) *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List brains",
		Run: func(cmd *cobra.Command, args []string) {
			brains, err := runner.app.BrainList()
			if err != nil {
				runner.Err(err)

				return
			}
			util.PrintBrains(runner.out, brains)
		},
	}
}

func newCmdBrainInspect(runner *CommandsRunner) *cobra.Command {
	return &cobra.Command{
		Use:   "inspect <id>",
		Short: "Display detailed information on one or more brains",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			util.IterateArgs(args, func(reference string) {
				brain, err := runner.app.GetBrainByReference(reference)
				if err != nil {
					runner.Err(err)

					return
				}
				util.PrintBrains(runner.out, []entity.BrainInterface{brain})
			})
		},
	}
}

func newCmdBrainDelete(runner *CommandsRunner) *cobra.Command {
	return &cobra.Command{
		Use:   "rm <id>",
		Short: "Remove one or more brains",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			util.IterateArgs(args, func(reference string) {
				brain, err := runner.app.GetBrainByReference(reference)
				if err != nil {
					runner.Err(err)

					return
				}

				err = runner.app.DeleteBrain(brain.GetID())
				if err != nil {
					runner.Err(err)

					return
				}
				runner.Out(fmt.Sprintf("Deleted: %d\n", brain.GetID()))
			})
		},
	}
}
