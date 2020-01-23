package training

import (
	"fmt"
	"github.com/dairlair/sentimentd/pkg/domain/entity"
	"github.com/dairlair/sentimentd/pkg/domain/service/tokenizer"
	"github.com/dairlair/sentimentd/pkg/infrastructure/model"
	mocks "github.com/dairlair/sentimentd/pkg/mocks/domain/service/training"
	"github.com/lukechampine/randmap/safe"
	"github.com/stretchr/testify/mock"
	"github.com/tjarratt/babble"
	"testing"
	"time"
)

func BenchmarkTrainingService_Train(b *testing.B) {
	brainId := int64(1)
	classes := generateClasses(10)
	samples := generateSamples(classes, 1000)

	classService := mocks.ClassServiceInterface{}
	class := &model.Class{
		Model:   model.Model{ID:777},
		BrainID: 1,
		Label:   "qwerty",
	}
	classService.On("FindOrCreate", mock.AnythingOfType("int64"), mock.AnythingOfType("string")).Return(class, nil)

	resultRepository := mocks.ResultsRepositoryInterface{}
	resultRepository.On("SaveResult", brainId, mock.AnythingOfType("result.TrainingResult")).Return(nil)

	tokenizerService := tokenizer.NewTokenizer()


	tokenService := mocks.TokenServiceInterface{}
	token := &model.Token{
		Model:   model.Model{
			ID:        1,
			CreatedAt: time.Time{},
			DeletedAt: nil,
		},
		BrainID: 1,
		Text:    "text",
	}
	tokenService.On("FindOrCreate", mock.AnythingOfType("int64"), mock.AnythingOfType("string")).Return(token, nil)

	service := NewTrainingService(&classService, &resultRepository, &tokenizerService, &tokenService)

	_ = service.Train(brainId, samples, func() {})
}

func generateClasses(count int64) map[string]entity.ClassInterface {
	classes := make(map[string]entity.ClassInterface, count)

	for i := int64(0); i < count; i++ {
		label := fmt.Sprintf("class_%d", i)
		class := &model.Class{
			Model:   model.Model{ID:count},
			BrainID: 1,
			Label:   label,
		}
		classes[label] = class
	}

	return classes
}

func generateSamples(classes map[string]entity.ClassInterface, count int64) []entity.Sample {
	samples := make([]entity.Sample, count)
	babbler := babble.NewBabbler()
	babbler.Count = 25
	babbler.Separator = " "

	for i := int64(0); i < count; i++ {
		class := randmap.Val(classes).(entity.ClassInterface)
		sample := entity.Sample{
			Sentence: babbler.Babble(),
			Classes:  []string{class.GetLabel()},
		}
		samples[i] = sample
	}

	return samples
}