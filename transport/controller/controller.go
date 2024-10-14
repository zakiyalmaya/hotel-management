package controller

import (
	"github.com/zakiyalmaya/hotel-management/application"
	"github.com/zakiyalmaya/hotel-management/transport/controller/guest"
	"github.com/zakiyalmaya/hotel-management/transport/controller/room"
)

type Controller struct {
	RoomCtrl  *room.RoomController
	GuestCtrl *guest.GuestController
}

func NewController(application *application.Application) *Controller {
	return &Controller{
		RoomCtrl:  room.NewRoomController(application.RoomSvc),
		GuestCtrl: guest.NewGuestController(application.GuestSvc),
	}
}
