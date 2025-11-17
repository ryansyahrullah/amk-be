package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	authModel "github.com/ryansyahrullah/amk-be/auth/model"
)

// PermissionMiddleware memeriksa hak akses CRUD berdasarkan role.
type PermissionMiddleware struct {
	db *gorm.DB
}

// NewPermissionMiddleware membuat middleware baru dengan koneksi DB.
func NewPermissionMiddleware(db *gorm.DB) *PermissionMiddleware {
	return &PermissionMiddleware{db: db}
}

// Require memastikan role memiliki izin terhadap table/action tertentu.
func (m *PermissionMiddleware) Require(tableName, action string) fiber.Handler {
	normalizedAction := strings.ToLower(action)

	return func(c *fiber.Ctx) error {
		roleSlug, _ := c.Locals("role_slug").(string)
		if roleSlug == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  "error",
				"message": "missing role",
			})
		}

		var role authModel.Role
		if err := m.db.Where("slug = ?", roleSlug).First(&role).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"status":  "error",
					"message": "role not found",
				})
			}
			return err
		}

		var perm authModel.RolePermission
		if err := m.db.Where("role_id = ? AND table_name = ?", role.ID, tableName).
			First(&perm).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"status":  "error",
					"message": "permission not configured",
				})
			}
			return err
		}

		allowed := false
		switch normalizedAction {
		case "create":
			allowed = perm.CanCreate
		case "read":
			allowed = perm.CanRead
		case "update":
			allowed = perm.CanUpdate
		case "delete":
			allowed = perm.CanDelete
		default:
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  "error",
				"message": "unknown action",
			})
		}

		if !allowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  "error",
				"message": "permission denied",
			})
		}

		return c.Next()
	}
}
