package cmd

import (
	"testing"
)

func BenchmarkRoot_Analyse(b *testing.B) {
	rootCmd.SetArgs([]string{"analyse", "skynet", "Text should to be analysed"})
	rootCmd.Execute()
}
