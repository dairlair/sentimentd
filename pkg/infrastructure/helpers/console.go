// Here you can find helpers which are provide ability to print domain entities to any file.
package helpers

import (
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	"io"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Console struct {
	writer io.Writer
}

func NewConsole(writer io.Writer) Console {
	return Console{
		writer: writer,
	}
}

func (con *Console) PrintBrains(brains []BrainInterface) {
	cols := []string{"ID", "NAME", "DESCRIPTION"}
	var rows [][]string
	for _, brain := range brains {
		row := []string{
			strconv.FormatInt(brain.GetID(), 10),
			brain.GetName(),
			brain.GetDescription(),
		}
		rows = append(rows, row)
	}
	printTable(cols, rows)
}

func printTable(cols []string, rows [][]string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', tabwriter.TabIndent)
	printRow(w, cols)
	for _, row := range rows {
		printRow(w, row)
	}
	_ = w.Flush()
}

func printRow(w io.Writer, cols []string) {
	_, _ = fmt.Fprintln(w, strings.Join(cols, "\t")+"\t")
}