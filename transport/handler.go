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
}
