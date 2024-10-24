package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/hotel-management/application"
	"github.com/zakiyalmaya/hotel-management/transport/controller"
)

func Handler(application *application.Application, r *fiber.App) {
	ctrl := controller.NewController(application)

	r.Post("/room", ctrl.RoomCtrl.Create)
	r.Get("/room", ctrl.RoomCtrl.GetByName)
	r.Get("/rooms", ctrl.RoomCtrl.GetAll)
	r.Put("/room/:name", ctrl.RoomCtrl.Update)

	r.Post("/guest", ctrl.GuestCtrl.Create)
	r.Get("/guest", ctrl.GuestCtrl.GetByID)

	r.Post("/booking", ctrl.BookingCtrl.Books)
	r.Get("/booking", ctrl.BookingCtrl.GetByRegisterNumber)
	r.Put("/payment", ctrl.BookingCtrl.UpdatePayment)
	r.Put("/reschedule", ctrl.BookingCtrl.Reschedule)
}
