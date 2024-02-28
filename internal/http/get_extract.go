package http

import (
	"context"
	"strconv"

	"github.com/anti-duhring/goncurrency/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func getExtract(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error("error converting id to int", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	clientExtract, err := clientsService.GetClientExtract(context.Background(), id)
	if err != nil {
		logger.Error("error calling service.GetClientExtract", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": clientExtract,
	})
}
