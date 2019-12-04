package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sentimentd",
		Short: "sentimentd short",
		Long: `sentimentd
long`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}