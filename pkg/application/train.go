package application

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
)

func (app *App) Train (brainID int64, samples []Sample, cb func ()) error {
	return app.trainingService.Train(brainID, samples, cb)
}