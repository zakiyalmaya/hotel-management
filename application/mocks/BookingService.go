// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/zakiyalmaya/hotel-management/model"
)

// BookingService is an autogenerated mock type for the BookingService type
type BookingService struct {
	mock.Mock
}

// Books provides a mock function with given fields: _a0
func (_m *BookingService) Books(_a0 *model.BookingEntity) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Books")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.BookingEntity) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByRegisterNumber provides a mock function with given fields: registerNumber
func (_m *BookingService) GetByRegisterNumber(registerNumber string) (*model.BookingResponse, error) {
	ret := _m.Called(registerNumber)

	if len(ret) == 0 {
		panic("no return value specified for GetByRegisterNumber")
	}

	var r0 *model.BookingResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.BookingResponse, error)); ok {
		return rf(registerNumber)
	}
	if rf, ok := ret.Get(0).(func(string) *model.BookingResponse); ok {
		r0 = rf(registerNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BookingResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(registerNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reschedule provides a mock function with given fields: _a0
func (_m *BookingService) Reschedule(_a0 *model.BookingEntity) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Reschedule")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.BookingEntity) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePayment provides a mock function with given fields: _a0
func (_m *BookingService) UpdatePayment(_a0 *model.BookingEntity) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePayment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.BookingEntity) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBookingService creates a new instance of BookingService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBookingService(t interface {
	mock.TestingT
	Cleanup(func())
}) *BookingService {
	mock := &BookingService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
