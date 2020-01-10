package classifier

import (
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/domain/service/training/result"
	"math"
)

type TrainedModelInterface interface {
	// Returns the total samples count used in the model's training dataset
	GetSamplesCount() int64
	// Returns map with classes frequencies ()
	GetClassFrequency() result.ClassFrequency
	// Returns count of unique tokens in the model's training dataset
	GetUniqueTokensCount() int64
	// Returns total tokens count in each class map[<Class ID> => <Tokens Count>]
	GetClassSizes() entity.ClassSizeMap
	// Returns token frequency (how much time each token was found in each class)
	GetTokenFrequency() result.TokenFrequency
}

type NaiveBayesClassifier struct {
	model TrainedModelInterface
}

func NewNaiveBayesClassifier (model TrainedModelInterface) *NaiveBayesClassifier {
	return &NaiveBayesClassifier{
		model: model,
	}
}

func (c *NaiveBayesClassifier) Classify () entity.Prediction {
	probabilities := make(map[int64]float64, len(c.model.GetClassFrequency()))
	for classID, _ := range c.model.GetClassFrequency() {
		probabilities[classID] = calculateClassProbability(c.model, classID)
	}

	// Create an probability space
	probabilities = createProbabilitySpace(probabilities)

	prediction := entity.NewPrediction(probabilities)
	return prediction
}

func calculateClassProbability(model TrainedModelInterface, classID int64) float64 {
	r := math.Log(float64(model.GetClassFrequency()[classID]) / float64(model.GetSamplesCount()))

	return r
}

func createProbabilitySpace(probabilities map[int64]float64) map[int64]float64 {
	var denominator float64 = 0
	for _, probability := range probabilities {
		denominator += math.Exp(probability)
	}

	for classID, probability := range probabilities {
		probabilities[classID] = math.Exp(probability) / denominator
	}

	return probabilities
}