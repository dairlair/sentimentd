package service

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
	"github.com/jinzhu/gorm"
)

type TokenService struct {
	tokenRepository TokenRepositoryInterface
}

type TokenServiceInterface interface {
	FindOrCreate (brainID int64, text string) (TokenInterface, error)
}

func NewTokenService(tokenRepository TokenRepositoryInterface) *TokenService {
	return &TokenService{
		tokenRepository: tokenRepository,
	}
}

func (service TokenService) FindOrCreate (brainID int64, text string) (TokenInterface, error) {
	token, err := service.tokenRepository.FindByBrainAndText(brainID, text)

	if err == gorm.ErrRecordNotFound {
		return service.tokenRepository.Create(brainID, text)
	}

	if err != nil {
		return nil, err
	}

	return token, nil
}