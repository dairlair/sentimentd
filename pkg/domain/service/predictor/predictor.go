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
	GetTokenIDs(brainID int64, tokens []string) ([]int64, error)
}

type ResultsRepositoryInterface interface {
	GetTrainedModel(brainID int64) (classifier.TrainedModelInterface, error)
}

type Predictor struct {
	tokenizer         TokenizerInterface
	tokenRepository   TokenRepositoryInterface
	resultsRepository ResultsRepositoryInterface
}

func NewPredictor(
	tokenizer TokenizerInterface,
	tokenRepository TokenRepositoryInterface,
	resultsRepository ResultsRepositoryInterface,
) *Predictor {
	return &Predictor{
		tokenizer:         tokenizer,
		tokenRepository:   tokenRepository,
		resultsRepository: resultsRepository,
	}
}

func (p *Predictor) Predict(brainID int64, text string) (prediction entity.Prediction, err error) {
	fmt.Printf("Prediction with brain#%d for text '%s'\n", brainID, text)

	// Divide text into the tokens
	tokens := p.tokenizer.Tokenize(text)
	fmt.Printf("Found tokens: %s\n", strings.Join(tokens, ", "))

	// Found these tokens in the TokenRepository
	tokenIDs, err := p.tokenRepository.GetTokenIDs(brainID, tokens)
	if err != nil {
		return prediction, err
	}
	// @TODO Replace to the Laplace smothering
	for len(tokenIDs) < len(tokens) {
		tokenIDs = append(tokenIDs, 0)
	}

	// Retrieve summarized training data for specified brain
	trainingResult, err := p.resultsRepository.GetTrainedModel(brainID)
	if err != nil {
		return prediction, err
	}

	c := classifier.NewNaiveBayesClassifier(trainingResult)

	prediction = c.Classify(tokenIDs)

	return prediction, nil
}
