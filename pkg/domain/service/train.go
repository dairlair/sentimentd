package service

import (
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
	"strings"
)

type TrainService struct {
	classRepository ClassRepositoryInterface
}

func NewTrainService(classRepository ClassRepositoryInterface) *TrainService {
	return &TrainService{
		classRepository: classRepository,
	}
}

func (service TrainService) Train (brainID int64, samples []Sample) error {
	fmt.Printf("Train Brain #%d\n", brainID)
	for _, sample := range samples {
		fmt.Printf("Sample: \n")
		fmt.Printf("  sentence: %s\n", sample.Sentence)
		fmt.Printf("  classes: %s\n", strings.Join(sample.Classes, ", "))
		for _, class := range sample.Classes {
			class, err := service.classRepository.Create(brainID, class)
			if err != nil {
				fmt.Printf("class creation error: %s\n", err)
				continue
			}

			fmt.Printf("class created successfully: %d\n", class.GetID())
		}
	}

	return nil
}