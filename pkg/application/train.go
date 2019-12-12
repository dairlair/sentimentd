package application

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
)

func (app *App) Train (brainID int64, samples []Sample) error {
	return app.trainService.Train(brainID, samples)
}