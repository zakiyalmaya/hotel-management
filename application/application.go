package application

import (
	"github.com/zakiyalmaya/hotel-management/application/booking"
	"github.com/zakiyalmaya/hotel-management/application/guest"
	"github.com/zakiyalmaya/hotel-management/application/room"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
)

type Application struct {
	Repos      *repository.Repositories
	RoomSvc    room.RoomService
	GuestSvc   guest.GuestService
	BookingSvc booking.BookingService
}

func NewApplication(repos *repository.Repositories) *Application {
	return &Application{
		Repos:      repos,
		RoomSvc:    room.NewRoomServiceImpl(repos),
		GuestSvc:   guest.NewGuestServiceImpl(repos),
		BookingSvc: booking.NewBookingServiceImpl(repos),
	}
}
