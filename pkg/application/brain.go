package application

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
)

func (app *App) BrainList () ([]BrainInterface, error) {
	return app.brainRepository.GetAll()
}

func (app *App) GetBrainByID (id int64) (BrainInterface, error) {
	return app.brainRepository.GetByID(id)
}

func (app *App) GetBrainByName (name string) (BrainInterface, error) {
	return app.brainRepository.GetByName(name)
}

func (app *App) CreateBrain (name string, description string) (BrainInterface, error) {
	return app.brainRepository.Create(name, description)
}

func (app *App)  DeleteBrain (id int64) error {
	return app.brainRepository.Delete(id)
}