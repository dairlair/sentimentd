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

	tx := repo.db.Begin()

	training := model.Training{
		BrainID:      brainID,
		Comment:      "-",
		SamplesCount: result.SamplesCount,
	}

	if err := tx.Create(&training).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Save the classes frequencies
	if err := saveTrainingClasses(tx, training, result); err != nil {
		tx.Rollback()
		return nil
	}

	// Save the tokens frequencies
	if err := saveTrainingTokens(tx, training, result); err != nil {
		tx.Rollback()
		return nil
	}

	return tx.Commit().Error
}

func saveTrainingClasses(db *gorm.DB, training model.Training, result result.TrainingResult) error {
	for classID, samplesCount := range result.ClassFrequency {
		trainingClass := model.TrainingClass{
			BrainID:      training.BrainID,
			TrainingID:   training.Model.ID,
			ClassID:      classID,
			SamplesCount: samplesCount,
		}
		if err := db.Create(&trainingClass).Error; err != nil {
			return err
		}
	}

	return nil
}

func saveTrainingTokens(db *gorm.DB, training model.Training, result result.TrainingResult) error {
	for classID, tokensFrequencies := range result.TokenFrequency {
		for tokenID, samplesCount := range tokensFrequencies {
			trainingClass := model.TrainingToken{
				BrainID:      training.BrainID,
				TrainingID:   training.Model.ID,
				ClassID:      classID,
				TokenID:      tokenID,
				SamplesCount: samplesCount,
			}
			if err := db.Create(&trainingClass).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func (repo resultRepository) GetTrainingResults(brainID int64) (r result.TrainingResult, err error) {
	r.SamplesCount, err = getBrainSamplesCount(repo.db, brainID)
	if err != nil {
		return r, err
	}
	r.ClassFrequency, err = getClassFrequency(repo.db, brainID)
	if err != nil {
		return r, err
	}
	return r, nil
}

func getBrainSamplesCount(db *gorm.DB, brainID int64) (int64, error) {
	r := struct{ SamplesCount int64 }{}
	sql := "SELECT SUM(samples_count) AS samples_count FROM trainings WHERE deleted_at IS NULL and brain_id = ?"
	if err := db.Raw(sql, brainID).Scan(&r).Error; err != nil {
		return 0, err
	}

	return r.SamplesCount, nil
}

func getClassFrequency(db *gorm.DB, brainID int64) (r result.ClassFrequency, err error) {
	r = make(map[int64]int64)

	rows, err := db.
		Table("training_classes").
		Select("class_id, SUM(samples_count) AS samples_count").
		Where("deleted_at IS NULL and brain_id = ?", brainID).
		Group("class_id").
		Rows()


	if err != nil {
		return r, err
	}

	defer rows.Close()

	for rows.Next() {
		classFrequency := struct {
			ClassID      int64
			SamplesCount int64
		}{}
		err = db.ScanRows(rows, &classFrequency)
		if err != nil {
			return r, err
		}
		r[classFrequency.ClassID] = classFrequency.SamplesCount
	}

	return r, nil
}
