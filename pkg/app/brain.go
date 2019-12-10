package app

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
)

func (app *App) BrainList () ([]BrainInterface, error) {
	return app.brainRepository.GetAll()
}

func (app *App) CreateBrain (name string, description string) (BrainInterface, error) {
	return app.brainRepository.Create(name, description)
}