package room

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zakiyalmaya/hotel-management/constant"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository/mocks"
	"github.com/zakiyalmaya/hotel-management/model"
)

var (
	mockRoom *mocks.RoomRepository
)

func Test_Create(t *testing.T) {
	mockRoom = new(mocks.RoomRepository)

	testCases := []struct {
		name     string
		room     *model.RoomEntity
		mockCall func()
		err      error
	}{
		{
			name: "Given valid request when insert new room then return success",
			room: &model.RoomEntity{
				Name:   "room name",
				Floor:  1,
				Type:   "type",
				Price:  100,
				Status: constant.RoomStatusAvailable,
			},
			mockCall: func() {
				mockRoom.On("Create", &model.RoomEntity{
					Name:   "room name",
					Floor:  1,
					Type:   "type",
					Price:  100,
					Status: constant.RoomStatusAvailable,
				}).Return(nil).Once()
			},
			err: nil,
		},
		{
			name: "Given error when insert new room then return error",
			room: &model.RoomEntity{
				Name:   "room name",
				Floor:  1,
				Type:   "type",
				Price:  100,
				Status: constant.RoomStatusAvailable,
			},
			mockCall: func() {
				mockRoom.On("Create", &model.RoomEntity{
					Name:   "room name",
					Floor:  1,
					Type:   "type",
					Price:  100,
					Status: constant.RoomStatusAvailable,
				}).Return(errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewRoomServiceImpl(&repository.Repositories{RoomRepo: mockRoom})
			err := service.Create(tc.room)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_GetByName(t *testing.T) {
	mockRoom = new(mocks.RoomRepository)

	testCases := []struct {
		name     string
		request  string
		mockCall func()
		err      error
		response *model.RoomResponse
	}{
		{
			name:    "Given valid request when get room by name then return success response",
			request: "room name",
			mockCall: func() {
				mockRoom.On("GetByName", "room name").Return(&model.RoomEntity{
					ID:     1,
					Name:   "room name",
					Floor:  1,
					Type:   "type",
					Price:  100,
					Status: constant.RoomStatusAvailable,
				}, nil).Once()
			},
			err: nil,
			response: &model.RoomResponse{
				ID:     1,
				Name:   "room name",
				Floor:  1,
				Type:   "type",
				Price:  100,
				Status: constant.RoomStatusAvailable.Enum(),
			},
		},
		{
			name:    "Given error when get room by name then return error",
			request: "room name",
			mockCall: func() {
				mockRoom.On("GetByName", "room name").Return(nil, errors.New("error")).Once()
			},
			err:      errors.New("error"),
			response: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewRoomServiceImpl(&repository.Repositories{RoomRepo: mockRoom})
			resp, err := service.GetByName(tc.request)
			assert.Equal(t, tc.response, resp)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_GetAll(t *testing.T) {
	mockRoom = new(mocks.RoomRepository)

	var (
		floor  = "1"
		status = "1"
	)

	testCases := []struct {
		name     string
		request  *model.GetAllRoomRequest
		mockCall func()
		err      error
		response []*model.RoomResponse
	}{
		{
			name: "Given valid request when get all rooms then return success",
			request: &model.GetAllRoomRequest{
				Floor:  &floor,
				Status: &status,
			},
			mockCall: func() {
				mockRoom.On("GetAll", &model.GetAllRoomRequest{
					Floor:  &floor,
					Status: &status,
				}).Return([]*model.RoomEntity{
					{ID: 1, Price: 0, Status: constant.RoomStatusAvailable},
					{ID: 2, Price: 0, Status: constant.RoomStatusAvailable},
				}, nil).Once()
			},
			err: nil,
			response: []*model.RoomResponse{
				{ID: 1, Price: 0, Status: constant.RoomStatusAvailable.Enum()},
				{ID: 2, Price: 0, Status: constant.RoomStatusAvailable.Enum()},
			},
		},
		{
			name: "Given valid request when get all rooms then return success",
			request: &model.GetAllRoomRequest{
				Floor:  &floor,
				Status: &status,
			},
			mockCall: func() {
				mockRoom.On("GetAll", &model.GetAllRoomRequest{
					Floor:  &floor,
					Status: &status,
				}).Return(nil, errors.New("error")).Once()
			},
			err:      errors.New("error"),
			response: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewRoomServiceImpl(&repository.Repositories{RoomRepo: mockRoom})
			resp, err := service.GetAll(tc.request)
			assert.Equal(t, tc.response, resp)
			assert.Equal(t, tc.err, err)
		})
	}
}

func Test_Update(t *testing.T) {
	mockRoom = new(mocks.RoomRepository)

	var (
		floor       = 1
		status      = 1
		price       = 100.00
		description = "description"
	)

	testCases := []struct {
		name     string
		request  *model.UpdateRoomRequest
		mockCall func()
		err      error
	}{
		{
			name: "Given invalid request when update room then return success response",
			request: &model.UpdateRoomRequest{
				Name:        "room name",
				Floor:       &floor,
				Status:      &status,
				Price:       &price,
				Description: &description,
			},
			mockCall: func() {
				mockRoom.On("Update", &model.RoomEntity{
					Name:        "room name",
					Floor:       floor,
					Status:      constant.RoomStatus(status),
					Price:       price,
					Description: &description,
				}).Return(nil).Once()
			},
			err: nil,
		},
		{
			name: "Given error when update room then return error",
			request: &model.UpdateRoomRequest{
				Name:        "room name",
				Floor:       &floor,
				Status:      &status,
				Price:       &price,
				Description: &description,
			},
			mockCall: func() {
				mockRoom.On("Update", &model.RoomEntity{
					Name:        "room name",
					Floor:       floor,
					Status:      constant.RoomStatus(status),
					Price:       price,
					Description: &description,
				}).Return(errors.New("error")).Once()
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockCall()

			service := NewRoomServiceImpl(&repository.Repositories{RoomRepo: mockRoom})
			err := service.Update(tc.request)
			assert.Equal(t, tc.err, err)
		})
	}
}
