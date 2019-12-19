// The training service accepts batch of samples and use them to collect frequencies of classes
// and features (words) in classes.

package training

import (
	. "github.com/dairlair/sentimentd/pkg/domain/entity"
	. "github.com/dairlair/sentimentd/pkg/domain/service/training/result"
)

type TokenizerInterface interface {
	Tokenize(string) []string
}

type ResultsRepositoryInterface interface {
	SaveResult(result TrainingResult) error
}

type ClassServiceInterface interface {
	FindOrCreate (brainID int64, label string) (ClassInterface, error)
}

type TokenServiceInterface interface {
	FindOrCreate (brainID int64, text string) (TokenInterface, error)
}

type TrainingService struct {
	classService ClassServiceInterface
	resultsRepository ResultsRepositoryInterface
	tokenizer    TokenizerInterface
	tokenService TokenServiceInterface
}

func NewTrainingService(
	classService ClassServiceInterface,
	resultsRepository ResultsRepositoryInterface,
	tokenizer TokenizerInterface,
	tokenService TokenServiceInterface,
) *TrainingService {
	return &TrainingService{
		classService: classService,
		resultsRepository: resultsRepository,
		tokenizer:    tokenizer,
		tokenService: tokenService,
	}
}

func (service TrainingService) Train(brainID int64, samples []Sample, cb func()) error {
	trainingResult := NewTrainingResult()
	for _, sample := range samples {
		for _, label := range sample.Classes {
			class, err := service.classService.FindOrCreate(brainID, label)
			if err != nil {
				return err
			}
			trainingResult.IncClassCount(class.GetID())

			texts := service.tokenizer.Tokenize(sample.Sentence)
			for _, text := range texts {
				token, err := service.tokenService.FindOrCreate(brainID, text)
				if err != nil {
					return err
				}
				trainingResult.IncTokenCount(class.GetID(), token.GetID())
			}

			trainingResult.IncSamplesCount()
		}
		cb()
	}

	if err := service.saveTrainingResult(*trainingResult); err != nil {
		return err
	}

	return nil
}

// This method saves the collected frequencies into the database
func (service TrainingService) saveTrainingResult(trainingResult TrainingResult) error {
	return service.resultsRepository.SaveResult(trainingResult)
}
