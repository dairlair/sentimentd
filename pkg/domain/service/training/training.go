// The training service accepts batch of samples and use them to collect frequencies of classes
// and features (words) in classes.

package training

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
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
	FindOrCreate(brainID int64, text string) int64
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
		tokenService:      NewTextServiceCache(tokenService),
	}
}

// Train accepts samples and returns the training result
func (service Service) Train(brainID int64, samples []entity.Sample) error {
	textTrainingResult := service.textTrain(samples)
	trainingResult := result.NewTrainingResult()

	// Copy the samples count
	trainingResult.SamplesCount = textTrainingResult.SamplesCount

	// Copy the classes frequencies
	fmt.Println("Class frequencies preparing")
	bar := pb.StartNew(len(textTrainingResult.TextClassFrequency))
	var classLabelIDMap = make(map[string]int64, len(textTrainingResult.TextClassFrequency))
	for classLabel, count := range textTrainingResult.TextClassFrequency {
		class, err := service.classService.FindOrCreate(brainID, classLabel)
		if err != nil {
			return err
		}
		classLabelIDMap[classLabel] = class.GetID()
		trainingResult.ClassFrequency[class.GetID()] = count
		bar.Increment()
	}
	bar.Finish()

	// Copy the tokens frequencies
	for classLabel, tokenFrequency := range textTrainingResult.TextTokenFrequency {
		classID := classLabelIDMap[classLabel]
		fmt.Printf("Save tokens for class [%s]\n", classLabel)
		bar := pb.StartNew(len(tokenFrequency))
		trainingResult.TokenFrequency[classID] = make(map[int64]int64, 1)
		for tokenText, count := range tokenFrequency {
			tokenID := service.tokenService.FindOrCreate(brainID, tokenText)
			trainingResult.TokenFrequency[classID][tokenID] = count
			bar.Increment()
		}
		bar.Finish()
	}

	if err := service.saveTrainingResult(brainID, *trainingResult); err != nil {
		return err
	}

	return nil
}

// Train accepts samples and returns the training result
func (service Service) textTrain(samples []entity.Sample) *result.TextTrainingResult {
	textTrainingResult := result.NewTextTrainingResult()
	fmt.Println("The text train")
	bar := pb.StartNew(len(samples))
	for _, sample := range samples {
		for _, classLabel := range sample.Classes {
			textTrainingResult.IncClassCount(classLabel)
			texts := service.tokenizer.Tokenize(sample.Sentence)
			for _, tokenText := range texts {
				textTrainingResult.IncTokenCount(classLabel, tokenText)
			}
			textTrainingResult.IncSamplesCount()
		}
		bar.Increment()
	}
	bar.Finish()
	return textTrainingResult
}

// This method saves the collected frequencies into the database
func (service Service) saveTrainingResult(brainID int64, trainingResult result.TrainingResult) error {
	return service.resultsRepository.SaveResult(brainID, trainingResult)
}
