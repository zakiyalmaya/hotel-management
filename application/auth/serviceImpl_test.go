package auth

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-playground/assert/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/mock"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository/mocks"
	"github.com/zakiyalmaya/hotel-management/model"
)

var (
	mockUser        *mocks.UserRepository
	mockRedisClient *redis.Client
)

func Test_Login(t *testing.T) {
	mockUser = new(mocks.UserRepository)
	mockRedisServer, err := miniredis.Run()
	if err != nil {
		t.Fatalf(err.Error())
	}

	mockRedisClient = redis.NewClient(&redis.Options{
		Addr: mockRedisServer.Addr(),
	})

	duration := 15 * time.Minute

	testCases := []struct {
		name     string
		request  interface{}
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when login then return success response",
			request: &model.AuthRequest{
				Username: "username",
				Password: "password",
			},
			mockCall: func() {
				mockUser.On("GetByUsername", "username").Return(&model.UserEntity{
					Username: "username",
					Password: "$2a$10$ajTCV9x1V1.SDLkVIQMGX.2zASIoBojgofRV0s3Mnwg5cIfMNIsyy",
				}, nil).Once()

				mockRedisClient.Set(context.Background(), "jwt-token-username", mock.Anything, duration).Err()
			},
			err: nil,
		},
		{
			name: "Given wrong password when login then return error response",
			request: &model.AuthRequest{
				Username: "username",
				Password: "password",
			},
			mockCall: func() {
				mockUser.On("GetByUsername", "username").Return(&model.UserEntity{
					Username: "username",
					Password: "wrong-password",
				}, nil).Once()
			},
			err: errors.New("wrong password"),
		},
		{
			name: "Given error not found when get user by username then return error response",
			request: &model.AuthRequest{
				Username: "username",
				Password: "password",
			},
			mockCall: func() {
				mockUser.On("GetByUsername", "username").Return(nil, sql.ErrNoRows).Once()
			},
			err: errors.New("user not found"),
		},
		{
			name: "Given error when get user by username then return error response",
			request: &model.AuthRequest{
				Username: "username",
				Password: "password",
			},
			mockCall: func() {
				mockUser.On("GetByUsername", "username").Return(nil, errors.New("error")).Once()
			},
			err: errors.New("error getting hotelier by username"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewAuthServiceImpl(&repository.Repositories{UserRepo: mockUser, RedCl: mockRedisClient})
			_, err := service.Login(tc.request.(*model.AuthRequest))
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_Logout(t *testing.T) {
	mockRedisServer, err := miniredis.Run()
	if err != nil {
		t.Fatalf(err.Error())
	}

	mockRedisClient = redis.NewClient(&redis.Options{
		Addr: mockRedisServer.Addr(),
	})

	testCases := []struct {
		name     string
		request  string
		mockCall func()
		err      error
	}{
		{
			name:    "Given valid request when logout then return success response",
			request: "username",
			mockCall: func() {
				mockRedisClient.Del(context.Background(), "jwt-token-username").Err()
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewAuthServiceImpl(&repository.Repositories{RedCl: mockRedisClient})
			err := service.Logout(tc.request)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_RefreshAuthToken(t *testing.T) {
	mockRedisServer, err := miniredis.Run()
	if err != nil {
		t.Fatalf(err.Error())
	}

	mockRedisClient = redis.NewClient(&redis.Options{
		Addr: mockRedisServer.Addr(),
	})

	testCases := []struct {
		name     string
		request  string
		mockCall func()
		err      bool
	}{
		{
			name:    "Given error get token from redis when refresh auth token then return error response",
			request: "username",
			mockCall: func() {
				mockRedisClient.Get(context.Background(), "jwt-token-username").Err()
			},
			err: true,
		},
		{
			name:    "Given invalid token when parse token then return error response",
			request: "username",
			mockCall: func() {
				mockRedisClient.Get(context.Background(), "jwt-token-username").Result()
				mockRedisClient.Set(context.Background(), "jwt-token-username", mock.Anything, 15*time.Minute).Err()
			},
			err: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewAuthServiceImpl(&repository.Repositories{RedCl: mockRedisClient})
			_, err := service.RefreshAuthToken(tc.request)
			assert.Equal(t, tc.err, err != nil)
		})
	}
}
