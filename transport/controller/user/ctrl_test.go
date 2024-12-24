package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/zakiyalmaya/hotel-management/application/mocks"
	"github.com/zakiyalmaya/hotel-management/model"
)

func Test_Create(t *testing.T) {
	app := fiber.New()
	mockUser := new(mocks.UserService)

	testCases := []struct {
		name       string
		request    interface{}
		mockCall   func()
		statusCode int
	}{
		{
			name: "Given valid request when create a new user then return success response",
			request: &model.CreateUserRequest{
				Name:     "John Doe",
				Username: "username",
				Email:    "example@mail.com",
				Password: "password",
			},
			mockCall: func() {
				mockUser.On("Create", &model.CreateUserRequest{
					Name:     "John Doe",
					Username: "username",
					Email:    "example@mail.com",
					Password: "password",
				}).Return(nil).Once()
			},
			statusCode: fiber.StatusCreated,
		},
		{
			name: "Given error when create a new user then return error response",
			request: &model.CreateUserRequest{
				Name:     "John Doe",
				Username: "username",
				Email:    "example@mail.com",
				Password: "password",
			},
			mockCall: func() {
				mockUser.On("Create", &model.CreateUserRequest{
					Name:     "John Doe",
					Username: "username",
					Email:    "example@mail.com",
					Password: "password",
				}).Return(errors.New("error")).Once()
			},
			statusCode: fiber.StatusInternalServerError,
		},
		{
			name:       "Given missing request when create a new user then return error response",
			request:    &model.CreateUserRequest{},
			mockCall:   func() {},
			statusCode: fiber.StatusBadRequest,
		},
		{
			name:       "Given invalid request when create a new user then return error response",
			request:    "invalid request",
			mockCall:   func() {},
			statusCode: fiber.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := NewUserController(mockUser)
			app.Post("/users", ctrl.Create)

			body, _ := json.Marshal(tc.request)
			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			tc.mockCall()

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tc.statusCode, resp.StatusCode)
		})
	}
}

func Test_ChangePassword(t *testing.T) {
	app := fiber.New()
	mockUser := new(mocks.UserService)

	testCases := []struct {
		name       string
		username   interface{}
		request    interface{}
		mockCall   func()
		statusCode int
	}{
		{
			name: "Given valid request when change password then return success response",
			username: "testuser",
			request: &model.ChangePasswordRequest{
				OldPassword: "oldpassword",
				NewPassword: "newpassword",
			},
			mockCall: func() {
				mockUser.On("ChangePassword", &model.ChangePasswordRequest{
					Username:    "testuser",
					OldPassword: "oldpassword",
					NewPassword: "newpassword",
				}).Return(nil).Once()
			},
			statusCode: fiber.StatusOK,
		},
		{
			name: "Given error when change password then return error response",
			username: "testuser",
			request: &model.ChangePasswordRequest{
				OldPassword: "oldpassword",
				NewPassword: "newpassword",
			},
			mockCall: func() {
				mockUser.On("ChangePassword", &model.ChangePasswordRequest{
					Username:    "testuser",
					OldPassword: "oldpassword",
					NewPassword: "newpassword",
				}).Return(errors.New("error")).Once()
			},
			statusCode: fiber.StatusInternalServerError,
		},
		{
			name:       "Given invalid request when change password then return error response",
			username:   "testuser",
			request:    "invalid request",
			mockCall:   func() {},
			statusCode: fiber.StatusBadRequest,
		},
		{
			name:       "Given error invalid username when change password then return error response",
			request:    &model.ChangePasswordRequest{},
			mockCall:   func() {},
			statusCode: fiber.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := NewUserController(mockUser)
			app.Post("/api/password", func(ctx *fiber.Ctx) error {
				ctx.Locals("username", tc.username)
				return ctrl.ChangePassword(ctx)
			})

			body, _ := json.Marshal(tc.request)
			req, _ := http.NewRequest("POST", "/api/password", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			tc.mockCall()

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tc.statusCode, resp.StatusCode)
		})
	}
}
