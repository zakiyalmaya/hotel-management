package booking

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/zakiyalmaya/hotel-management/constant"
	"github.com/zakiyalmaya/hotel-management/model"
)

func Test_Books(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	testCases := []struct {
		name     string
		booking  *model.BookingEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when reservation room then return success response",
			booking: &model.BookingEntity{
				RegisterNumber: "register_number",
				GuestID:        1,
				RoomName:       "101",
				CheckIn:        time.Date(2000, time.December, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.December, 02, 00, 00, 00, 00, time.UTC),
				PaidAmount:     1000000,
				PaymentMethod:  constant.PaymentMethodCash,
				PaymentStatus:  constant.PaymentStatusPending,
			},
			mockCall: func() {
				mock.ExpectExec("INSERT INTO bookings").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			err: nil,
		},
		{
			name: "Given error canceled when reservation room then return error response",
			booking: &model.BookingEntity{
				RegisterNumber: "register_number",
				GuestID:        1,
				RoomName:       "101",
				CheckIn:        time.Date(2000, time.December, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.December, 02, 00, 00, 00, 00, time.UTC),
				PaidAmount:     1000000,
				PaymentMethod:  constant.PaymentMethodCash,
				PaymentStatus:  constant.PaymentStatusPending,
			},
			mockCall: func() {
				mock.ExpectExec("INSERT INTO bookings").WillReturnError(sqlmock.ErrCancelled)
			},
			err: sqlmock.ErrCancelled,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			repo := NewBookingRepository(sqlxDB)

			err := repo.Books(tc.booking)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_GetByRegisterNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	testCases := []struct {
		name     string
		request  string
		mockCall func()
		response *model.BookingDetail
		err      error
	}{
		{
			name:    "Given valid request when get reservation data then return success",
			request: "register_number",
			mockCall: func() {
				mock.ExpectQuery("SELECT").WithArgs("register_number").WillReturnRows(sqlmock.NewRows([]string{"register_number", "guest_id", "first_name", "last_name",
					"identity_number", "room_name", "floor", "type", "status", "check_in", "check_out", "paid_amount", "payment_method", "payment_status", "additional_request", "created_at"}).
					AddRow("register_number", "1", "John", "Doe", "identity", "room", "1", "", "1", time.Time{}, time.Time{}, "1000", "2", "1", nil, nil))
			},
			response: &model.BookingDetail{
				BookingEntity: model.BookingEntity{
					RegisterNumber: "register_number",
					GuestID:        1,
					RoomName:       "room",
					PaymentStatus:  constant.PaymentStatusPending,
					PaidAmount:     1000,
					PaymentMethod:  constant.PaymentMethodBankTransfer,
				},
				GuestEntity: model.GuestEntity{
					FirstName: "John",
					LastName:  "Doe",
					Identity:  "identity",
				},
				RoomEntity: model.RoomEntity{
					Floor:  1,
					Status: constant.RoomStatusAvailable,
				},
			},
			err: nil,
		},
		{
			name:    "Given error when get reservation data then return error",
			request: "register_number",
			mockCall: func() {
				mock.ExpectQuery("SELECT").WithArgs("register_number").WillReturnError(errors.New(("error")))
			},
			response: nil,
			err:      errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			repo := NewBookingRepository(sqlxDB)

			resp, err := repo.GetByRegisterNumber(tc.request)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.response, resp)
		})
	}
}

func Test_UpdatePayment(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	testCases := []struct {
		name     string
		booking  *model.BookingEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when update payment then return success response",
			booking: &model.BookingEntity{
				RegisterNumber: "register_number",
				PaymentStatus:  constant.PaymentStatusCanceled,
			},
			mockCall: func() {
				mock.ExpectExec("UPDATE bookings").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			err: nil,
		},
		{
			name: "Given error when update payment then return error response",
			booking: &model.BookingEntity{
				RegisterNumber: "register_number",
				PaymentStatus:  constant.PaymentStatusCanceled,
			},
			mockCall: func() {
				mock.ExpectExec("UPDATE bookings").WillReturnError(sqlmock.ErrCancelled)
			},
			err: sqlmock.ErrCancelled,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			repo := NewBookingRepository(sqlxDB)

			err := repo.UpdatePayment(tc.booking)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_Reschedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	testCases := []struct {
		name     string
		booking  *model.BookingEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when reschedule reservation then return success response",
			booking: &model.BookingEntity{
				RegisterNumber: "register_number",
				CheckIn:        time.Date(2000, time.March, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.March, 03, 00, 00, 00, 00, time.UTC),
				PaidAmount:     1000000,
			},
			mockCall: func() {
				mock.ExpectExec("UPDATE bookings").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			err: nil,
		},
		{
			name: "Given error when reschedule reservation then return error response",
			booking: &model.BookingEntity{
				RegisterNumber: "register_number",
				CheckIn:        time.Date(2000, time.March, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.March, 03, 00, 00, 00, 00, time.UTC),
				PaidAmount:     1000000,
			},
			mockCall: func() {
				mock.ExpectExec("UPDATE bookings").WillReturnError(sqlmock.ErrCancelled)
			},
			err: sqlmock.ErrCancelled,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			repo := NewBookingRepository(sqlxDB)

			err := repo.Reschedule(tc.booking)
			assert.Equal(t, tc.err, err)
		})
	}
}
