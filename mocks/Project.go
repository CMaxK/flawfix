// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// Project is an autogenerated mock type for the Project type
type Project struct {
	mock.Mock
}

// GetID provides a mock function with given fields:
func (_m *Project) GetID() uuid.UUID {
	ret := _m.Called()

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// NewProject creates a new instance of Project. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProject(t interface {
	mock.TestingT
	Cleanup(func())
}) *Project {
	mock := &Project{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
