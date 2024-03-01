package http

import (
	"strconv"

	"github.com/anti-duhring/goncurrency/internal/transactions"
	"github.com/gofiber/fiber/v2"
)

type TransactionRequest struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

func transaction(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var request TransactionRequest
	if err := c.BodyParser(&request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = transactionsService.CreateTransaction(c.Context(), id, &transactions.Transaction{
		Amount:      request.Value,
		Operation:   request.Type,
		Description: request.Description,
	})

	return nil
}
