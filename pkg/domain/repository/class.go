// This repository helps us to work with classes.
package repository

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
)

type ClassRepositoryInterface interface {
	//GetAllByBrainID(brainID int64) ([]ClassInterface, error)
	//GetByBrainIDAndLabel(brainID int64, label string) (ClassInterface, error)
	Create(brainID int64, label string) (ClassInterface, error)
	FindByBrainAndLabel(brainID int64, label string) (ClassInterface, error)
}