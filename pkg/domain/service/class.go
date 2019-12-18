package service

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
	"github.com/jinzhu/gorm"
)

type ClassService struct {
	classRepository ClassRepositoryInterface
}

func NewClassService(classRepository ClassRepositoryInterface) *ClassService {
	return &ClassService{
		classRepository: classRepository,
	}
}

func (service ClassService) FindOrCreate (brainID int64, label string) (ClassInterface, error) {
	class, err := service.classRepository.FindByBrainAndLabel(brainID, label)

	if err == gorm.ErrRecordNotFound {
		return service.classRepository.Create(brainID, label)
	}

	if err != nil {
		return nil, err
	}

	return class, nil
}