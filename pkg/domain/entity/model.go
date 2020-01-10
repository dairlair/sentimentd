package entity

import (
	"github.com/dairlair/sentimentd/pkg/domain/service/training/result"
)

type TrainedModel struct {
	SamplesCount int64
	UniqueTokensCount int64
	ClassFrequency result.ClassFrequency
}

func (m TrainedModel) GetSamplesCount() int64 {
	return m.SamplesCount
}

func (m TrainedModel) GetUniqueTokensCount() int64 {
	return m.UniqueTokensCount
}

func (m TrainedModel) GetClassFrequency() result.ClassFrequency {
	return m.ClassFrequency
}