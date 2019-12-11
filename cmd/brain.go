package cmd

import (
	"fmt"
	"github.com/dairlair/sentimentd/cmd/utils"
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(brainCmd)
	brainCmd.AddCommand(brainCreateCmd)
	brainCmd.AddCommand(brainListCmd)
	brainCmd.AddCommand(brainInspectCmd)
	brainCmd.AddCommand(brainDeleteCmd)
}

var brainCmd = &cobra.Command{
	Use:   "brain",
	Short: "Brains operations",
	Long:  `Brains is the predictive models of sentimentd service`,
}

var brainCreateCmd = &cobra.Command{
	Use:   "create <name> [description]",
	Short: "Create a brain",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var name = args[0]
		var description = ""
		if len(args) > 1 {
			description = args[1]
		}
		brain, err := app.CreateBrain(name, description)
		if err != nil {
			fmt.Printf("Error: %s\n", err)

			return
		}
		console.PrintBrains([]entity.BrainInterface{brain})
	},
}

var brainListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List brains",
	Run: func(cmd *cobra.Command, args []string) {
		brains, err := app.BrainList()
		if err != nil {
			fmt.Printf("Error: %s\n", err)

			return
		}
		console.PrintBrains(brains)
	},
}

var brainInspectCmd = &cobra.Command{
	Use:   "inspect <id>",
	Short: "Display detailed information on one or more brains",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		utils.IterateArgs(args, func(id int64) {
			brain, err := app.GetBrainByID(id)
			if err != nil {
				fmt.Printf("Error: %s\n", err)

				return
			}
			console.PrintBrains([]entity.BrainInterface{brain})
		})
	},
}

var brainDeleteCmd = &cobra.Command{
	Use:   "rm <id>",
	Short: "Remove one or more brains",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.IterateArgs(args, func(id int64) {
			err := app.DeleteBrain(id)
			if err != nil {
				fmt.Printf("Error: %s\n", err)

				return
			}
			fmt.Printf("Deleted: %d\n", id)
		})
	},
}
