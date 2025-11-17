package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ryansyahrullah/amk-be/auth/dto"
	"github.com/ryansyahrullah/amk-be/auth/service"
	"github.com/ryansyahrullah/amk-be/pkg/utils"
)

// PasswordHandler menangani endpoint lupa/reset password.
type PasswordHandler struct {
	service *service.PasswordService
}

// NewPasswordHandler membuat handler baru.
func NewPasswordHandler(service *service.PasswordService) *PasswordHandler {
	return &PasswordHandler{service: service}
}

// ForgotPassword memproses POST /auth/forgot-password.
func (h *PasswordHandler) ForgotPassword(c *fiber.Ctx) error {
	var req dto.PasswordForgotRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid body")
	}

	if err := h.service.ForgotPassword(c.Context(), req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.Success(c, fiber.Map{"message": "OTP terkirim"})
}

// ResetPassword memproses POST /auth/reset-password.
func (h *PasswordHandler) ResetPassword(c *fiber.Ctx) error {
	var req dto.PasswordResetRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid body")
	}

	if err := h.service.ResetPassword(c.Context(), req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.Success(c, fiber.Map{"message": "Password updated"})
}
