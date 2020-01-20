package cli

import (
	"errors"
	"github.com/dairlair/sentimentd/pkg/application"
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	mocks "github.com/dairlair/sentimentd/pkg/mocks/domain/entity"
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

func TestBrainListCommand_Failed(t *testing.T) {
	appMock := cli.AppInterface{}
	appMock.On("BrainList").Return(nil, errors.New("fail"))
	runner := NewCommandsRunner(&appMock, os.Stdin, os.Stdout, os.Stderr)
	command := newCmdBrainList(runner)
	command.Run(command, []string{})
	appMock.AssertExpectations(t)
}

func TestBrainCreateCommand_Successful(t *testing.T) {
	brainStub := createBrainStub()
	appMock := cli.AppInterface{}
	appMock.On("CreateBrain", brainStub.GetName(), brainStub.GetDescription()).Return(&brainStub, nil)
	runner := NewCommandsRunner(&appMock, os.Stdin, os.Stdout, os.Stderr)
	command :=newCmdBrainCreate(runner)
	command.Run(command, []string{brainStub.GetName(), brainStub.GetDescription()})
	appMock.AssertExpectations(t)
}

func TestBrainCreateCommand_Failed(t *testing.T) {
	brainStub := createBrainStub()
	appMock := cli.AppInterface{}
	appMock.On("CreateBrain", brainStub.GetName(), brainStub.GetDescription()).Return(nil, errors.New("fail"))
	runner := NewCommandsRunner(&appMock, os.Stdin, os.Stdout, os.Stderr)
	command :=newCmdBrainCreate(runner)
	command.Run(command, []string{brainStub.GetName(), brainStub.GetDescription()})
	appMock.AssertExpectations(t)
}

func TestBrainInspectCommand_Successful(t *testing.T) {
	brainStub := createBrainStub()
	appMock := cli.AppInterface{}
	appMock.On("GetBrainByReference", "skynet").Return(&brainStub, nil)
	runner := NewCommandsRunner(&appMock, os.Stdin, os.Stdout, os.Stderr)
	command :=newCmdBrainInspect(runner)
	command.Run(command, []string{"skynet"})
	appMock.AssertExpectations(t)
}

func TestBrainInspectCommand_Failed(t *testing.T) {
	appMock := cli.AppInterface{}
	appMock.On("GetBrainByReference", "skynet").Return(nil, errors.New("fail"))
	runner := NewCommandsRunner(&appMock, os.Stdin, os.Stdout, os.Stderr)
	command :=newCmdBrainInspect(runner)
	command.Run(command, []string{"skynet"})
	appMock.AssertExpectations(t)
}

func TestBrainDeleteCommand_Successful(t *testing.T) {
	brainStub := createBrainStub()
	appMock := cli.AppInterface{}
	appMock.On("GetBrainByReference", "skynet").Return(&brainStub, nil)
	appMock.On("DeleteBrain", brainStub.GetID()).Return(nil)
	runner := NewCommandsRunner(&appMock, os.Stdin, os.Stdout, os.Stderr)
	command :=newCmdBrainDelete(runner)
	command.Run(command, []string{"skynet"})
	appMock.AssertExpectations(t)
}

func TestBrainDeleteCommand_Failed_NotFound(t *testing.T) {
	appMock := cli.AppInterface{}
	appMock.On("GetBrainByReference", "skynet").Return(nil, errors.New("fail"))
	runner := NewCommandsRunner(&appMock, os.Stdin, os.Stdout, os.Stderr)
	command :=newCmdBrainDelete(runner)
	command.Run(command, []string{"skynet"})
	appMock.AssertExpectations(t)
}

func TestBrainDeleteCommand_Failed_NotDeleted(t *testing.T) {
	brainStub := createBrainStub()
	appMock := cli.AppInterface{}
	appMock.On("GetBrainByReference", "skynet").Return(&brainStub, nil)
	appMock.On("DeleteBrain", brainStub.GetID()).Return(errors.New("fail"))
	runner := NewCommandsRunner(&appMock, os.Stdin, os.Stdout, os.Stderr)
	command :=newCmdBrainDelete(runner)
	command.Run(command, []string{"skynet"})
	appMock.AssertExpectations(t)
}

func createBrainStub() mocks.BrainInterface {
	id := int64(1)
	name := "Skynet"
	description := "Artificial Intelligence from Terminator movie"
	brainStub := mocks.BrainInterface{}
	brainStub.On("GetID").Return(id)
	brainStub.On("GetName").Return(name)
	brainStub.On("GetDescription").Return(description)
	return brainStub
}