package user

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository/mocks"
	"github.com/zakiyalmaya/hotel-management/model"
)

var (
	mockUser *mocks.UserRepository
)

func Test_Create(t *testing.T) {
	mockUser = new(mocks.UserRepository)

	testCases := []struct {
		name     string
		request  interface{}
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when create a new user then return success response",
			request: &model.CreateUserRequest{
				Name:     "name",
				Username: "username",
				Password: "password",
				Email:    "email",
			},
			mockCall: func() {
				mockUser.On("Create", mock.Anything).Return(nil).Once()
			},
			err: nil,
		},
		{
			name: "Given error when create a new user then return error response",
			request: &model.CreateUserRequest{
				Name:     "name",
				Username: "username",
				Password: "password",
				Email:    "email",
			},
			mockCall: func() {
				mockUser.On("Create", mock.Anything).Return(errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewUserServiceImpl(&repository.Repositories{UserRepo: mockUser})
			err := service.Create(tc.request.(*model.CreateUserRequest))
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_ChangePassword(t *testing.T) {
	mockUser = new(mocks.UserRepository)

	testCases := []struct {
		name     string
		request  interface{}
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when change password then return success response",
			request: &model.ChangePasswordRequest{
				Username:    "username",
				OldPassword: "password",
				NewPassword: "newPassword",
			},
			mockCall: func() {
				mockUser.On("GetByUsername", mock.Anything).Return(&model.UserEntity{
					Username: "username",
					Password: "$2a$10$ajTCV9x1V1.SDLkVIQMGX.2zASIoBojgofRV0s3Mnwg5cIfMNIsyy",
				}, nil).Once()

				mockUser.On("UpdatePassword", mock.Anything).Return(nil).Once()
			},
			err: nil,
		},
		{
			name: "Given new password same as old password when change password then return error response",
			request: &model.ChangePasswordRequest{
				Username:    "username",
				OldPassword: "password",
				NewPassword: "password",
			},
			mockCall: func() {
				mockUser.On("GetByUsername", mock.Anything).Return(&model.UserEntity{
					Username: "username",
					Password: "$2a$10$ajTCV9x1V1.SDLkVIQMGX.2zASIoBojgofRV0s3Mnwg5cIfMNIsyy",
				}, nil).Once()
			},
			err: errors.New("new password must be different from old password"),
		},
		{
			name: "Given wrong old password when change password then return error response",
			request: &model.ChangePasswordRequest{
				Username:    "username",
				OldPassword: "wrong password",
				NewPassword: "newPassword",
			},
			mockCall: func() {
				mockUser.On("GetByUsername", mock.Anything).Return(&model.UserEntity{
					Username: "username",
					Password: "$2a$10$ajTCV9x1V1.SDLkVIQMGX.2zASIoBojgofRV0s3Mnwg5cIfMNIsky",
				}, nil).Once()
			},
			err: errors.New("wrong password"),
		},
		{
			name: "Given error when change password then return error response",
			request: &model.ChangePasswordRequest{
				Username:    "username",
				OldPassword: "password",
				NewPassword: "newPassword",
			},
			mockCall: func() {
				mockUser.On("GetByUsername", mock.Anything).Return(&model.UserEntity{
					Username: "username",
					Password: "$2a$10$ajTCV9x1V1.SDLkVIQMGX.2zASIoBojgofRV0s3Mnwg5cIfMNIsyy",
				}, nil).Once()

				mockUser.On("UpdatePassword", mock.Anything).Return(errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
		{
			name: "Given error when get user by username then return error response",
			request: &model.ChangePasswordRequest{
				Username:    "username",
				OldPassword: "password",
				NewPassword: "newPassword",
			},
			mockCall: func() {
				mockUser.On("GetByUsername", mock.Anything).Return(nil, errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
		{
			name: "Given error no rows when get user by username then return error response",
			request: &model.ChangePasswordRequest{
				Username:    "username",
				OldPassword: "password",
				NewPassword: "newPassword",
			},
			mockCall: func() {
				mockUser.On("GetByUsername", mock.Anything).Return(nil, sql.ErrNoRows).Once()
			},
			err: errors.New("user not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewUserServiceImpl(&repository.Repositories{UserRepo: mockUser})
			err := service.ChangePassword(tc.request.(*model.ChangePasswordRequest))
			assert.Equal(t, tc.err, err)
		})
	}
}
