// This repository allows to save the training results
package repository

import (
	"github.com/dairlair/sentimentd/pkg/domain/service/training/result"
	"github.com/jinzhu/gorm"
)

type resultsRepository struct {
	repository
}
func NewResultsRepository(db *gorm.DB) resultsRepository {
	return resultsRepository{
		repository: repository{
			db: db,
		},
	}
}

// These function carefully save the training results
func (r resultsRepository) SaveResult(result result.TrainingResult) error {
	return nil
}