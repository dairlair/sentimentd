// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/dairlair/sentimentd/pkg/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// BrainRepositoryInterface is an autogenerated mock type for the BrainRepositoryInterface type
type BrainRepositoryInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: name, description
func (_m *BrainRepositoryInterface) Create(name string, description string) (entity.BrainInterface, error) {
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

// Delete provides a mock function with given fields: id
func (_m *BrainRepositoryInterface) Delete(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *BrainRepositoryInterface) GetAll() ([]entity.BrainInterface, error) {
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

// GetByID provides a mock function with given fields: id
func (_m *BrainRepositoryInterface) GetByID(id int64) (entity.BrainInterface, error) {
	ret := _m.Called(id)

	var r0 entity.BrainInterface
	if rf, ok := ret.Get(0).(func(int64) entity.BrainInterface); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.BrainInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
