package repository

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
)

type BrainRepositoryInterface interface {
	GetAll() ([]BrainInterface, error)
	GetByID(id int64) (BrainInterface, error)
	Create(name string, description string) (BrainInterface, error)
	Delete(id int64) error
}