package application

import (
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	"strings"
)

func (app *App) Train (id int64, sample Sample) error {
	fmt.Printf("Train Brain #%d with sample: '\n", id)
	fmt.Printf("  sentence: %s\n", sample.Sentence)
	fmt.Printf("  classes: %s\n", strings.Join(sample.Classes, ", "))
	return nil
}