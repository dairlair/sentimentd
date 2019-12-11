// This package provides ability to interact with sentimentd application through CLI interface.
// Obviously it MUST have ONLY these dependencies:
//   * the application
//   * the input stream
//   * the output stream (standard and error).
package cli

import (
	"github.com/dairlair/sentimentd/pkg/application"
	"io"
)

type CommandLineInterface struct {
	app *application.App
	stdin io.Reader
	stdout io.Writer
	stderr io.Writer
}