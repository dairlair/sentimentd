// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/dairlair/sentimentd/pkg/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// ClassRepositoryInterface is an autogenerated mock type for the ClassRepositoryInterface type
type ClassRepositoryInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: brainID, label
func (_m *ClassRepositoryInterface) Create(brainID int64, label string) (entity.ClassInterface, error) {
	ret := _m.Called(brainID, label)

	var r0 entity.ClassInterface
	if rf, ok := ret.Get(0).(func(int64, string) entity.ClassInterface); ok {
		r0 = rf(brainID, label)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.ClassInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(brainID, label)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByBrainAndLabel provides a mock function with given fields: brainID, label
func (_m *ClassRepositoryInterface) FindByBrainAndLabel(brainID int64, label string) (entity.ClassInterface, error) {
	ret := _m.Called(brainID, label)

	var r0 entity.ClassInterface
	if rf, ok := ret.Get(0).(func(int64, string) entity.ClassInterface); ok {
		r0 = rf(brainID, label)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entity.ClassInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(brainID, label)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
