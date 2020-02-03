package application

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
)

func (app *App) BrainList() ([]BrainInterface, error) {
	return app.brainRepository.GetAll()
}

func (app *App) GetBrainByReference(reference string) (BrainInterface, error) {
	return app.brainRepository.GetByReference(reference)
}

func (app *App) CreateBrain(name string, description string) (BrainInterface, error) {
	return app.brainRepository.Create(name, description)
}

func (app *App) DeleteBrain(id int64) error {
	return app.brainRepository.Delete(id)
}
