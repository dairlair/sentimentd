package service

import (
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type TokenService struct {
	tokenRepository TokenRepositoryInterface
}

func NewTokenService(tokenRepository TokenRepositoryInterface) *TokenService {
	return &TokenService{
		tokenRepository: tokenRepository,
	}
}

func (service TokenService) FindOrCreate (brainID int64, text string) int64 {
	token, err := service.tokenRepository.FindByBrainAndText(brainID, text)

	if err == gorm.ErrRecordNotFound {
		token, err = service.tokenRepository.Create(brainID, text)
	}

	if err != nil {
		log.Fatal(err)
	}

	return token.GetID()
}