// This package provides ability to interact with sentimentd application through CLI interface.
// Obviously it MUST have ONLY these dependencies:
//   * the application
//   * the input stream
//   * the output stream (standard and error).
//
// How it works:
//   1. You create CmdFactory using the NewCmdFactory function with app, in, out, err.
//   2. The created factory returns Cobra command which interact with app and in/out/err streams.
//   3. You attach these command anywhere.

package cli

import (
	"github.com/dairlair/sentimentd/pkg/application"
	"io"
)

type CmdFactory struct {
	app *application.App
	in io.Reader
	out io.Writer
	err io.Writer
}

func NewCmdFactory(app *application.App, in io.Reader, out, err io.Writer) *CmdFactory {
	return &CmdFactory{
		app: app,
		in:  in,
		out: out,
		err: err,
	}
}