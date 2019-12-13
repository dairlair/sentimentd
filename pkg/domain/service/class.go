// The training service accepts batch of samples and use them to collect frequencies of classes
// and features (words) in classes.

package service

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
)

type ClassService struct {
	classRepository ClassRepositoryInterface
}

type ClassServiceInterface interface {
	FindOrCreate (brainID int64, label string) (ClassInterface, error)
}

func NewClassService(classRepository ClassRepositoryInterface) *ClassService {
	return &ClassService{
		classRepository: classRepository,
	}
}

func (service ClassService) FindOrCreate (brainID int64, label string) (ClassInterface, error) {
	class, err := service.classRepository.FindByBrainAndLabel(brainID, label)
	if err != nil {
		return nil, err
	}

	return class, nil
}