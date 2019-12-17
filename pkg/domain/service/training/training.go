// The training service accepts batch of samples and use them to collect frequencies of classes
// and features (words) in classes.

package training

import (
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/domain/service"
	"github.com/dairlair/sentimentd/pkg/domain/service/training/result"
	"strings"
)

type TokenizerInterface interface {
	Tokenize(string) []string
}

type TrainingService struct {
	tokenizer    TokenizerInterface
	tokenService service.TokenServiceInterface
	classService service.ClassServiceInterface
}

func NewTrainingService(
	tokenizer TokenizerInterface,
	tokenService service.TokenServiceInterface,
	classService service.ClassServiceInterface,
	) *TrainingService {
	return &TrainingService{
		tokenizer:    tokenizer,
		tokenService: tokenService,
		classService: classService,
	}
}

func (service TrainingService) Train (brainID int64, samples []Sample) (result.TrainingResult, error) {
	resultsCollector := result.NewTrainingResult()
	fmt.Printf("Train Brain #%d\n", brainID)
	for _, sample := range samples {
		fmt.Printf("Sample: \n")
		fmt.Printf("  sentence: %s\n", sample.Sentence)
		fmt.Printf("  classes: %s\n", strings.Join(sample.Classes, ", "))
		for _, label := range sample.Classes {
			class, err := service.classService.FindOrCreate(brainID, label)
			if err != nil {
				return *resultsCollector, err
			}
			fmt.Printf("class retrieved successfully: %d#%s from %s\n", class.GetID(), class.GetLabel(), label)
			resultsCollector.IncClassCount(class.GetID())

			texts := service.tokenizer.Tokenize(sample.Sentence)
			for _, text := range texts {
				token, err := service.tokenService.FindOrCreate(brainID, text)
				if err != nil {
					return *resultsCollector, err
				}
				fmt.Printf("token retrieved successfully: %d#%s from %s\n", class.GetID(), token.GetText(), text)
				resultsCollector.IncTokenCount(class.GetID(), token.GetID())
			}

			resultsCollector.IncSamplesCount()
		}
	}

	return *resultsCollector, nil
}