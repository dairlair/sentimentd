package application

import (
	"github.com/dairlair/sentimentd/pkg/domain/entity"
)

func (app *App) Predict(brainID int64, text string) (prediction entity.Prediction, err error) {
	return entity.NewPrediction(map[int64]float64{}), err
}