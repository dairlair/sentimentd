package cmd

import (
	"fmt"
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
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
	Long: `Create a brain`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create: " + strings.Join(args, ";"))
		var name = args[0]
		var description = ""
		if len(args) > 1 {
			description = args[1]
		}
		brain, err := application.CreateBrain(name, description)
		if err != nil {
			return
		}
		console.PrintBrains([]entity.BrainInterface{brain})
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

var brainDeleteCmd = &cobra.Command{
	Use:   "rm <id>",
	Short: "Remove one or more brains",
	Long: `Remove one or more brains`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			id, err := strconv.ParseInt(arg, 10, 64)
			if err != nil {
				fmt.Printf("Error: %s is invalid reference", arg)
				continue
			}
			err = application.DeleteBrain(id)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				fmt.Printf("Deleted: %d\n", id)
			}
		}
	},
}