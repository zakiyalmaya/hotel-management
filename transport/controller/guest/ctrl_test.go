package guest

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/zakiyalmaya/hotel-management/application/mocks"
	"github.com/zakiyalmaya/hotel-management/model"
)

var (
	mockGuest *mocks.GuestService
)

func Test_Create(t *testing.T) {
	app := fiber.New()
	mockGuest = new(mocks.GuestService)

	testCases := []struct {
		name       string
		request    interface{}
		mockCall   func()
		statusCode int
	}{
		{
			name: "Given valid request when create new guest then return success response",
			request: &model.CreateGuestRequest{
				FirstName:   "John",
				LastName:    "Doe",
				Identity:    "identity",
				DateOfBirth: "01-12-2000",
				PhoneNumber: "08123456789",
				Email:       "example@mail.com",
			},
			mockCall: func() {
				mockGuest.On("Create", &model.GuestEntity{
					FirstName:   "John",
					LastName:    "Doe",
					Identity:    "identity",
					DateOfBirth: time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC),
					PhoneNumber: "08123456789",
					Email:       "example@mail.com",
				}).Return(nil).Once()
			},
			statusCode: http.StatusCreated,
		},
		{
			name: "Given valid request when create new guest then return success response",
			request: &model.CreateGuestRequest{
				FirstName:   "John",
				LastName:    "Doe",
				Identity:    "identity",
				DateOfBirth: "01-12-2000",
				PhoneNumber: "08123456789",
				Email:       "example@mail.com",
			},
			mockCall: func() {
				mockGuest.On("Create", &model.GuestEntity{
					FirstName:   "John",
					LastName:    "Doe",
					Identity:    "identity",
					DateOfBirth: time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC),
					PhoneNumber: "08123456789",
					Email:       "example@mail.com",
				}).Return(errors.New("error")).Once()
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "Given invalid date of birth when create new guest then return bad request response",
			request: &model.CreateGuestRequest{
				FirstName:   "John",
				LastName:    "Doe",
				Identity:    "identity",
				DateOfBirth: "01/12/2000",
				PhoneNumber: "08123456789",
				Email:       "example@mail.com",
			},
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Given invalid request when create new guest then return bad request response",
			request:    &model.CreateGuestRequest{},
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Given error parsing request when create new guest then return bad request response",
			request:    "test",
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := NewGuestController(mockGuest)
			app.Post("/guest", ctrl.Create)

			body, _ := json.Marshal(tc.request)
			req, _ := http.NewRequest("POST", "/guest", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			tc.mockCall()

			res, _ := app.Test(req)
			assert.Equal(t, tc.statusCode, res.StatusCode)
		})
	}
}

func Test_GetByID(t *testing.T) {
	app := fiber.New()
	mockGuest = new(mocks.GuestService)

	testCases := []struct {
		name       string
		request    string
		mockCall   func()
		statusCode int
	}{
		{
			name:    "Given valid id request when get guest by id then return success response",
			request: "1",
			mockCall: func() {
				mockGuest.On("GetByID", 1).Return(&model.GuestResponse{
					ID:   1,
					Name: "test",
				}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			name:    "Given valid id request when get guest by id then return success response",
			request: "1",
			mockCall: func() {
				mockGuest.On("GetByID", 1).Return(nil, errors.New("error")).Once()
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "Given invalid id request when get guest by id then return bad request",
			request:    "a",
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := NewGuestController(mockGuest)
			app.Get("/guest", ctrl.GetByID)

			req, _ := http.NewRequest("GET", "/guest?id="+tc.request, nil)
			req.Header.Set("Content-Type", "application/json")
			tc.mockCall()

			res, _ := app.Test(req)
			assert.Equal(t, tc.statusCode, res.StatusCode)
		})
	}
}
