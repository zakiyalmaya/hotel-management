package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/hotel-management/application"
	"github.com/zakiyalmaya/hotel-management/transport/controller"
)

func Handler(application *application.Application, r *fiber.App) {
	ctrl := controller.NewController(application)

	r.Post("/api/room", ctrl.RoomCtrl.Create)
	r.Get("/api/room", ctrl.RoomCtrl.GetByName)
	r.Get("/api/rooms", ctrl.RoomCtrl.GetAll)
	r.Put("/api/room/:name", ctrl.RoomCtrl.Update)

	r.Post("/api/guest", ctrl.GuestCtrl.Create)
	r.Get("/api/guest", ctrl.GuestCtrl.GetByID)

	r.Post("/api/booking", ctrl.BookingCtrl.Books)
	r.Get("/api/booking", ctrl.BookingCtrl.GetByRegisterNumber)
	r.Put("/api/payment", ctrl.BookingCtrl.UpdatePayment)
	r.Put("/api/reschedule", ctrl.BookingCtrl.Reschedule)
}
