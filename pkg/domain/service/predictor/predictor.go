package predictor

import (
	"fmt"
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/domain/service/classifier"
	"strings"
)

type TokenizerInterface interface {
	Tokenize(sentence string) []string
}

type TokenRepositoryInterface interface {
	GetTokenIDs(tokens []string) ([]int64, error)
}

type ResultsRepositoryInterface interface {
	GetTrainedModel(brainID int64) (classifier.TrainedModelInterface, error)
}

type Predictor struct {
	tokenizer TokenizerInterface
	resultsRepository ResultsRepositoryInterface
}

func NewPredictor (tokenizer TokenizerInterface, resultsRepository ResultsRepositoryInterface ) *Predictor {
	return &Predictor{
		tokenizer: tokenizer,
		resultsRepository: resultsRepository,
	}
}

func (p *Predictor) Predict (brainID int64, text string) (prediction entity.Prediction, err error) {
	fmt.Printf("Prediction with %d for '%s'\n", brainID, text)

	// Divide text into the tokens
	tokens := p.tokenizer.Tokenize(text)
	fmt.Printf("Found tokens: %s\n", strings.Join(tokens, ", "))

	// Retrieve summarized training data for specified brain
	trainingResult, err := p.resultsRepository.GetTrainedModel(brainID)
	if err != nil {
		return prediction, err
	}

	c := classifier.NewNaiveBayesClassifier(trainingResult)

	prediction = c.Classify()

	return prediction, nil
}