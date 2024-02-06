package v1

import (
	"context"
	"gateway/internal/controller/http/dto"
	"github.com/gofiber/fiber/v2"
	desc "github.com/nordew/GoChitChat-External/gen/go/user"
	"net/http"
)

func (h *Handler) signUp(c *fiber.Ctx) error {
	var input dto.SignUpRequest

	if err := c.BodyParser(&input); err != nil {
		return writeInvalidJSONResponse(c)
	}

	resp, err := h.userClient.Create(context.Background(), &desc.CreateUserRequest{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return writeErrorResponse(c, "failed to sign up", err.Error())
	}

	serverResp := fiber.Map{
		"user_id":       resp.Id,
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
	}

	return c.Status(http.StatusCreated).JSON(serverResp)
}

func (h *Handler) signIn(c *fiber.Ctx) error {
	var input dto.SignInRequest

	if err := c.BodyParser(&input); err != nil {
		return writeInvalidJSONResponse(c)
	}

	resp, err := h.userClient.Login(context.Background(), &desc.LoginUserRequest{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return writeErrorResponse(c, "failed to sign in", err.Error())
	}

	serverResp := fiber.Map{
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
	}

	return c.Status(http.StatusOK).JSON(serverResp)
}

func (h *Handler) refresh(c *fiber.Ctx) error {
	token := c.Get("refresh_token")

	resp, err := h.userClient.Refresh(context.Background(), &desc.RefreshUserRequest{
		RefreshToken: token,
	})
	if err != nil {
		return writeErrorResponse(c, "error refreshing token", err.Error())
	}

	serverResp := fiber.Map{
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
	}

	return c.Status(http.StatusOK).JSON(serverResp)
}
