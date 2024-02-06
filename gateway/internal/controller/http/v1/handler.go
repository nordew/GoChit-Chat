package v1

import (
	"github.com/gofiber/fiber/v2"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
)

type Handler struct {
	userClient desc.UserClient
}

func NewHandler(userClient desc.UserClient) *Handler {
	return &Handler{
		userClient: userClient,
	}
}

func (h *Handler) Init() *fiber.App {
	app := fiber.New()

	auth := app.Group("/auth")
	{
		auth.Post("/sign-up", h.signUp)
		auth.Get("/sign-in", h.signIn)
		auth.Get("/refresh", h.refresh)
	}

	return app
}

func writeInvalidJSONResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Invalid JSON",
	})
}

func writeErrorResponse(c *fiber.Ctx, errType, errDescription string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":             errType,
		"error_description": errDescription,
	})
}
