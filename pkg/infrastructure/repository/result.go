// This repository allows to save the training results
package repository

import (
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/domain/service/classifier"
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

func (repo resultRepository) GetTrainedModel(brainID int64) (trainedModel classifier.TrainedModelInterface, err error) {
	m := entity.TrainedModel{}
	if m.SamplesCount, err = getSamplesCount(repo.db, brainID); err != nil {
		return nil, err
	}
	if m.UniqueTokensCount, err = getUniqueTokensCount(repo.db, brainID); err != nil {
		return nil, err
	}
	if m.ClassFrequency, err = getClassFrequency(repo.db, brainID); err != nil {
		return nil, err
	}
	if m.ClassSize, err = getClassSize(repo.db, brainID); err != nil {
		return nil, err
	}
	if m.TokenFrequency ,err = getTokenFrequency(repo.db, brainID); err != nil {
		return nil, err
	}
	return &m, nil
}

func getSamplesCount(db *gorm.DB, brainID int64) (int64, error) {
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

func getUniqueTokensCount(db *gorm.DB, brainID int64) (int64, error) {
	r := struct{ UniqueTokensCount int64 }{}

	sql := `
		SELECT count(*) as unique_tokens_count FROM (
			SELECT DISTINCT token_id FROM training_tokens WHERE brain_id = ? AND deleted_at IS NULL
		) AS t 
	`

	if err := db.Raw(sql, brainID).Scan(&r).Error; err != nil {
		return 0, err
	}

	return r.UniqueTokensCount, nil
}

func getClassSize (db *gorm.DB, brainID int64) (entity.ClassSizeMap, error) {
	r := entity.ClassSizeMap{}

	rows, err := db.
		Table("training_tokens").
		Select("class_id, SUM(samples_count) AS tokens_count").
		Where("deleted_at IS NULL and brain_id = ?", brainID).
		Group("class_id").
		Rows()


	if err != nil {
		return r, err
	}

	defer rows.Close()

	for rows.Next() {
		classSize := struct {
			ClassID      int64
			TokensCount int64
		}{}
		err = db.ScanRows(rows, &classSize)
		if err != nil {
			return r, err
		}
		r[classSize.ClassID] = classSize.TokensCount
	}

	return r, nil
}

func getTokenFrequency (db *gorm.DB, brainID int64) (result.TokenFrequency, error) {
	r := result.TokenFrequency{}

	rows, err := db.
		Table("training_tokens").
		Select("class_id, token_id, SUM(samples_count) AS count").
		Where("deleted_at IS NULL and brain_id = ?", brainID).
		Group("class_id, token_id").
		Rows()

	if err != nil {
		return r, err
	}

	defer rows.Close()

	for rows.Next() {
		tokenFrequency := struct {
			ClassID      int64
			TokenID      int64
			Count int64
		}{}
		err = db.ScanRows(rows, &tokenFrequency)
		if err != nil {
			return r, err
		}
		if r[tokenFrequency.ClassID] == nil {
			r[tokenFrequency.ClassID] = make(map[int64]int64)
		}
		r[tokenFrequency.ClassID][tokenFrequency.TokenID] = tokenFrequency.Count
	}

	return r, nil
}