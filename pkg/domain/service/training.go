// The training service accepts batch of samples and use them to collect frequencies of classes
// and features (words) in classes.

package service

import (
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	"strings"
)

type TokenizerInterface interface {
	Tokenize(string) []string
}

type TrainingService struct {
	tokenizer TokenizerInterface
	classService ClassServiceInterface
}

func NewTrainingService(tokenizer TokenizerInterface, classService ClassServiceInterface) *TrainingService {
	return &TrainingService{
		tokenizer: tokenizer,
		classService: classService,
	}
}

func (service TrainingService) Train (brainID int64, samples []Sample) (error, trainingResul) {
	fmt.Printf("Train Brain #%d\n", brainID)
	for _, sample := range samples {
		fmt.Printf("Sample: \n")
		fmt.Printf("  sentence: %s\n", sample.Sentence)
		fmt.Printf("  classes: %s\n", strings.Join(sample.Classes, ", "))
		for _, label := range sample.Classes {
			class, err := service.classService.FindOrCreate(brainID, label)
			if err != nil {
				fmt.Printf("class getting error: %s\n", err)
				continue
			}

			tokens := service.tokenizer.Tokenize(sample.Sentence)
			for _, token := range tokens {

			}

			fmt.Printf("class retrieved successfully: %s#%d from %s\n", class.GetLabel(), class.GetID(), label)
		//	class, err := service.classRepository.Create(brainID, class)
		//	if err != nil {
		//		fmt.Printf("class creation error: %s\n", err)
		//		continue
		//	}
		//
		//	//fmt.Printf("class created successfully: %d\n", class.GetID())
		}
	}

	return nil
}