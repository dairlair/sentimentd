package predictor

import (
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/domain/service/classifier"
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

// Predictor is a structure which does tokenization and pass tokenized input with trained model to Classifier
type Predictor struct {
	tokenizer         TokenizerInterface
	tokenRepository   TokenRepositoryInterface
	resultsRepository ResultsRepositoryInterface
}

// NewPredictor returns a Predictor instance
func NewPredictor(
	tokenizer TokenizerInterface,
	tokenRepository TokenRepositoryInterface,
	resultsRepository ResultsRepositoryInterface,
) *Predictor {
	resultsRepositoryCache := newResultsRepositoryCache(resultsRepository)
	return &Predictor{
		tokenizer:         tokenizer,
		tokenRepository:   tokenRepository,
		resultsRepository: resultsRepositoryCache,
	}
}

// Predict returns prediction based on trained model provided by specified brain
func (p *Predictor) Predict(brainID int64, text string) (prediction entity.Prediction, err error) {
	// Divide text into the tokens
	tokens := p.tokenizer.Tokenize(text)

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
	trainedModel, err := p.resultsRepository.GetTrainedModel(brainID)
	if err != nil {
		return prediction, err
	}

	c := classifier.NewNaiveBayesClassifier(trainedModel)

	prediction = c.Classify(tokenIDs)

	return prediction, nil
}

type resultsRepositoryCache struct {
	repo  ResultsRepositoryInterface
	cache map[int64]classifier.TrainedModelInterface
}

func newResultsRepositoryCache(repo ResultsRepositoryInterface) *resultsRepositoryCache {
	return &resultsRepositoryCache{
		repo:  repo,
		cache: make(map[int64]classifier.TrainedModelInterface, 0),
	}
}

func (c *resultsRepositoryCache) GetTrainedModel(brainID int64) (classifier.TrainedModelInterface, error) {
	if _, ok := c.cache[brainID]; !ok {
		trainedModel, err := c.repo.GetTrainedModel(brainID)
		if err != nil {

			return nil, err
		}
		c.cache[brainID] = trainedModel
	}

	return c.cache[brainID], nil
}
