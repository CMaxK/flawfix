// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	uuid "github.com/google/uuid"
	models "github.com/l3montree-dev/flawfix/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// organizationRepository is an autogenerated mock type for the organizationRepository type
type organizationRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: tx, t
func (_m *organizationRepository) Create(tx interface{}, t *models.Organization) error {
	ret := _m.Called(tx, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, *models.Organization) error); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: tx, id
func (_m *organizationRepository) Delete(tx interface{}, id uuid.UUID) error {
	ret := _m.Called(tx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, uuid.UUID) error); ok {
		r0 = rf(tx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: ids
func (_m *organizationRepository) List(ids []uuid.UUID) ([]models.Organization, error) {
	ret := _m.Called(ids)

	var r0 []models.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func([]uuid.UUID) ([]models.Organization, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]uuid.UUID) []models.Organization); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Organization)
		}
	}

	if rf, ok := ret.Get(1).(func([]uuid.UUID) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields: id
func (_m *organizationRepository) Read(id uuid.UUID) (models.Organization, error) {
	ret := _m.Called(id)

	var r0 models.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (models.Organization, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) models.Organization); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Organization)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transaction provides a mock function with given fields: _a0
func (_m *organizationRepository) Transaction(_a0 func(interface{}) error) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(interface{}) error) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: tx, t
func (_m *organizationRepository) Update(tx interface{}, t *models.Organization) error {
	ret := _m.Called(tx, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, *models.Organization) error); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newOrganizationRepository creates a new instance of organizationRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newOrganizationRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *organizationRepository {
	mock := &organizationRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}