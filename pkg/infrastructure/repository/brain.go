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
			db: db,
		},
	}
}

func (repo *BrainRepository) GetAll() ([]BrainInterface, error) {
	var brains []Brain
	if err := repo.repository.db.Find(&brains).Error; err != nil {
		return nil, err
	}

	// @See https://github.com/golang/go/wiki/InterfaceSlice#what-can-i-do-instead
	brainsInterfaces := make([]BrainInterface, len(brains))
	for i, brain := range brains {
		copiedBrain := brain
		brainsInterfaces[i] = &copiedBrain
	}

	return brainsInterfaces, nil
}

func (repo *BrainRepository) GetByID(id int64) (BrainInterface, error) {
	var brain Brain
	if err := repo.repository.db.First(&brain, id).Error; err != nil {
		return nil, err
	}

	return &brain, nil
}

func (repo *BrainRepository) GetByName(name string) (BrainInterface, error) {
	var brain Brain
	if err := repo.repository.db.First(&brain, Brain{Name:name}).Error; err != nil {
		return nil, err
	}

	return &brain, nil
}

func (repo *BrainRepository) Create(name string, description string) (BrainInterface, error) {
	brain := Brain{Name: name, Description: description}
	if err := repo.repository.db.Create(&brain).Error; err != nil {
		return nil, err
	}

	return &brain, nil
}

func (repo *BrainRepository) Delete(id int64) error {
	dbc := repo.repository.db.Delete(Brain{}, id)

	if err := dbc.Error; err != nil {
		
		return err
	}

	if dbc.RowsAffected != 1 {

		return errors.New(fmt.Sprintf("no such brain: %d", id))
	}

	return nil
}
