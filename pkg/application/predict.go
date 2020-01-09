package application

import (
	"github.com/dairlair/sentimentd/pkg/domain/entity"
)

func (app *App) Predict(brainID int64, text string) (prediction entity.Prediction, err error) {
	return app.predictor.Predict(brainID, text)
}