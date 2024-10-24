package booking

import "github.com/zakiyalmaya/hotel-management/model"

type BookingRepository interface {
	Books(booking *model.BookingEntity) error
	GetByRegisterNumber(registerNumber string) (*model.BookingDetail, error)
	UpdatePayment(booking *model.BookingEntity) error
	Reschedule(booing *model.BookingEntity) error
}
