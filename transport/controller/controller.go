package controller

import (
	"github.com/zakiyalmaya/hotel-management/application"
	"github.com/zakiyalmaya/hotel-management/transport/controller/booking"
	"github.com/zakiyalmaya/hotel-management/transport/controller/guest"
	"github.com/zakiyalmaya/hotel-management/transport/controller/room"
	"github.com/zakiyalmaya/hotel-management/transport/controller/user"
)

type Controller struct {
	RoomCtrl    *room.RoomController
	GuestCtrl   *guest.GuestController
	BookingCtrl *booking.BookingController
	UserCtrl    *user.UserController
}

func NewController(application *application.Application) *Controller {
	return &Controller{
		RoomCtrl:    room.NewRoomController(application.RoomSvc),
		GuestCtrl:   guest.NewGuestController(application.GuestSvc),
		BookingCtrl: booking.NewBookingController(application.BookingSvc),
		UserCtrl:    user.NewUserController(application.UserSvc),
	}
}
