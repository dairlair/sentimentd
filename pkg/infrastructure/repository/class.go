package repository

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
	. "github.com/dairlair/sentimentd/pkg/infrastructure/model"
	"github.com/jinzhu/gorm"
)

type ClassRepository struct {
	repository
}

func NewClassRepository(db *gorm.DB) ClassRepositoryInterface {
	return &ClassRepository{
		repository: repository{
			db: db,
		},
	}
}

func (repo *ClassRepository) Create(brainID int64, label string) (ClassInterface, error) {
	class := Class{BrainID: brainID, Label:label}
	if err := repo.repository.db.Create(&class).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (repo *ClassRepository) FindByBrainAndLabel(brainID int64, label string) (ClassInterface, error) {
	class := Class{BrainID: brainID, Label:label}
	if err := repo.repository.db.First(&class).Error; err != nil {
		return nil, err
	}

	return &class, nil
}