package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(trainCmd)
	brainCmd.AddCommand(brainCreateCmd)
	brainCmd.AddCommand(brainListCmd)
	brainCmd.AddCommand(brainInspectCmd)
	brainCmd.AddCommand(brainDeleteCmd)
}

var trainCmd = &cobra.Command{
	Use:   "train <brain>",
	Short: "Train specified brain with data read from STDIN",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Train brain %s\n", args[0])
	},
}