package application

import (
	"github.com/dairlair/sentimentd/pkg/domain/entity"
)

// Train implements AppInterface and allows to train brains
func (app *App) Train(brainID int64, samples []entity.Sample, cb func()) error {
	return app.trainingService.Train(brainID, samples, cb)
}
