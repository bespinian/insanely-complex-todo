package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Message string `json:"error"`
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if err := ctx.Status(code).JSON(ErrorResponse{Message: e.Message}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return nil
}
