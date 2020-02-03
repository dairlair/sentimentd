// The training service accepts batch of samples and use them to collect frequencies of classes
// and features (words) in classes.

package training

import (
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/domain/service/training/result"
)

// TokenizerInterface defines dependency
type TokenizerInterface interface {
	Tokenize(string) []string
}

// ResultsRepositoryInterface defines dependency
type ResultsRepositoryInterface interface {
	SaveResult(brainID int64, res result.TrainingResult) error
}

// ClassServiceInterface defines dependency
type ClassServiceInterface interface {
	FindOrCreate(brainID int64, label string) (entity.ClassInterface, error)
}

// TokenServiceInterface defines dependency
type TokenServiceInterface interface {
	FindOrCreate(brainID int64, text string) (entity.TokenInterface, error)
}

// Service provides functionality to brain's training
type Service struct {
	classService      ClassServiceInterface
	resultsRepository ResultsRepositoryInterface
	tokenizer         TokenizerInterface
	tokenService      TokenServiceInterface
}

// NewTrainingService is a constructor for TrainingService
func NewTrainingService(
	classService ClassServiceInterface,
	resultsRepository ResultsRepositoryInterface,
	tokenizer TokenizerInterface,
	tokenService TokenServiceInterface,
) *Service {
	return &Service{
		classService:      classService,
		resultsRepository: resultsRepository,
		tokenizer:         tokenizer,
		tokenService:      tokenService,
	}
}

// Train accepts samples and returns the training result
func (service Service) Train(brainID int64, samples []entity.Sample, cb func()) error {
	trainingResult := result.NewTrainingResult()
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

	if err := service.saveTrainingResult(brainID, *trainingResult); err != nil {
		return err
	}

	return nil
}

// This method saves the collected frequencies into the database
func (service Service) saveTrainingResult(brainID int64, trainingResult result.TrainingResult) error {
	return service.resultsRepository.SaveResult(brainID, trainingResult)
}
