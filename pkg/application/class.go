package application

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
)

func (app *App) GetClassByID (classID int64) (ClassInterface, error) {
	return app.classRepository.FindByID(classID)
}