package transport

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/hotel-management/application"
	"github.com/zakiyalmaya/hotel-management/middleware"
	"github.com/zakiyalmaya/hotel-management/transport/controller"
)

func Handler(application *application.Application, redcl *redis.Client, r *fiber.App) {
	ctrl := controller.NewController(application)

	r.Post("/api/room", middleware.AuthMiddleware(redcl), ctrl.RoomCtrl.Create)
	r.Get("/api/room", middleware.AuthMiddleware(redcl), ctrl.RoomCtrl.GetByName)
	r.Get("/api/rooms", middleware.AuthMiddleware(redcl), ctrl.RoomCtrl.GetAll)
	r.Put("/api/room/:name", middleware.AuthMiddleware(redcl), ctrl.RoomCtrl.Update)

	r.Post("/api/guest", middleware.AuthMiddleware(redcl), ctrl.GuestCtrl.Create)
	r.Get("/api/guest", middleware.AuthMiddleware(redcl), ctrl.GuestCtrl.GetByID)

	r.Post("/api/booking", middleware.AuthMiddleware(redcl), ctrl.BookingCtrl.Books)
	r.Get("/api/booking", middleware.AuthMiddleware(redcl), ctrl.BookingCtrl.GetByRegisterNumber)
	r.Put("/api/payment", middleware.AuthMiddleware(redcl), ctrl.BookingCtrl.UpdatePayment)
	r.Put("/api/reschedule", middleware.AuthMiddleware(redcl), ctrl.BookingCtrl.Reschedule)

	r.Post("/api/register", ctrl.UserCtrl.Create)
	r.Put("/api/password", middleware.AuthMiddleware(redcl), ctrl.UserCtrl.ChangePassword)

	r.Post("/auth/login", ctrl.AuthCtrl.Login)
	r.Post("/auth/logout", middleware.AuthMiddleware(redcl), ctrl.AuthCtrl.Logout)
	r.Post("/auth/refresh", middleware.AuthMiddleware(redcl), ctrl.AuthCtrl.Refresh)
}
