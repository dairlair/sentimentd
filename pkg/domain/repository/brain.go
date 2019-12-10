package repository

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
)

type BrainRepositoryInterface interface {
	GetAll() ([]BrainInterface, error)
	Create(name string, description string) (BrainInterface, error)
}