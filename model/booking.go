package model

import (
	"time"

	"github.com/zakiyalmaya/hotel-management/constant"
)

type BookingEntity struct {
	ID                int                    `db:"id"`
	RegisterNumber    string                 `db:"register_number"`
	GuestID           int                    `db:"guest_id"`
	RoomName          string                 `db:"room_name"`
	CheckIn           time.Time              `db:"check_in"`
	CheckOut          time.Time              `db:"check_out"`
	PaidAmount        float64                `db:"paid_amount"`
	PaymentMethod     constant.PaymentMethod `db:"payment_method"`
	PaymentStatus     constant.PaymentStatus `db:"payment_status"`
	AdditionalRequest *string                `db:"additional_request"`
	CreatedAt         *time.Time             `db:"created_at"`
	UpdatedAt         *time.Time             `db:"updated_at"`
}

type BookingDetail struct {
	BookingEntity
	GuestEntity
	RoomEntity
}

type BookingRequest struct {
	GuestID         int    `json:"guest_id" validate:"required"`
	RoomName        string `json:"room_name" validate:"required"`
	CheckIn         string `json:"check_in" validate:"required"`
	CheckOut        string `json:"check_out" validate:"required"`
	PaymentMethod   int    `json:"payment_method" validate:"required"`
	AdditionRequest string `json:"additional_request,omitempty"`
}

type BookingResponse struct {
	RegisterNumber    string  `json:"register_number"`
	GuestID           int     `json:"guest_id"`
	GuestName         string  `json:"guest_name"`
	GuestIdentity     string  `json:"guest_identity"`
	RoomName          string  `json:"room_name"`
	RoomFloor         int     `json:"room_floor"`
	RoomType          string  `json:"room_type"`
	RoomStatus        string  `json:"room_status"`
	CheckIn           string  `json:"check_in"`
	CheckOut          string  `json:"check_out"`
	PaidAmount        float64 `json:"paid_amount"`
	PaymentMethod     string  `json:"payment_method"`
	PaymentStatus     string  `json:"payment_status"`
	AdditionalRequest *string `json:"additional_request"`
	CreatedAt         string  `json:"created_at"`
}

type UpdatePaymentRequest struct {
	RegisterNumber string `json:"register_number" validate:"required"`
	PaymentStatus  int    `json:"payment_status" validate:"required"`
}

type ResceduleRequest struct {
	RegisterNumber string `json:"register_number" validate:"required"`
	CheckIn        string `json:"check_in" validate:"required"`
	CheckOut       string `json:"check_out" validate:"required"`
}
