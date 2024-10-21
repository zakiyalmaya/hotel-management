package application

import (
	"github.com/zakiyalmaya/hotel-management/application/guest"
	"github.com/zakiyalmaya/hotel-management/application/room"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
)

type Application struct {
	Repos    *repository.Repositories
	RoomSvc  room.RoomService
	GuestSvc guest.GuestService
}

func NewApplication(repos *repository.Repositories) *Application {
	return &Application{
		Repos:    repos,
		RoomSvc:  room.NewRoomServiceImpl(repos),
		GuestSvc: guest.NewGuestServiceImpl(repos),
	}
}
