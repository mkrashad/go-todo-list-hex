// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	task "github.com/mkrashad/go-todo/internal/task"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: _a0
func (_m *Repository) CreateTask(_a0 task.Task) (task.Task, error) {
	ret := _m.Called(_a0)

	var r0 task.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(task.Task) (task.Task, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(task.Task) task.Task); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(task.Task)
	}

	if rf, ok := ret.Get(1).(func(task.Task) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTaskById provides a mock function with given fields: id
func (_m *Repository) DeleteTaskById(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTasks provides a mock function with given fields:
func (_m *Repository) GetAllTasks() []task.Task {
	ret := _m.Called()

	var r0 []task.Task
	if rf, ok := ret.Get(0).(func() []task.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]task.Task)
		}
	}

	return r0
}

// GetTaskById provides a mock function with given fields: id
func (_m *Repository) GetTaskById(id uint64) (task.Task, error) {
	ret := _m.Called(id)

	var r0 task.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (task.Task, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) task.Task); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(task.Task)
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTaskById provides a mock function with given fields: id, _a1
func (_m *Repository) UpdateTaskById(id uint64, _a1 task.Task) (task.Task, error) {
	ret := _m.Called(id, _a1)

	var r0 task.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, task.Task) (task.Task, error)); ok {
		return rf(id, _a1)
	}
	if rf, ok := ret.Get(0).(func(uint64, task.Task) task.Task); ok {
		r0 = rf(id, _a1)
	} else {
		r0 = ret.Get(0).(task.Task)
	}

	if rf, ok := ret.Get(1).(func(uint64, task.Task) error); ok {
		r1 = rf(id, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}