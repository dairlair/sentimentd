package application

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/domain/service/training/result"
)

func (app *App) Train (brainID int64, samples []Sample) (result.TrainingResult, error) {
	return app.trainingService.Train(brainID, samples)
}