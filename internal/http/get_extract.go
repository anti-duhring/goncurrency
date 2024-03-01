package http

import (
	"context"
	"errors"
	"strconv"

	"github.com/anti-duhring/goncurrency/internal/clients"
	"github.com/anti-duhring/goncurrency/internal/transactions"
	"github.com/anti-duhring/goncurrency/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type GetExtractResponse struct {
	Saldo             clients.ClientExtract      `json:"saldo"`
	UltimasTransacoes []transactions.Transaction `json:"ultimas_transacoes"`
}

func getExtract(c *fiber.Ctx) error {
	idStr := c.Params("id")
	ctx := context.Background()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error("error converting id to int", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	clientExtract, err := clientsService.GetClientExtract(ctx, id)
	if err != nil {
		logger.Error("error calling service.GetClientExtract", err)

		if errors.Is(err, clients.ErrClientNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	transactions, err := transactionsService.GetTransactionsFromClient(ctx, id, 10)
	if err != nil {
		logger.Error("error calling service.GetTransactionsFromClient", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"saldo":              clientExtract,
		"ultimas_transacoes": *transactions,
	})
}
