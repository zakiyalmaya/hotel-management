package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/hotel-management/application"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
	"github.com/zakiyalmaya/hotel-management/transport"
)

func main() {
	// instatiate repository
	db := repository.DBConnection()
	defer db.Close()

	repository := repository.NewRespository(db)

	// instantiate application
	application := application.NewApplication(repository)

	// instantiate fiber
	r := fiber.New()

	// instantiate transport
	transport.Handler(application, r)

	r.Listen(":3000")
}
