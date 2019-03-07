// Code generated by mockery v1.0.0
package mocks

import io "io"
import mock "github.com/stretchr/testify/mock"
import models "github.com/MTES-MCT/filharmonic-api/models"

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *Storage) Get(id string) (io.Reader, error) {
	ret := _m.Called(id)

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func(string) io.Reader); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: file
func (_m *Storage) Put(file models.File) (string, error) {
	ret := _m.Called(file)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.File) string); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.File) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
