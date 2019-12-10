package repository

import (
	"errors"
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
	. "github.com/dairlair/sentimentd/pkg/infrastructure/model"
	"github.com/jinzhu/gorm"
)

type BrainRepository struct {
	repository
}

func NewBrainRepository(db *gorm.DB) BrainRepositoryInterface {
	return &BrainRepository{
		repository: repository{
			db:db,
		},
	}
}

func (repo *BrainRepository) GetAll() ([]BrainInterface, error) {
	var brains []Brain
	repo.repository.db.Find(&brains)

	// @See https://github.com/golang/go/wiki/InterfaceSlice#what-can-i-do-instead
	brainsInterfaces := make([]BrainInterface, len(brains))

	for i, brain := range brains {
		copiedBrain := brain
		brainsInterfaces[i] = &copiedBrain
	}

	return brainsInterfaces, nil
}

func (repo *BrainRepository) Create(name string, description string) (BrainInterface, error) {
	brain := Brain{Name: name, Description: description}
	repo.repository.db.Create(&brain)
	return &brain, nil
}

func (repo *BrainRepository) Delete(id int64) error {
	var brain Brain
	brain.Model.ID = id

	if repo.repository.db.Delete(&brain).RowsAffected != 1 {
		return errors.New(fmt.Sprintf("no such brain: %d", id))
	}

	return nil
}