package repository

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
	. "github.com/dairlair/sentimentd/pkg/infrastructure/model"
	"github.com/jinzhu/gorm"
)

type TokenRepository struct {
	repository
}

func NewTokenRepository(db *gorm.DB) TokenRepositoryInterface {
	return &TokenRepository{
		repository: repository{
			db: db,
		},
	}
}

func (repo *TokenRepository) Create(brainID int64, text string) (TokenInterface, error) {
	token := Token{BrainID: brainID, Text:text}
	if err := repo.repository.db.Create(&token).Error; err != nil {
		return nil, err
	}

	return &token, nil
}

func (repo *TokenRepository) FindByBrainAndText(brainID int64, text string) (TokenInterface, error) {
	token := Token{BrainID: brainID, Text:text}
	if err := repo.repository.db.Where(token).First(&token).Error; err != nil {
		return nil, err
	}

	return &token, nil
}