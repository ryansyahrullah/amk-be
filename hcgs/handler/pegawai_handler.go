package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ryansyahrullah/amk-be/hcgs/dto"
	"github.com/ryansyahrullah/amk-be/hcgs/service"
	"github.com/ryansyahrullah/amk-be/pkg/utils"
)

// PegawaiHandler menangani endpoint profil pegawai.
type PegawaiHandler struct {
	service *service.PegawaiService
}

// NewPegawaiHandler membuat handler baru.
func NewPegawaiHandler(service *service.PegawaiService) *PegawaiHandler {
	return &PegawaiHandler{service: service}
}

// GetMe mengembalikan profil pegawai berdasarkan token.
func (h *PegawaiHandler) GetMe(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return utils.Error(c, fiber.StatusUnauthorized, "invalid token")
	}

	resp, err := h.service.GetProfile(c.Context(), userID)
	if err != nil {
		return err
	}
	return utils.Success(c, resp)
}

// UpdateMe memperbarui profil pegawai.
func (h *PegawaiHandler) UpdateMe(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return utils.Error(c, fiber.StatusUnauthorized, "invalid token")
	}

	var req dto.PegawaiUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid body")
	}

	resp, err := h.service.UpdateProfile(c.Context(), userID, req)
	if err != nil {
		return err
	}
	return utils.Success(c, resp)
}
