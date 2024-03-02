package http

import (
	"errors"
	"strconv"

	"github.com/anti-duhring/goncurrency/internal/transactions"
	"github.com/anti-duhring/goncurrency/pkg/logger"
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

	clientExtract, err := transactionsService.CreateTransaction(c.Context(), id, &transactions.Transaction{
		Amount:      request.Value,
		Operation:   request.Type,
		Description: request.Description,
	})
	if err != nil {
		if errors.Is(err, transactions.ErrAccountLimitExceeded) {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}
		logger.Error("error calling service.CreateTransaction", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"limite": clientExtract.AccountLimit,
		"saldo":  clientExtract.Balance,
	})
}
