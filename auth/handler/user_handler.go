package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/ryansyahrullah/amk-be/auth/dto"
	"github.com/ryansyahrullah/amk-be/auth/repository"
	"github.com/ryansyahrullah/amk-be/auth/service"
	"github.com/ryansyahrullah/amk-be/pkg/utils"
)

// UserHandler menangani endpoint CRUD user.
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler membuat handler baru.
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Create membuat user baru.
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req dto.UserCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid body")
	}

	resp, err := h.service.CreateUser(c.Context(), req)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.Created(c, resp)
}

// List menampilkan user dengan filter sederhana.
func (h *UserHandler) List(c *fiber.Ctx) error {
	filter := repository.UserFilter{
		Query:    c.Query("q"),
		RoleSlug: c.Query("role"),
	}

	if active := c.Query("active"); active != "" {
		val := active == "true"
		filter.OnlyActive = &val
	}

	resp, err := h.service.ListUsers(c.Context(), filter)
	if err != nil {
		return err
	}

	return utils.Success(c, resp)
}

// Get mengambil user by id.
func (h *UserHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid id")
	}
	resp, err := h.service.GetUser(c.Context(), uint(id))
	if err != nil {
		return err
	}
	return utils.Success(c, resp)
}

// Update memperbarui user tertentu.
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid id")
	}
	var req dto.UserUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid body")
	}
	resp, err := h.service.UpdateUser(c.Context(), uint(id), req)
	if err != nil {
		return err
	}
	return utils.Success(c, resp)
}
