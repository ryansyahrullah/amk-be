package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ryansyahrullah/amk-be/auth/dto"
	"github.com/ryansyahrullah/amk-be/auth/service"
	"github.com/ryansyahrullah/amk-be/pkg/utils"
)

// AuthHandler menampung endpoint terkait autentikasi.
type AuthHandler struct {
	service *service.AuthService
}

// NewAuthHandler membuat handler baru.
func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Login memproses POST /auth/login.
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.AuthLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid body")
	}

	resp, err := h.service.Login(c.Context(), req)
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.Success(c, resp)
}
