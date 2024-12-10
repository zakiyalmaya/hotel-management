package user

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/jmoiron/sqlx"
	"github.com/zakiyalmaya/hotel-management/model"
)

func Test_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	testCases := []struct {
		name     string
		user     *model.UserEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when create user then return success response",
			user: &model.UserEntity{
				Username: "username",
				Password: "password",
				Name:     "name",
				Email:    "email",
			},
			mockCall: func() {
				mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			err: nil,
		},
		{
			name: "Given error when create user then return error response",
			user: &model.UserEntity{
				Username: "username",
				Password: "password",
				Name:     "name",
				Email:    "email",
			},
			mockCall: func() {
				mock.ExpectExec("INSERT INTO users").WillReturnError(errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()
			repo := NewUserRepository(sqlxDB)

			err := repo.Create(tc.user)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_GetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	currentTime := time.Now()

	testCases := []struct {
		name     string
		username string
		result   *model.UserEntity
		mockCall func()
		err      error
	}{
		{
			name:     "Given valid request when get user by username then return success response",
			username: "username",
			result: &model.UserEntity{
				ID:        1,
				Username:  "username",
				Password:  "password",
				Name:      "name",
				Email:     "email",
				CreatedAt: &currentTime,
				UpdatedAt: &currentTime,
			},
			mockCall: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password", "name", "email", "created_at", "updated_at"}).
					AddRow(1, "username", "password", "name", "email", currentTime, currentTime)
				mock.ExpectQuery("SELECT id, name, username, password, email, created_at, updated_at FROM users WHERE username = ?").WithArgs("username").WillReturnRows(rows)
			},
			err: nil,
		},
		{
			name:     "Given error when get user by username then return error response",
			username: "username",
			mockCall: func() {
				mock.ExpectQuery("SELECT id, name, username, password, email, created_at, updated_at FROM users WHERE username = ?").WithArgs("username").WillReturnError(errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()
			repo := NewUserRepository(sqlxDB)

			result, err := repo.GetByUsername(tc.username)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.result, result)
		})
	}
}

func Test_UpdatePassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	testCases := []struct {
		name     string
		request  *model.UserEntity
		mockCall func()
		err      error
	}{
		{
			name:   "Given valid request when update password user then return success response",
			request: &model.UserEntity{
				Username: "username",
				Password: "password",
			},
			mockCall: func() {
				mock.ExpectExec("UPDATE users SET").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			err: nil,
		},
		{
			name:   "Given error when update password user then return error response",
			request: &model.UserEntity{
				Username: "username",
				Password: "password",
			},
			mockCall: func() {
				mock.ExpectExec("UPDATE users SET").WillReturnError(errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()
			repo := NewUserRepository(sqlxDB)

			err := repo.UpdatePassword(tc.request)
			assert.Equal(t, tc.err, err)
		})
	}
}
