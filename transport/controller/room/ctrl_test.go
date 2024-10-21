package room

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/hotel-management/application/mocks"
	"github.com/zakiyalmaya/hotel-management/constant"
	"github.com/zakiyalmaya/hotel-management/model"
)

func Test_Create(t *testing.T) {
	app := fiber.New()
	mockRoom := new(mocks.RoomService)

	testCases := []struct {
		name       string
		request    interface{}
		mockCall   func()
		statusCode int
	}{
		{
			name: "Given valid request when create a new room then return success response",
			request: &model.CreateRoomRequest{
				Name:   "room name",
				Floor:  1,
				Type:   "type",
				Price:  100.00,
				Status: 1,
			},
			mockCall: func() {
				mockRoom.On("Create", &model.RoomEntity{
					Name:   "room name",
					Floor:  1,
					Type:   "type",
					Price:  100.00,
					Status: constant.RoomStatus(1),
				}).Return(nil).Once()
			},
			statusCode: http.StatusCreated,
		},
		{
			name: "Given error when create a new room then return error response",
			request: &model.CreateRoomRequest{
				Name:   "room name",
				Floor:  1,
				Type:   "type",
				Price:  100.00,
				Status: 1,
			},
			mockCall: func() {
				mockRoom.On("Create", &model.RoomEntity{
					Name:   "room name",
					Floor:  1,
					Type:   "type",
					Price:  100.00,
					Status: constant.RoomStatus(1),
				}).Return(errors.New("error")).Once()
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "Given missing request when create a new room then return bad request response",
			request: &model.CreateRoomRequest{
				Name: "room name",
			},
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Given invalid request when create a new room then return bad request",
			request:    "test",
			mockCall:   func() {},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := NewRoomController(mockRoom)
			app.Post("/room", ctrl.Create)

			body, _ := json.Marshal(tc.request)
			req, _ := http.NewRequest("POST", "/room", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			tc.mockCall()

			res, _ := app.Test(req)
			assert.Equal(t, tc.statusCode, res.StatusCode)
		})
	}
}

func Test_GetByName(t *testing.T) {
	app := fiber.New()
	mockRoom := new(mocks.RoomService)

	testCases := []struct {
		name       string
		request    string
		mockCall   func()
		statusCode int
	}{
		{
			name:    "Given valid request when get room by name then return success",
			request: "roomname",
			mockCall: func() {
				mockRoom.On("GetByName", "roomname").Return(&model.RoomResponse{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			name:    "Given error when get room by name then return error response",
			request: "roomname",
			mockCall: func() {
				mockRoom.On("GetByName", "roomname").Return(nil, errors.New("error")).Once()
			},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := NewRoomController(mockRoom)
			app.Get("/room", ctrl.GetByName)

			req, _ := http.NewRequest("GET", "/room?name="+tc.request, nil)
			req.Header.Set("Content-Type", "application/json")
			tc.mockCall()

			res, _ := app.Test(req)
			assert.Equal(t, tc.statusCode, res.StatusCode)
		})
	}
}

func Test_GetAll(t *testing.T) {
	app := fiber.New()
	mockRoom := new(mocks.RoomService)

	var (
		floor  = "1"
		status = "1"
	)

	testCases := []struct {
		name       string
		request    map[string]string
		mockCall   func()
		statusCode int
	}{
		{
			name: "Given valid request when get all rooms then return success response",
			request: map[string]string{
				"floor":  floor,
				"status": status,
			},
			mockCall: func() {
				mockRoom.On("GetAll", &model.GetAllRoomRequest{
					Floor:  &floor,
					Status: &status,
				}).Return([]*model.RoomResponse{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			name: "Given error when get all rooms then return error response",
			request: map[string]string{
				"floor":  floor,
				"status": status,
			},
			mockCall: func() {
				mockRoom.On("GetAll", &model.GetAllRoomRequest{
					Floor:  &floor,
					Status: &status,
				}).Return(nil, errors.New("error")).Once()
			},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := NewRoomController(mockRoom)
			app.Get("/rooms", ctrl.GetAll)

			query := url.Values{}
			for k, v := range tc.request {
				query.Set(k, v)
			}
			req, _ := http.NewRequest("GET", "/rooms?"+query.Encode(), nil)
			req.Header.Set("Content-Type", "application/json")
			tc.mockCall()

			res, _ := app.Test(req)
			assert.Equal(t, tc.statusCode, res.StatusCode)
		})
	}
}

func Test_Update(t *testing.T) {
	app := fiber.New()
	mockRoom := new(mocks.RoomService)

	var (
		floor  = 1
		status = 1
		price  = 100.00
		desc   = "description"
	)

	testCases := []struct {
		name       string
		param      string
		request    interface{}
		mockCall   func()
		statusCode int
	}{
		{
			name:  "Given valid request when update room data then return success",
			param: "roomname",
			request: &model.UpdateRoomRequest{
				Name:        "roomname",
				Floor:       &floor,
				Status:      &status,
				Price:       &price,
				Description: &desc,
			},
			mockCall: func() {
				mockRoom.On("Update", &model.UpdateRoomRequest{
					Name:        "roomname",
					Floor:       &floor,
					Status:      &status,
					Price:       &price,
					Description: &desc,
				}).Return(nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			name:  "Given error when update room data then return error response",
			param: "roomname",
			request: &model.UpdateRoomRequest{
				Name:        "roomname",
				Floor:       &floor,
				Status:      &status,
				Price:       &price,
				Description: &desc,
			},
			mockCall: func() {
				mockRoom.On("Update", &model.UpdateRoomRequest{
					Name:        "roomname",
					Floor:       &floor,
					Status:      &status,
					Price:       &price,
					Description: &desc,
				}).Return(errors.New("error")).Once()
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:  "Given invalid request when update room data then return error",
			param: "roomname",
			request: "test",
			mockCall: func() {},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := NewRoomController(mockRoom)
			app.Put("/room/:name", ctrl.Update)

			body, _ := json.Marshal(tc.request)
			req, _ := http.NewRequest("PUT", "/room/"+tc.param, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			tc.mockCall()

			res, _ := app.Test(req)
			assert.Equal(t, tc.statusCode, res.StatusCode)
		})
	}
}
