package server

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/ybalcin/another-identity-service/handlers"
)

func InitAccountHandlers(router fiber.Router) {
	handlers := NewAccountHandler()

	accountApi := router.Group("/v1/accounts")

	accountApi.Get("/", handlers.GetAllHandler)
}
