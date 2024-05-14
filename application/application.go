package application

import (
	"github.com/zakiyalmaya/hotel-management/application/room"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
)

type Application struct {
	Repos   *repository.Repositories
	RoomSvc room.Service
}

func NewApplication(repos *repository.Repositories) *Application {
	return &Application{
		Repos:   repos,
		RoomSvc: room.NewRoomServiceImpl(repos),
	}
}
