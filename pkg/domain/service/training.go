// The training service accepts batch of samples and use them to collect frequencies of classes
// and features (words) in classes.

package service

import (
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	"strings"
)

type TrainingService struct {
	classService ClassServiceInterface
}

func NewTrainingService(classService ClassServiceInterface) *TrainingService {
	return &TrainingService{
		classService: classService,
	}
}

func (service TrainingService) Train (brainID int64, samples []Sample) error {
	fmt.Printf("Train Brain #%d\n", brainID)
	for _, sample := range samples {
		fmt.Printf("Sample: \n")
		fmt.Printf("  sentence: %s\n", sample.Sentence)
		fmt.Printf("  classes: %s\n", strings.Join(sample.Classes, ", "))
		//for _, class := range sample.Classes {
		//	class, err := service.classRepository.Create(brainID, class)
		//	if err != nil {
		//		fmt.Printf("class creation error: %s\n", err)
		//		continue
		//	}
		//
		//	//fmt.Printf("class created successfully: %d\n", class.GetID())
		//}
	}

	return nil
}