package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/ryansyahrullah/amk-be/auth/dto"
	"github.com/ryansyahrullah/amk-be/auth/service"
	"github.com/ryansyahrullah/amk-be/pkg/utils"
)

// RoleHandler menangani endpoint role & permission.
type RoleHandler struct {
	service *service.RoleService
}

// NewRoleHandler membuat handler baru.
func NewRoleHandler(service *service.RoleService) *RoleHandler {
	return &RoleHandler{service: service}
}

// List menampilkan semua role.
func (h *RoleHandler) List(c *fiber.Ctx) error {
	roles, err := h.service.ListRoles(c.Context())
	if err != nil {
		return err
	}
	return utils.Success(c, roles)
}

// Create membuat role baru.
func (h *RoleHandler) Create(c *fiber.Ctx) error {
	var body struct {
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
	}
	if err := c.BodyParser(&body); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid body")
	}

	role, err := h.service.CreateRole(c.Context(), body.Name, body.Slug, body.Description)
	if err != nil {
		return err
	}
	return utils.Created(c, role)
}

// Update memperbarui role.
func (h *RoleHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid id")
	}
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := c.BodyParser(&body); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid body")
	}

	role, err := h.service.UpdateRole(c.Context(), uint(id), body.Name, body.Description)
	if err != nil {
		return err
	}
	return utils.Success(c, role)
}

// UpdatePermissions memperbarui hak akses role.
func (h *RoleHandler) UpdatePermissions(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid id")
	}
	var body []dto.PermissionRequest
	if err := c.BodyParser(&body); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid body")
	}

	perms, err := h.service.UpsertPermissions(c.Context(), uint(id), body)
	if err != nil {
		return err
	}
	return utils.Success(c, perms)
}
