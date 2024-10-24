package booking

import "github.com/zakiyalmaya/hotel-management/model"

type BookingService interface {
	Books(booking *model.BookingEntity) error
	GetByRegisterNumber(registerNumber string) (*model.BookingResponse, error)
	UpdatePayment(booking *model.BookingEntity) error
	Reschedule(booking *model.BookingEntity) error
}