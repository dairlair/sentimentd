// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/dairlair/sentimentd/pkg/domain/entity"
	mock "github.com/stretchr/testify/mock"

	result "github.com/dairlair/sentimentd/pkg/domain/service/training/result"
)

// TrainedModelInterface is an autogenerated mock type for the TrainedModelInterface type
type TrainedModelInterface struct {
	mock.Mock
}

// GetClassFrequency provides a mock function with given fields:
func (_m *TrainedModelInterface) GetClassFrequency() result.ClassFrequency {
	ret := _m.Called()

	var r0 result.ClassFrequency
	if rf, ok := ret.Get(0).(func() result.ClassFrequency); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(result.ClassFrequency)
		}
	}

	return r0
}

// GetClassSizes provides a mock function with given fields:
func (_m *TrainedModelInterface) GetClassSizes() entity.ClassSizeMap {
	ret := _m.Called()

	var r0 entity.ClassSizeMap
	if rf, ok := ret.Get(0).(func() entity.ClassSizeMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.ClassSizeMap)
		}
	}

	return r0
}

// GetSamplesCount provides a mock function with given fields:
func (_m *TrainedModelInterface) GetSamplesCount() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetTokenFrequency provides a mock function with given fields:
func (_m *TrainedModelInterface) GetTokenFrequency() result.TokenFrequency {
	ret := _m.Called()

	var r0 result.TokenFrequency
	if rf, ok := ret.Get(0).(func() result.TokenFrequency); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(result.TokenFrequency)
		}
	}

	return r0
}

// GetUniqueTokensCount provides a mock function with given fields:
func (_m *TrainedModelInterface) GetUniqueTokensCount() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}