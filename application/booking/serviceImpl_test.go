package booking

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zakiyalmaya/hotel-management/constant"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository/mocks"
	"github.com/zakiyalmaya/hotel-management/model"
)

var (
	mockBooking *mocks.BookingRepository
	mockRoom    *mocks.RoomRepository
)

func Test_Books(t *testing.T) {
	mockBooking = new(mocks.BookingRepository)
	mockRoom = new(mocks.RoomRepository)

	testCases := []struct {
		name     string
		booking  *model.BookingEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when booking a room then return success",
			booking: &model.BookingEntity{
				RegisterNumber: "register_number",
				GuestID:        1,
				RoomName:       "room",
				CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
			},
			mockCall: func() {
				mockRoom.On("GetByName", "room").Return(&model.RoomEntity{Name: "room", Status: constant.RoomStatusAvailable, Price: 1000}, nil).Once()
				mockBooking.On("Books", &model.BookingEntity{
					RegisterNumber: "register_number",
					GuestID:        1,
					RoomName:       "room",
					CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
					CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
					PaidAmount:     1000,
				}).Return(nil).Once()
			},
			err: nil,
		},
		{
			name: "Given error get room data when booking a room then return error",
			booking: &model.BookingEntity{
				RegisterNumber: "register_number",
				GuestID:        1,
				RoomName:       "room",
				CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
			},
			mockCall: func() {
				mockRoom.On("GetByName", "room").Return(nil, errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
		{
			name: "Given error when booking a room then return error",
			booking: &model.BookingEntity{
				RegisterNumber: "register_number",
				GuestID:        1,
				RoomName:       "room",
				CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
			},
			mockCall: func() {
				mockRoom.On("GetByName", "room").Return(&model.RoomEntity{Name: "room", Status: constant.RoomStatusAvailable, Price: 1000}, nil).Once()
				mockBooking.On("Books", &model.BookingEntity{
					RegisterNumber: "register_number",
					GuestID:        1,
					RoomName:       "room",
					CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
					CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
					PaidAmount:     1000,
				}).Return(errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
		{
			name: "Given invalid room status when booking a room then return success",
			booking: &model.BookingEntity{
				RegisterNumber: "register_number",
				GuestID:        1,
				RoomName:       "room",
				CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
			},
			mockCall: func() {
				mockRoom.On("GetByName", "room").Return(&model.RoomEntity{Name: "room", Status: constant.RoomStatusMaintenance, Price: 1000}, nil).Once()
			},
			err: errors.New("room not available"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewBookingServiceImpl(&repository.Repositories{RoomRepo: mockRoom, BookingRepo: mockBooking})
			err := service.Books(tc.booking)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_GetByRegisterNumber(t *testing.T) {
	mockBooking = new(mocks.BookingRepository)
	time := time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC)

	testCases := []struct {
		name     string
		request  string
		mockCall func()
		err      error
		response *model.BookingResponse
	}{
		{
			name:    "Given valid request when get reservation data by register number then return success response",
			request: "register_number",
			mockCall: func() {
				mockBooking.On("GetByRegisterNumber", "register_number").Return(&model.BookingDetail{
					BookingEntity: model.BookingEntity{
						RegisterNumber: "register_number",
						CreatedAt:      &time,
						CheckIn:        time,
						CheckOut:       time,
						PaymentMethod:  constant.PaymentMethodCash,
						PaymentStatus:  constant.PaymentStatusCompleted,
					},
					RoomEntity: model.RoomEntity{
						Status: constant.RoomStatusBooked,
					},
				}, nil).Once()
			},
			err: nil,
			response: &model.BookingResponse{
				RegisterNumber: "register_number",
				GuestName:      " ",
				PaymentMethod:  constant.PaymentMethodCash.Enum(),
				PaymentStatus:  constant.PaymentStatusCompleted.Enum(),
				CheckIn:        "01 April 2000",
				CheckOut:       "01 April 2000",
				CreatedAt:      "01-04-2000 00:00:00",
				RoomStatus:     constant.RoomStatusBooked.Enum(),
			},
		},
		{
			name:    "Given err when get reservation data by register number then return error",
			request: "register_number",
			mockCall: func() {
				mockBooking.On("GetByRegisterNumber", "register_number").Return(nil, errors.New("error")).Once()
			},
			err:      errors.New("error"),
			response: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewBookingServiceImpl(&repository.Repositories{BookingRepo: mockBooking})
			resp, err := service.GetByRegisterNumber(tc.request)
			assert.Equal(t, tc.response, resp)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_UpdatePayment(t *testing.T) {
	mockBooking = new(mocks.BookingRepository)
	mockRoom = new(mocks.RoomRepository)

	testCases := []struct {
		name     string
		request  *model.BookingEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given request payment status completed when update payment then return success",
			request: &model.BookingEntity{
				RegisterNumber: "register_number",
				PaymentStatus:  constant.PaymentStatusCompleted,
			},
			mockCall: func() {
				mockBooking.On("UpdatePayment", &model.BookingEntity{
					RegisterNumber: "register_number",
					PaymentStatus:  constant.PaymentStatusCompleted,
				}).Return(nil).Once()

				mockBooking.On("GetByRegisterNumber", "register_number").Return(&model.BookingDetail{
					BookingEntity: model.BookingEntity{
						RegisterNumber: "register_number",
						PaymentStatus:  constant.PaymentStatusCompleted,
					},
					RoomEntity: model.RoomEntity{
						Floor: 1,
					},
				}, nil).Once()

				mockRoom.On("Update", &model.RoomEntity{Status: constant.RoomStatusBooked}).Return(nil).Once()
			},
			err: nil,
		},
		{
			name: "Given request payment status refunded when update payment then return success",
			request: &model.BookingEntity{
				RegisterNumber: "register_number",
				PaymentStatus:  constant.PaymentStatusRefunded,
			},
			mockCall: func() {
				mockBooking.On("UpdatePayment", &model.BookingEntity{
					RegisterNumber: "register_number",
					PaymentStatus:  constant.PaymentStatusRefunded,
				}).Return(nil).Once()

				mockBooking.On("GetByRegisterNumber", "register_number").Return(&model.BookingDetail{
					BookingEntity: model.BookingEntity{
						RegisterNumber: "register_number",
						PaymentStatus:  constant.PaymentStatusRefunded,
					},
					RoomEntity: model.RoomEntity{
						Floor: 1,
					},
				}, nil).Once()

				mockRoom.On("Update", &model.RoomEntity{Status: constant.RoomStatusAvailable}).Return(nil).Once()
			},
			err: nil,
		},
		{
			name: "Given error when update payment status to completed then return error",
			request: &model.BookingEntity{
				RegisterNumber: "register_number",
				PaymentStatus:  constant.PaymentStatusCompleted,
			},
			mockCall: func() {
				mockBooking.On("UpdatePayment", &model.BookingEntity{
					RegisterNumber: "register_number",
					PaymentStatus:  constant.PaymentStatusCompleted,
				}).Return(nil).Once()

				mockBooking.On("GetByRegisterNumber", "register_number").Return(&model.BookingDetail{
					BookingEntity: model.BookingEntity{
						RegisterNumber: "register_number",
						PaymentStatus:  constant.PaymentStatusCompleted,
					},
					RoomEntity: model.RoomEntity{
						Floor: 1,
					},
				}, nil).Once()

				mockRoom.On("Update", &model.RoomEntity{Status: constant.RoomStatusBooked}).Return(errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
		{
			name: "Given error when update payment status to refunded then return errors",
			request: &model.BookingEntity{
				RegisterNumber: "register_number",
				PaymentStatus:  constant.PaymentStatusRefunded,
			},
			mockCall: func() {
				mockBooking.On("UpdatePayment", &model.BookingEntity{
					RegisterNumber: "register_number",
					PaymentStatus:  constant.PaymentStatusRefunded,
				}).Return(nil).Once()

				mockBooking.On("GetByRegisterNumber", "register_number").Return(&model.BookingDetail{
					BookingEntity: model.BookingEntity{
						RegisterNumber: "register_number",
						PaymentStatus:  constant.PaymentStatusRefunded,
					},
					RoomEntity: model.RoomEntity{
						Floor: 1,
					},
				}, nil).Once()

				mockRoom.On("Update", &model.RoomEntity{Status: constant.RoomStatusAvailable}).Return(errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
		{
			name: "Given error when get reservation data by register number then return error",
			request: &model.BookingEntity{
				RegisterNumber: "register_number",
				PaymentStatus:  constant.PaymentStatusRefunded,
			},
			mockCall: func() {
				mockBooking.On("UpdatePayment", &model.BookingEntity{
					RegisterNumber: "register_number",
					PaymentStatus:  constant.PaymentStatusRefunded,
				}).Return(nil).Once()

				mockBooking.On("GetByRegisterNumber", "register_number").Return(nil, errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
		{
			name: "Given error when update payment then return error",
			request: &model.BookingEntity{
				RegisterNumber: "register_number",
				PaymentStatus:  constant.PaymentStatusRefunded,
			},
			mockCall: func() {
				mockBooking.On("UpdatePayment", &model.BookingEntity{
					RegisterNumber: "register_number",
					PaymentStatus:  constant.PaymentStatusRefunded,
				}).Return(errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewBookingServiceImpl(&repository.Repositories{BookingRepo: mockBooking, RoomRepo: mockRoom})
			err := service.UpdatePayment(tc.request)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_Reschedule(t *testing.T) {
	mockBooking = new(mocks.BookingRepository)
	mockRoom = new(mocks.RoomRepository)

	testCases := []struct {
		name     string
		request  *model.BookingEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when reschedule booking then return success",
			request: &model.BookingEntity{
				CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
				RegisterNumber: "register_number",
			},
			mockCall: func() {
				mockBooking.On("GetByRegisterNumber", "register_number").Return(&model.BookingDetail{
					BookingEntity: model.BookingEntity{
						RegisterNumber: "register_number",
						PaymentStatus:  constant.PaymentStatusPending,
						RoomName:       "room_name",
					},
					RoomEntity: model.RoomEntity{
						Name:   "room_name",
						Status: constant.RoomStatusAvailable,
					},
				}, nil).Once()

				mockRoom.On("GetByName", "room_name").Return(&model.RoomEntity{
					Name:  "room_name",
					Price: 1000,
				}, nil).Once()

				mockBooking.On("Reschedule", &model.BookingEntity{
					CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
					CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
					RegisterNumber: "register_number",
					PaidAmount:     1000,
				}).Return(nil).Once()
			},
			err: nil,
		},
		{
			name: "Given error when reschedule booking then return error",
			request: &model.BookingEntity{
				CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
				RegisterNumber: "register_number",
			},
			mockCall: func() {
				mockBooking.On("GetByRegisterNumber", "register_number").Return(&model.BookingDetail{
					BookingEntity: model.BookingEntity{
						RegisterNumber: "register_number",
						PaymentStatus:  constant.PaymentStatusPending,
						RoomName:       "room_name",
					},
					RoomEntity: model.RoomEntity{
						Name:   "room_name",
						Status: constant.RoomStatusAvailable,
					},
				}, nil).Once()

				mockRoom.On("GetByName", "room_name").Return(&model.RoomEntity{
					Name:  "room_name",
					Price: 1000,
				}, nil).Once()

				mockBooking.On("Reschedule", &model.BookingEntity{
					CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
					CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
					RegisterNumber: "register_number",
					PaidAmount:     1000,
				}).Return(errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
		{
			name: "Given error get room data when reschedule booking then return success",
			request: &model.BookingEntity{
				CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
				RegisterNumber: "register_number",
			},
			mockCall: func() {
				mockBooking.On("GetByRegisterNumber", "register_number").Return(&model.BookingDetail{
					BookingEntity: model.BookingEntity{
						RegisterNumber: "register_number",
						PaymentStatus:  constant.PaymentStatusPending,
						RoomName:       "room_name",
					},
					RoomEntity: model.RoomEntity{
						Name:   "room_name",
						Status: constant.RoomStatusAvailable,
					},
				}, nil).Once()

				mockRoom.On("GetByName", "room_name").Return(nil, errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
		{
			name: "Given invalid room status when reschedule booking then return success",
			request: &model.BookingEntity{
				CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
				RegisterNumber: "register_number",
			},
			mockCall: func() {
				mockBooking.On("GetByRegisterNumber", "register_number").Return(&model.BookingDetail{
					BookingEntity: model.BookingEntity{
						RegisterNumber: "register_number",
						PaymentStatus:  constant.PaymentStatusPending,
						RoomName:       "room_name",
					},
					RoomEntity: model.RoomEntity{
						Name:   "room_name",
						Status: constant.RoomStatusBooked,
					},
				}, nil).Once()
			},
			err: errors.New("invalid room status"),
		},
		{
			name: "Given invalid payment status when reschedule booking then return success",
			request: &model.BookingEntity{
				CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
				RegisterNumber: "register_number",
			},
			mockCall: func() {
				mockBooking.On("GetByRegisterNumber", "register_number").Return(&model.BookingDetail{
					BookingEntity: model.BookingEntity{
						RegisterNumber: "register_number",
						PaymentStatus:  constant.PaymentStatusCanceled,
						RoomName:       "room_name",
					},
					RoomEntity: model.RoomEntity{
						Name:   "room_name",
						Status: constant.RoomStatusAvailable,
					},
				}, nil).Once()
			},
			err: errors.New("invalid payment status"),
		},
		{
			name: "Given error get reservation data when reschedule booking then return success",
			request: &model.BookingEntity{
				CheckIn:        time.Date(2000, time.April, 01, 00, 00, 00, 00, time.UTC),
				CheckOut:       time.Date(2000, time.April, 02, 00, 00, 00, 00, time.UTC),
				RegisterNumber: "register_number",
			},
			mockCall: func() {
				mockBooking.On("GetByRegisterNumber", "register_number").Return(nil, errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewBookingServiceImpl(&repository.Repositories{RoomRepo: mockRoom, BookingRepo: mockBooking})
			err := service.Reschedule(tc.request)
			assert.Equal(t, tc.err, err)
		})
	}
}
