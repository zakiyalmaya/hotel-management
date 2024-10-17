package guest

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository/mocks"
	"github.com/zakiyalmaya/hotel-management/model"
)

var (
	mockGuest *mocks.GuestRepository
)

func Test_Create(t *testing.T) {
	mockGuest = new(mocks.GuestRepository)

	testCases := []struct {
		name     string
		guest    *model.GuestEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when create guest then return success response",
			guest: &model.GuestEntity{
				FirstName:   "first",
				LastName:    "last",
				Identity:    "identity",
				DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				PhoneNumber: "phone",
				Email:       "email",
			},
			mockCall: func() {
				mockGuest.On("Create", &model.GuestEntity{
					FirstName:   "first",
					LastName:    "last",
					Identity:    "identity",
					Email:       "email",
					PhoneNumber: "phone",
					DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				}).Return(nil).Once()
			},
			err: nil,
		},
		{
			name: "Given error when create guest then retrun error",
			guest: &model.GuestEntity{
				FirstName:   "first",
				LastName:    "last",
				Identity:    "identity",
				DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				PhoneNumber: "phone",
				Email:       "email",
			},
			mockCall: func() {
				mockGuest.On("Create", &model.GuestEntity{
					FirstName:   "first",
					LastName:    "last",
					Identity:    "identity",
					Email:       "email",
					PhoneNumber: "phone",
					DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				}).Return(errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewGuestServiceImpl(&repository.Repositories{GuestRepo: mockGuest})
			err := service.Create(tc.guest)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_GetByID(t *testing.T) {
	mockGuest = new(mocks.GuestRepository)

	testCases := []struct {
		name     string
		id       int
		mockCall func()
		response *model.GuestResponse
		err      error
	}{
		{
			name: "Given valid request when get guest by id then return success response",
			id:   1,
			mockCall: func() {
				mockGuest.On("GetByID", 1).Return(&model.GuestEntity{
					ID:          1,
					FirstName:   "John",
					LastName:    "Doe",
					Identity:    "123",
					DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
					PhoneNumber: "123",
					Email:       "example@mail.com",
				}, nil).Once()
			},
			response: &model.GuestResponse{
				ID:          1,
				Name:        "John Doe",
				Identity:    "123",
				DateOfBirth: "01 January 2000",
				PhoneNumber: "123",
				Email:       "example@mail.com",
			},
			err: nil,
		},
		{
			name: "Given error when get guest by id then return error",
			id:   1,
			mockCall: func() {
				mockGuest.On("GetByID", 1).Return(nil, errors.New("error")).Once()
			},
			response: nil,
			err:      errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewGuestServiceImpl(&repository.Repositories{GuestRepo: mockGuest})
			resp, err := service.GetByID(tc.id)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.response, resp)
		})
	}
}
