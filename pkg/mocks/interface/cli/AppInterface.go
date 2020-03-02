// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/dairlair/sentimentd/pkg/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// AppInterface is an autogenerated mock type for the AppInterface type
type AppInterface struct {
	mock.Mock
}

// BrainList provides a mock function with given fields:
func (_m *AppInterface) BrainList() ([]entity.BrainInterface, error) {
	ret := _m.Called()

	var r0 []entity.BrainInterface
	if rf, ok := ret.Get(0).(func() []entity.BrainInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.BrainInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBrain provides a mock function with given fields: name, description
func (_m *AppInterface) CreateBrain(name string, description string) (entity.BrainInterface, error) {
	ret := _m.Called(name, description)

	var r0 entity.BrainInterface
	if rf, ok := ret.Get(0).(func(string, string) entity.BrainInterface); ok {
		r0 = rf(name, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.BrainInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBrain provides a mock function with given fields: id
func (_m *AppInterface) DeleteBrain(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBrainByReference provides a mock function with given fields: reference
func (_m *AppInterface) GetBrainByReference(reference string) (entity.BrainInterface, error) {
	ret := _m.Called(reference)

	var r0 entity.BrainInterface
	if rf, ok := ret.Get(0).(func(string) entity.BrainInterface); ok {
		r0 = rf(reference)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.BrainInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClassByID provides a mock function with given fields: classID
func (_m *AppInterface) GetClassByID(classID int64) (entity.ClassInterface, error) {
	ret := _m.Called(classID)

	var r0 entity.ClassInterface
	if rf, ok := ret.Get(0).(func(int64) entity.ClassInterface); ok {
		r0 = rf(classID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.ClassInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(classID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HumanizedPredict provides a mock function with given fields: brainID, text
func (_m *AppInterface) HumanizedPredict(brainID int64, text string) (entity.HumanizedPrediction, error) {
	ret := _m.Called(brainID, text)

	var r0 entity.HumanizedPrediction
	if rf, ok := ret.Get(0).(func(int64, string) entity.HumanizedPrediction); ok {
		r0 = rf(brainID, text)
	} else {
		r0 = ret.Get(0).(entity.HumanizedPrediction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(brainID, text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Predict provides a mock function with given fields: brainID, text
func (_m *AppInterface) Predict(brainID int64, text string) (entity.Prediction, error) {
	ret := _m.Called(brainID, text)

	var r0 entity.Prediction
	if rf, ok := ret.Get(0).(func(int64, string) entity.Prediction); ok {
		r0 = rf(brainID, text)
	} else {
		r0 = ret.Get(0).(entity.Prediction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(brainID, text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Train provides a mock function with given fields: brainID, samples
func (_m *AppInterface) Train(brainID int64, samples []entity.Sample) error {
	ret := _m.Called(brainID, samples)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, []entity.Sample) error); ok {
		r0 = rf(brainID, samples)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
