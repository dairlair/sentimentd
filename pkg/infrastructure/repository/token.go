package repository

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	. "github.com/dairlair/sentimentd/pkg/infrastructure/model"
	"github.com/jinzhu/gorm"
)

type TokenRepository struct {
	repository
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
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

func (repo *TokenRepository) GetTokenIDs(brainID int64, tokens []string) ([]int64, error) {
	models := make([]Token, len(tokens))

	if err := repo.db.Where("brain_id = ? AND text IN (?)", brainID, tokens).Find(&models).Error; err != nil {
		return []int64{}, err
	}

	var IDs []int64
	for _, token := range models {
		IDs = append(IDs, token.ID)
	}

	return IDs, nil
}