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

	prediction := entity.NewPrediction(probabilities)
	return prediction
}

func calculateClassProbability(trainingResult result.TrainingResult, classID int64) float64 {
	return math.Log(float64(trainingResult.ClassFrequency[classID] / trainingResult.SamplesCount))
}
