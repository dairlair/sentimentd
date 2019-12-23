package cli

import (
	"github.com/dairlair/sentimentd/pkg/application"
	"os"
	"testing"
)

func TestCommandsRunner_NewCmdBrain(t *testing.T) {
	config := application.Config{}
	app := application.NewApp(config)
	runner := NewCommandsRunner(app, os.Stdin, os.Stdout, os.Stderr)
	runner.NewCmdBrain()
}
