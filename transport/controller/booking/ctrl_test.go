package booking

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/zakiyalmaya/hotel-management/application/mocks"
	"github.com/zakiyalmaya/hotel-management/constant"
	"github.com/zakiyalmaya/hotel-management/model"
)

var (
	mockBooing *mocks.BookingService
)

func Test_Books(t *testing.T) {
	app := fiber.New()
	mockBooing = new(mocks.BookingService)

	testCases := []struct {
		name       string
		request    interface{}
		mockCall   func()
		statusCode int
	}{
		{
			name: "Given valid request when booking room then return success",
			request: &model.BookingRequest{
				GuestID:         1,
				RoomName:        "room_name",
				CheckIn:         "01-12-2000",
				CheckOut:        "02-12-2000",
				PaymentMethod:   1,
				AdditionRequest: "additional_request",
			},
			mockCall: func() {
				mockBooing.On("Books", mock.Anything).Return(nil).Once()
			},
			statusCode: http.StatusCreated,
		},
		{
			name: "Given error when booking room then return error",
			request: &model.BookingRequest{
				GuestID:         1,
				RoomName:        "room_name",
				CheckIn:         "01-12-2000",
				CheckOut:        "02-12-2000",
				PaymentMethod:   1,
				AdditionRequest: "additional_request",
			},
			mockCall: func() {
				mockBooing.On("Books", mock.Anything).Return(errors.New("error")).Once()
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "Given invalid check out date format when booking room then return error",
			request: &model.BookingRequest{
				GuestID:         1,
				RoomName:        "room_name",
				CheckIn:         "01-12-2000",
				CheckOut:        "02/12/2000",
				PaymentMethod:   1,
				AdditionRequest: "additional_request",
			},
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Given invalid check in date format when booking room then return error",
			request: &model.BookingRequest{
				GuestID:         1,
				RoomName:        "room_name",
				CheckIn:         "01/12/2000",
				CheckOut:        "02-12-2000",
				PaymentMethod:   1,
				AdditionRequest: "additional_request",
			},
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Given check in date > check out date when booking room then return error",
			request: &model.BookingRequest{
				GuestID:         1,
				RoomName:        "room_name",
				CheckIn:         "11-12-2000",
				CheckOut:        "02-12-2000",
				PaymentMethod:   1,
				AdditionRequest: "additional_request",
			},
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Given invalid payment method when booking room then return error",
			request: &model.BookingRequest{
				GuestID:         1,
				RoomName:        "room_name",
				CheckIn:         "01-12-2000",
				CheckOut:        "02-12-2000",
				PaymentMethod:   10,
				AdditionRequest: "additional_request",
			},
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Given missing required request when booking room then return error",
			request: &model.BookingRequest{
				GuestID:         1,
				RoomName:        "room_name",
				CheckIn:         "11-12-2000",
				PaymentMethod:   1,
				AdditionRequest: "additional_request",
			},
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Given empty body request when booking room then return error",
			request:    "a",
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := NewBookingController(mockBooing)
			app.Post("/booking", ctrl.Books)

			body, _ := json.Marshal(tc.request)
			req, _ := http.NewRequest("POST", "/booking", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			tc.mockCall()

			res, _ := app.Test(req)
			assert.Equal(t, tc.statusCode, res.StatusCode)
		})
	}
}

func Test_GetByRegisterNumber(t *testing.T) {
	app := fiber.New()
	mockBooing = new(mocks.BookingService)

	testCases := []struct {
		name       string
		request    string
		mockCall   func()
		statusCode int
	}{
		{
			name:    "Given valid request when get reservation data by register number then return success",
			request: "register_number",
			mockCall: func() {
				mockBooing.On("GetByRegisterNumber", "register_number").Return(&model.BookingResponse{
					RegisterNumber: "register_number",
				}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			name:    "Given error when get reservation data by register number then return error",
			request: "register_number",
			mockCall: func() {
				mockBooing.On("GetByRegisterNumber", "register_number").Return(nil, errors.New("error")).Once()
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "Given invalid request when get reservation data by register number then return error",
			request:    "",
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := NewBookingController(mockBooing)
			app.Get("/booking", ctrl.GetByRegisterNumber)

			req, _ := http.NewRequest("GET", "/booking?register_number="+tc.request, nil)
			req.Header.Set("Content-Type", "application/json")
			tc.mockCall()

			res, _ := app.Test(req)
			assert.Equal(t, tc.statusCode, res.StatusCode)
		})
	}
}

func Test_UpdatePayment(t *testing.T) {
	app := fiber.New()
	mockBooing = new(mocks.BookingService)

	testCases := []struct {
		name       string
		request    interface{}
		mockCall   func()
		statusCode int
	}{
		{
			name: "Given valid request when update payment then return success",
			request: &model.UpdatePaymentRequest{
				RegisterNumber: "register_number",
				PaymentStatus:  2,
			},
			mockCall: func() {
				mockBooing.On("UpdatePayment", &model.BookingEntity{
					RegisterNumber: "register_number",
					PaymentStatus:  constant.PaymentStatusCompleted,
				}).Return(nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			name: "Given valid request when update payment then return success",
			request: &model.UpdatePaymentRequest{
				RegisterNumber: "register_number",
				PaymentStatus:  2,
			},
			mockCall: func() {
				mockBooing.On("UpdatePayment", &model.BookingEntity{
					RegisterNumber: "register_number",
					PaymentStatus:  constant.PaymentStatusCompleted,
				}).Return(errors.New("error")).Once()
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "Given missing required request when update payment then return error",
			request: &model.UpdatePaymentRequest{
				RegisterNumber: "register_number",
			},
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Given invalid payment status when update payment then return error",
			request: &model.UpdatePaymentRequest{
				RegisterNumber: "register_number",
				PaymentStatus:  10,
			},
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Given invalid request when update payment then return error",
			request:    "a",
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := NewBookingController(mockBooing)
			app.Put("/payment", ctrl.UpdatePayment)

			body, _ := json.Marshal(tc.request)
			req, _ := http.NewRequest("PUT", "/payment", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			tc.mockCall()

			res, _ := app.Test(req)
			assert.Equal(t, tc.statusCode, res.StatusCode)
		})
	}
}

func Test_Rescedule(t *testing.T) {
	app := fiber.New()
	mockBooing = new(mocks.BookingService)

	testCases := []struct {
		name       string
		request    interface{}
		mockCall   func()
		statusCode int
	}{
		{
			name: "Given valid request when reschedule reservation then return success",
			request: &model.ResceduleRequest{
				RegisterNumber: "register_number",
				CheckIn:        "01-12-2024",
				CheckOut:       "02-12-2024",
			},
			mockCall: func() {
				mockBooing.On("Reschedule", &model.BookingEntity{
					RegisterNumber: "register_number",
					CheckIn:        time.Date(2024, time.December, 01, 00, 00, 00, 00, time.UTC),
					CheckOut:       time.Date(2024, time.December, 02, 00, 00, 00, 00, time.UTC),
				}).Return(nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			name: "Given error when reschedule reservation then return error",
			request: &model.ResceduleRequest{
				RegisterNumber: "register_number",
				CheckIn:        "01-12-2024",
				CheckOut:       "02-12-2024",
			},
			mockCall: func() {
				mockBooing.On("Reschedule", &model.BookingEntity{
					RegisterNumber: "register_number",
					CheckIn:        time.Date(2024, time.December, 01, 00, 00, 00, 00, time.UTC),
					CheckOut:       time.Date(2024, time.December, 02, 00, 00, 00, 00, time.UTC),
				}).Return(errors.New("error")).Once()
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "Given invalid check out date when reschedule reservation then return error",
			request: &model.ResceduleRequest{
				RegisterNumber: "register_number",
				CheckIn:        "01-12-2024",
				CheckOut:       "02/12/2024",
			},
			mockCall: func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Given invalid check in date when reschedule reservation then return error",
			request: &model.ResceduleRequest{
				RegisterNumber: "register_number",
				CheckIn:        "01/12/2024",
				CheckOut:       "02-12-2024",
			},
			mockCall: func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Given missing required request when reschedule reservation then return error",
			request: &model.ResceduleRequest{
				CheckIn:        "01-12-2024",
				CheckOut:       "02-12-2024",
			},
			mockCall: func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Given invalid request when reschedule reservation then return error",
			request: "a",
			mockCall: func() {},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := NewBookingController(mockBooing)
			app.Put("/reschedule", ctrl.Reschedule)

			body, _ := json.Marshal(tc.request)
			req, _ := http.NewRequest("PUT", "/reschedule", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			tc.mockCall()

			res, _ := app.Test(req)
			assert.Equal(t, tc.statusCode, res.StatusCode)
		})
	}
}
