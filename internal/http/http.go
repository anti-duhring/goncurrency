package http

import (
	"database/sql"

	"github.com/anti-duhring/goncurrency/internal/clients"
	"github.com/gofiber/fiber/v2"
)

var clientsService *clients.Service

func Init(DB *sql.DB) *fiber.App {
	app := fiber.New()

	clientsRepository := clients.NewRepositoryPostgres(DB)
	clientsService = clients.NewService(clientsRepository)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "api is running",
		})
	})

	clients := app.Group("/clientes")
	clients.Get("/:id/extrato", getExtract)

	return app
}
