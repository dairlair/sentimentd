package classifier

import (
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/domain/service/training/result"
	"math"
)

type NaiveBayesClassifier struct {
	trainingResult result.TrainingResult
}

func NewNaiveBayesClassifier (trainingResult result.TrainingResult) *NaiveBayesClassifier {
	return &NaiveBayesClassifier{
		trainingResult: trainingResult,
	}
}

/**
 * @TODO Remove direct access to TrainingResult and use TrainingResultInterface instead
 */
func (c *NaiveBayesClassifier) Classify () entity.Prediction {
	probabilities := make(map[int64]float64, len(c.trainingResult.ClassFrequency))
	for classID, _ := range c.trainingResult.ClassFrequency {
		probabilities[classID] = calculateClassProbability(c.trainingResult, classID)
	}

	// Create an probability space
	probabilities = createProbabilitySpace(probabilities)

	prediction := entity.NewPrediction(probabilities)
	return prediction
}

func calculateClassProbability(trainingResult result.TrainingResult, classID int64) float64 {
	return math.Log(float64(trainingResult.ClassFrequency[classID]) / float64(trainingResult.SamplesCount))
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