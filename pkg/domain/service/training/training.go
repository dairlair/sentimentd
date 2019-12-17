// The training service accepts batch of samples and use them to collect frequencies of classes
// and features (words) in classes.

package training

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/domain/service"
	"github.com/dairlair/sentimentd/pkg/domain/service/training/result"
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

func (service TrainingService) Train(brainID int64, samples []Sample, cb func()) (result.TrainingResult, error) {
	resultsCollector := result.NewTrainingResult()
	for _, sample := range samples {
		for _, label := range sample.Classes {
			class, err := service.classService.FindOrCreate(brainID, label)
			if err != nil {
				return *resultsCollector, err
			}
			resultsCollector.IncClassCount(class.GetID())

			texts := service.tokenizer.Tokenize(sample.Sentence)
			for _, text := range texts {
				token, err := service.tokenService.FindOrCreate(brainID, text)
				if err != nil {
					return *resultsCollector, err
				}
				resultsCollector.IncTokenCount(class.GetID(), token.GetID())
			}

			resultsCollector.IncSamplesCount()
		}
		cb()
	}

	return *resultsCollector, nil
}
