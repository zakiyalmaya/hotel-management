package model

import (
	"time"

	"github.com/zakiyalmaya/hotel-management/constant"
)

type RoomEntity struct {
	ID          int                 `db:"id"`
	Name        string              `db:"name"`
	Floor       int                 `db:"floor"`
	Type        string              `db:"type"`
	Price       float64             `db:"price"`
	Status      constant.RoomStatus `db:"status"`
	Description *string             `db:"description"`
	CreatedAt   *time.Time          `db:"created_at"`
	UpdatedAt   *time.Time          `db:"updated_at"`
}

type RoomResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Floor       int     `json:"floor"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
	Description *string `json:"description"`
}

type RoomsResponse struct {
	Rooms []*RoomResponse `json:"rooms"`
}

type CreateRoomRequest struct {
	Name        string  `json:"name" validate:"required"`
	Floor       int     `json:"floor" validate:"required,gt=0"`
	Type        string  `json:"type" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Status      int     `json:"status" validate:"required"`
	Description *string `json:"description,omitempty"`
}

type GetAllRoomRequest struct {
	Floor  *string `json:"floor,omitempty"`
	Status *string `json:"status,omitempty"`
}

type UpdateRoomRequest struct {
	Name        string   `json:"name" validate:"required"`
	Floor       *int     `json:"floor,omitempty"`
	Status      *int     `json:"status,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Description *string  `json:"description,omitempty"`
}
