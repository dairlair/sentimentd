package application

import (
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"time"
)

func (app *App) Predict(brainID int64, text string) (prediction entity.Prediction, err error) {
	return app.predictor.Predict(brainID, text)
}

func (app *App) HumanizedPredict(brainID int64, text string) (entity.HumanizedPrediction, error) {
	t := time.Now()

	prediction, err := app.Predict(brainID, text)
	if err != nil {
		return entity.HumanizedPrediction{}, err
	}

	probabilities := make(map[string]float64, len(prediction.GetClassIDs()))
	for _, classID := range prediction.GetClassIDs() {
		class, _ := app.GetClassByID(classID)
		probabilities[class.GetLabel()] = prediction.GetClassProbability(classID)
	}

	duration := time.Since(t)

	return entity.NewHumanizedPrediction(probabilities, duration), nil
}