// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	user "github.com/mkrashad/go-todo/user/internal"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0
func (_m *Service) CreateUser(_a0 user.User) (user.User, error) {
	ret := _m.Called(_a0)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(user.User) (user.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(user.User) user.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(user.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserById provides a mock function with given fields: id
func (_m *Service) DeleteUserById(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUsers provides a mock function with given fields:
func (_m *Service) GetAllUsers() []user.User {
	ret := _m.Called()

	var r0 []user.User
	if rf, ok := ret.Get(0).(func() []user.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	return r0
}

// GetUserById provides a mock function with given fields: id
func (_m *Service) GetUserById(id uint64) (user.User, error) {
	ret := _m.Called(id)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (user.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) user.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserById provides a mock function with given fields: id, updatedUser
func (_m *Service) UpdateUserById(id uint64, updatedUser user.User) (user.User, error) {
	ret := _m.Called(id, updatedUser)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, user.User) (user.User, error)); ok {
		return rf(id, updatedUser)
	}
	if rf, ok := ret.Get(0).(func(uint64, user.User) user.User); ok {
		r0 = rf(id, updatedUser)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(uint64, user.User) error); ok {
		r1 = rf(id, updatedUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}