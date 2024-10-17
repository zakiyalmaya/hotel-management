package guest

import (
	"database/sql"
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
				DateOfBirth: time.Now(),
				PhoneNumber: "phone",
				Email:       "email",
			},
			mockCall: func() { mock.ExpectExec("INSERT INTO guests").WillReturnResult(sqlmock.NewResult(1, 1)) },
			err:      nil,
		},
		{
			name: "Given error connection when create guest then return error",
			guest: &model.GuestEntity{
				FirstName:   "first",
				LastName:    "last",
				Identity:    "identity",
				DateOfBirth: time.Now(),
				PhoneNumber: "phone",
				Email:       "email",
			},
			mockCall: func() { mock.ExpectExec("INSERT INTO guests").WillReturnError(sql.ErrConnDone) },
			err:      sql.ErrConnDone,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			repo := NewGuestRepository(sqlxDB)

			err := repo.Create(tc.guest)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	currentTime := time.Now()

	testCases := []struct {
		name     string
		id       int
		result   *model.GuestEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when get guest by id then return success response",
			id:   1,
			result: &model.GuestEntity{
				ID:          1,
				FirstName:   "first",
				LastName:    "last",
				Identity:    "identity",
				DateOfBirth: currentTime,
				PhoneNumber: "phone",
				Email:       "email",
				CreatedAt:   &currentTime,
				UpdatedAt:   &currentTime,
			},
			mockCall: func() {
				mock.ExpectQuery("SELECT id, first_name, last_name, identity_number, date_of_birth, phone_number, email, created_at, updated_at FROM guests WHERE id = ?").WillReturnRows(
					sqlmock.NewRows([]string{"id", "first_name", "last_name", "identity_number", "date_of_birth", "phone_number", "email", "created_at", "updated_at"}).
						AddRow(1, "first", "last", "identity", currentTime, "phone", "email", currentTime, currentTime),
				)
			},
			err: nil,
		},
		{
			name:   "Given error when get guest by id to db then return error response",
			id:     1,
			result: nil,
			mockCall: func() {
				mock.ExpectQuery("SELECT id, first_name, last_name, identity_number, date_of_birth, phone_number, email, created_at, updated_at FROM guests WHERE id = ?").WillReturnError(sql.ErrNoRows)
			},
			err: sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			repo := NewGuestRepository(sqlxDB)

			result, err := repo.GetByID(tc.id)
			assert.Equal(t, tc.result, result)
			assert.Equal(t, tc.err, err)
		})
	}
}
