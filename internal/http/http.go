package http

import (
	"database/sql"

	"github.com/anti-duhring/goncurrency/internal/clients"
	"github.com/anti-duhring/goncurrency/internal/transactions"
	"github.com/gofiber/fiber/v2"
)

var (
	clientsService      *clients.Service
	transactionsService *transactions.Service
)

func Init(DB *sql.DB) *fiber.App {
	app := fiber.New()

	clientsRepository := clients.NewRepositoryPostgres(DB)
	clientsService = clients.NewService(clientsRepository)

	transactionsRepository := transactions.NewRepositoryPostgres(DB)
	transactionsService = transactions.NewService(transactionsRepository, clientsRepository)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "api is running",
		})
	})

	clients := app.Group("/clientes")
	clients.Get("/:id/extrato", getExtract)
	clients.Post("/:id/transacoes", transaction)

	return app
}
