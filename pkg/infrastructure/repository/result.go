// This repository allows to save the training results
package repository

import (
	"github.com/dairlair/sentimentd/pkg/domain/service/training/result"
	"github.com/dairlair/sentimentd/pkg/infrastructure/model"
	"github.com/jinzhu/gorm"
)

type resultRepository struct {
	repository
}

func NewResultRepository(db *gorm.DB) resultRepository {
	return resultRepository{
		repository: repository{
			db: db,
		},
	}
}

// These function carefully save the training results
func (repo resultRepository) SaveResult(brainID int64, result result.TrainingResult) error {
	// @TODO Wrap to transaction this code, add defer rollback and commit
	training := model.Training{
		BrainID:      brainID,
		Comment:      "-",
		SamplesCount: result.SamplesCount,
	}

	if err := repo.repository.db.Create(&training).Error; err != nil {
		return err
	}

	// Save the classes frequencies
	if err := repo.saveTrainingClasses(training, result); err != nil {
		return nil
	}

	return nil
}

func (repo resultRepository) saveTrainingClasses(training model.Training, result result.TrainingResult) error {
	for classID, samplesCount := range result.ClassFrequency {
		trainingClass := model.TrainingClass{
			BrainID:      training.BrainID,
			TrainingID:   training.Model.ID,
			ClassID:      classID,
			SamplesCount: samplesCount,
		}
		if err := repo.db.Create(&trainingClass).Error; err != nil {
			return err
		}
	}

	return nil
}