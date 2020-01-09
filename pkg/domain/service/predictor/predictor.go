package predictor

import (
	"fmt"
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"strings"
)

type TokenizerInterface interface {
	Tokenize(sentence string) []string
}

type Predictor struct {
	tokenizer TokenizerInterface
}

func NewPredictor (tokenizer TokenizerInterface) *Predictor {
	return &Predictor{
		tokenizer: tokenizer,
	}
}

func (p *Predictor) Predict (brainID int64, text string) (prediction entity.Prediction, err error) {
	fmt.Printf("Prediction with %d for '%s'\n", brainID, text)
	tokens := p.tokenizer.Tokenize(text)
	fmt.Printf("Found tokens: %s\n", strings.Join(tokens, ", "))
	return prediction, nil
}