package util

import (
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"
)

func PrintBrains(out io.Writer, brains []BrainInterface) {
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
	printTable(out, cols, rows)
}

func printTable(out io.Writer, cols []string, rows [][]string) {
	w := tabwriter.NewWriter(out, 0, 0, 8, ' ', tabwriter.TabIndent)
	printRow(w, cols)
	for _, row := range rows {
		printRow(w, row)
	}
	_ = w.Flush()
}

func printRow(out io.Writer, cols []string) {
	_, _ = fmt.Fprintln(out, strings.Join(cols, "\t")+"\t")
}