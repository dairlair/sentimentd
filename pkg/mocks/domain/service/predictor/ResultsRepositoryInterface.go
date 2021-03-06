// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	classifier "github.com/dairlair/sentimentd/pkg/domain/service/classifier"
	mock "github.com/stretchr/testify/mock"
)

// ResultsRepositoryInterface is an autogenerated mock type for the ResultsRepositoryInterface type
type ResultsRepositoryInterface struct {
	mock.Mock
}

// GetTrainedModel provides a mock function with given fields: brainID
func (_m *ResultsRepositoryInterface) GetTrainedModel(brainID int64) (classifier.TrainedModelInterface, error) {
	ret := _m.Called(brainID)

	var r0 classifier.TrainedModelInterface
	if rf, ok := ret.Get(0).(func(int64) classifier.TrainedModelInterface); ok {
		r0 = rf(brainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(classifier.TrainedModelInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(brainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
