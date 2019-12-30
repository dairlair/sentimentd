package cli

import (
	"github.com/dairlair/sentimentd/pkg/application"
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	cli "github.com/dairlair/sentimentd/pkg/mocks/interface/cli"
	"os"
	"testing"
)

func TestCommandsRunner_NewCmdBrain(t *testing.T) {
	config := application.Config{}
	app := application.NewApp(config)
	runner := NewCommandsRunner(app, os.Stdin, os.Stdout, os.Stderr)
	runner.NewCmdBrain()
}

func TestBrainListCommand_Successful(t *testing.T) {
	appMock := cli.AppInterface{}
	appMock.On("BrainList").Return([]entity.BrainInterface{}, nil)
	runner := NewCommandsRunner(&appMock, os.Stdin, os.Stdout, os.Stderr)
	command := newCmdBrainList(runner)
	command.Run(command, []string{})
	appMock.AssertExpectations(t)
}
