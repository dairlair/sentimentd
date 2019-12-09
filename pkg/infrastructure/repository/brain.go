package repository

import (
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
	. "github.com/dairlair/sentimentd/pkg/infrastructure/model"
	"github.com/jinzhu/gorm"
)

type BrainRepository struct {
	repository
}

func NewBrainRepository(db *gorm.DB) BrainRepositoryInterface {
	return &BrainRepository{
		repository: repository{
			db:db,
		},
	}
}

func (repo *BrainRepository) GetAll() ([]BrainInterface, error) {
	var brains = []Brain{}
	repo.repository.db.Find(&brains)

	// @See https://github.com/golang/go/wiki/InterfaceSlice#what-can-i-do-instead
	brainsInterfaces := make([]BrainInterface, len(brains))

	x := len(brains)
	fmt.Printf("%d brains retrieved\n", x)

	return brainsInterfaces, nil
}