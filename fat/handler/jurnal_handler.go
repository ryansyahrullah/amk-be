package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/ryansyahrullah/amk-be/fat/dto"
	"github.com/ryansyahrullah/amk-be/fat/service"
	"github.com/ryansyahrullah/amk-be/pkg/utils"
)

// JurnalHandler menangani endpoint jurnal umum.
type JurnalHandler struct {
	service *service.JurnalService
}

// NewJurnalHandler membuat handler baru.
func NewJurnalHandler(service *service.JurnalService) *JurnalHandler {
	return &JurnalHandler{service: service}
}

// List menampilkan daftar jurnal.
func (h *JurnalHandler) List(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	resp, err := h.service.List(c.Context(), limit)
	if err != nil {
		return err
	}
	return utils.Success(c, resp)
}

// Create membuat jurnal baru.
func (h *JurnalHandler) Create(c *fiber.Ctx) error {
	var req dto.JurnalRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid body")
	}

	userID, _ := c.Locals("user_id").(uint)
	resp, err := h.service.Create(c.Context(), req, userID)
	if err != nil {
		return err
	}
	return utils.Created(c, resp)
}
