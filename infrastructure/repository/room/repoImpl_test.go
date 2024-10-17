package room

import (
	"database/sql"
	"errors"
	"testing"

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
		room     *model.RoomEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when create room then return success response",
			room: &model.RoomEntity{
				Name:   "name",
				Floor:  1,
				Type:   "type",
				Price:  100,
				Status: 1,
			},
			mockCall: func() {
				mock.ExpectExec("INSERT INTO rooms").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			err: nil,
		},
		{
			name: "Given error connection when create room then return error",
			room: &model.RoomEntity{
				Name:   "name",
				Floor:  1,
				Type:   "type",
				Price:  100,
				Status: 1,
			},
			mockCall: func() {
				mock.ExpectExec("INSERT INTO rooms").WillReturnError(sql.ErrConnDone)
			},
			err: sql.ErrConnDone,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			repo := NewRoomRepository(sqlxDB)

			err := repo.Create(tc.room)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_GetByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	testCases := []struct {
		name     string
		request  string
		result   *model.RoomEntity
		mockCall func()
		err      error
	}{
		{
			name:    "Given valid request when get room by name then return success response",
			request: "name",
			result: &model.RoomEntity{
				ID:          1,
				Name:        "name",
				Floor:       1,
				Type:        "type",
				Price:       100,
				Status:      1,
				Description: nil,
				CreatedAt:   nil,
				UpdatedAt:   nil,
			},
			mockCall: func() {
				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "floor", "type", "price", "status", "description", "created_at", "updated_at"}).
					AddRow(1, "name", 1, "type", 100, 1, nil, nil, nil))
			},
			err: nil,
		},
		{
			name:    "Given error connection when get room by name then return error",
			request: "name",
			result:  nil,
			mockCall: func() {
				mock.ExpectQuery("SELECT").WillReturnError(sql.ErrConnDone)
			},
			err: sql.ErrConnDone,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			repo := NewRoomRepository(sqlxDB)

			result, err := repo.GetByName(tc.request)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.result, result)
		})
	}
}

func Test_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	floor := "1"
	status := "1"

	testCases := []struct {
		name     string
		request  *model.GetAllRoomRequest
		result   []*model.RoomEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when get all room then return success response",
			request: &model.GetAllRoomRequest{
				Floor:  &floor,
				Status: &status,
			},
			result: []*model.RoomEntity{
				{
					ID:          1,
					Name:        "name",
					Floor:       1,
					Type:        "type",
					Price:       100,
					Status:      1,
					Description: nil,
					CreatedAt:   nil,
					UpdatedAt:   nil,
				},
			},
			mockCall: func() {
				mock.ExpectQuery("SELECT").WithArgs(floor, status).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "floor", "type", "price", "status", "description", "created_at", "updated_at"}).
					AddRow(1, "name", 1, "type", 100, 1, nil, nil, nil))
			},
			err: nil,
		},
		{
			name: "Given error connection when get all room then return error",
			request: &model.GetAllRoomRequest{
				Floor:  &floor,
				Status: &status,
			},
			result: nil,
			mockCall: func() {
				mock.ExpectQuery("SELECT").WithArgs(floor, status).WillReturnError(sql.ErrConnDone)
			},
			err: sql.ErrConnDone,
		},
		{
			name: "Given error when get all room then return error",
			request: &model.GetAllRoomRequest{
				Floor:  &floor,
				Status: &status,
			},
			result: nil,
			mockCall: func() {
				mock.ExpectQuery("SELECT").WithArgs(floor, status).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "floor", "type", "price"}).AddRow(1, "name", 1, "type", 100))
			},
			err: errors.New("sql: expected 5 destination arguments in Scan, not 9"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			repo := NewRoomRepository(sqlxDB)

			result, err := repo.GetAll(tc.request)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.result, result)
		})
	}
}

func Test_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	var description string = "description"

	testCases := []struct {
		name     string
		request  *model.RoomEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when update room then return success response",
			request: &model.RoomEntity{
				Name:        "name",
				Floor:       1,
				Type:        "type",
				Price:       100,
				Status:      1,
				Description: &description,
			},
			mockCall: func() {
				mock.ExpectExec("UPDATE rooms").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			err: nil,
		},
		{
			name: "Given error connection when update room then return error",
			request: &model.RoomEntity{
				Name:        "name",
				Floor:       1,
				Type:        "type",
				Price:       100,
				Status:      1,
				Description: &description,
			},
			mockCall: func() {
				mock.ExpectExec("UPDATE rooms").WillReturnError(sql.ErrConnDone)
			},
			err: sql.ErrConnDone,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			repo := NewRoomRepository(sqlxDB)

			err := repo.Update(tc.request)
			assert.Equal(t, tc.err, err)
		})
	}
}
