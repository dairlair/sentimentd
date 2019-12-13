// This repository helps us to work with tokens.
package repository

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
)

type TokenRepositoryInterface interface {
	Create(brainID int64, text string) (TokenInterface, error)
	FindByBrainAndText(brainID int64, text string) (TokenInterface, error)
}