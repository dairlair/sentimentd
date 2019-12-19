// This package provides ability to interact with sentimentd application through CLI interface.
// Obviously it MUST have ONLY these dependencies:
//   * the application
//   * the input stream
//   * the output stream (standard and error).
//
// How it works:
//   1. You create CommandsRunner using the NewCommandsRunner function with app, in, out, err.
//   2. The created factory returns Cobra command which interact with app and in/out/err streams.
//   3. You attach these command anywhere.

package cli

import (
	"fmt"
	"github.com/dairlair/sentimentd/pkg/application"
	"io"
)

type CommandsRunner struct {
	app *application.App
	in io.Reader
	out io.Writer
	err io.Writer
}

func NewCommandsRunner(app *application.App, in io.Reader, out, err io.Writer) *CommandsRunner {
	return &CommandsRunner{
		app: app,
		in:  in,
		out: out,
		err: err,
	}
}

func (runner *CommandsRunner) Out (s string) {
	writeToStream(runner.out, fmt.Sprintf("%s\n", s))
}

func (runner *CommandsRunner) Err (err error) {
	writeToStream(runner.err, fmt.Sprintf("error: %s\n", err))
}

func writeToStream(stream io.Writer, s string) {
	if _, err := fmt.Fprintf(stream, s); err != nil {
		panic("can not write to the output stream")
	}
}