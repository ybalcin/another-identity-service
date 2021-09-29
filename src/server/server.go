package server

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/ybalcin/another-identity-service/startup"
	"log"
)

func newServer() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")

	InitAccountHandlers(api)

	return app
}

func Serve() {
	server := newServer()

	log.Fatal(server.Listen(":" + AppConfig.Port))
}
