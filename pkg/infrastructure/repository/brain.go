package repository

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
)

type BrainRepository struct {
}

func NewBrainRepository() BrainRepositoryInterface {
	return &BrainRepository{}
}

func (repo *BrainRepository) GetAll() ([]BrainInterface, error) {
	brains := make([]BrainInterface, 1)

	return brains, nil
}