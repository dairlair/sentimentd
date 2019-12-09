package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(brainCmd)
	brainCmd.AddCommand(brainCreateCmd)
	brainCmd.AddCommand(brainListCmd)
	brainCmd.AddCommand(brainInspectCmd)
	brainCmd.AddCommand(brainRemoveCmd)
}

var brainCmd = &cobra.Command{
	Use:   "brain",
	Short: "Brains operations",
	Long:  `Brains is the predictive models of sentimentd service`,
}

var brainCreateCmd = &cobra.Command{
	Use:   "create <name> [description]",
	Short: "Create a brain",
	Long: `Create a brain`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create: " + strings.Join(args, ";"))
	},
}

var brainListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List brains",
	Long: `List brains`,
	Run: func(cmd *cobra.Command, args []string) {
		brains, err := application.BrainList()
		if err != nil {
			return
		}
		console.PrintBrains(brains)
	},
}

var brainInspectCmd = &cobra.Command{
	Use:   "inspect <id>",
	Short: "Display detailed information on one or more brains",
	Long: `Display detailed information on one or more brains`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Inspect: " + strings.Join(args, ";"))
	},
}

var brainRemoveCmd = &cobra.Command{
	Use:   "rm <id>",
	Short: "Remove one or more brains",
	Long: `Remove one or more brains`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Remove: " + strings.Join(args, ";"))
	},
}