// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	model "github.com/bersennaidoo/eopd/registration-service/domain/model"
	mock "github.com/stretchr/testify/mock"
)

// Storer is an autogenerated mock type for the Storer type
type Storer struct {
	mock.Mock
}

// Register provides a mock function with given fields: registration
func (_m *Storer) Register(registration *model.RegistrationRequest) error {
	ret := _m.Called(registration)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.RegistrationRequest) error); ok {
		r0 = rf(registration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: registration
func (_m *Storer) Update(registration *model.RegistrationRequest) error {
	ret := _m.Called(registration)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.RegistrationRequest) error); ok {
		r0 = rf(registration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// View provides a mock function with given fields: patientID
func (_m *Storer) View(patientID string) (*model.RegistrationRequest, error) {
	ret := _m.Called(patientID)

	var r0 *model.RegistrationRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.RegistrationRequest, error)); ok {
		return rf(patientID)
	}
	if rf, ok := ret.Get(0).(func(string) *model.RegistrationRequest); ok {
		r0 = rf(patientID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RegistrationRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(patientID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStorer creates a new instance of Storer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Storer {
	mock := &Storer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
