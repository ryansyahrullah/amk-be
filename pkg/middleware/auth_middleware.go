package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/ryansyahrullah/amk-be/pkg/utils"
)

// AuthMiddleware memvalidasi token JWT dan menyimpan klaim pada context.
type AuthMiddleware struct {
	secret string
}

// NewAuthMiddleware membuat middleware baru.
func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{secret: secret}
}

// Handle adalah middleware Fiber yang memeriksa Authorization header.
func (m *AuthMiddleware) Handle(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "missing bearer token",
		})
	}

	token := strings.TrimSpace(strings.TrimPrefix(header, "Bearer"))
	claims, err := utils.ParseToken(m.secret, token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid token",
		})
	}

	c.Locals("user_id", claims.UserID)
	c.Locals("nrp", claims.NRP)
	c.Locals("role_slug", claims.RoleSlug)

	return c.Next()
}
