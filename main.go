package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/hotel-management/application"
	"github.com/zakiyalmaya/hotel-management/config"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
	"github.com/zakiyalmaya/hotel-management/transport"
)

func main() {
	// load config
	config.Init()
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Panic("failed to load configuration")
	}

	// instatiate repository
	db := repository.DBConnection(config.Database.File)
	redcl := repository.RedisClient(config.Redis.Host, config.Redis.Port)
	defer db.Close()

	repository := repository.NewRespository(db, redcl)

	// instantiate application
	application := application.NewApplication(repository)

	// instantiate fiber
	r := fiber.New()

	// instantiate transport
	transport.Handler(application, redcl, r)

	r.Listen(config.Server.Port)
}
