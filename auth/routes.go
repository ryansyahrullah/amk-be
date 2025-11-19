package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/ryansyahrullah/amk-be/auth/handler"
	"github.com/ryansyahrullah/amk-be/auth/repository"
	"github.com/ryansyahrullah/amk-be/auth/service"
	"github.com/ryansyahrullah/amk-be/config"
	hcgsRepo "github.com/ryansyahrullah/amk-be/hcgs/repository"
	"github.com/ryansyahrullah/amk-be/pkg/mail"
	"github.com/ryansyahrullah/amk-be/pkg/middleware"
)

// RegisterRoutes mendaftarkan seluruh endpoint modul auth.
func RegisterRoutes(app *fiber.App, db *gorm.DB, cfg *config.AppConfig) {
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	permRepo := repository.NewPermissionRepository(db)
	resetRepo := repository.NewPasswordResetRepository(db)
	pegawaiRepo := hcgsRepo.NewPegawaiRepository(db)

	authService := service.NewAuthService(userRepo, cfg.JWTSecret, int64(cfg.JWTTTL.Seconds()))
	userService := service.NewUserService(userRepo, roleRepo, pegawaiRepo)
	roleService := service.NewRoleService(roleRepo, permRepo)
	mailer := mail.New(mail.Config{
		Host:     cfg.SMTPHost,
		Port:     cfg.SMTPPort,
		Username: cfg.SMTPUser,
		Password: cfg.SMTPPass,
		From:     cfg.SMTPFrom,
	})
	passwordService := service.NewPasswordService(userRepo, resetRepo, mailer, 15*time.Minute)

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	roleHandler := handler.NewRoleHandler(roleService)
	passwordHandler := handler.NewPasswordHandler(passwordService)

	authMW := middleware.NewAuthMiddleware(cfg.JWTSecret)
	permMW := middleware.NewPermissionMiddleware(db)

	group := app.Group("/auth")
	group.Post("/login", authHandler.Login)
	group.Post("/forgot-password", passwordHandler.ForgotPassword)
	group.Post("/reset-password", passwordHandler.ResetPassword)

	protected := group.Group("", authMW.Handle)
	protected.Post("/users", permMW.Require("au_users", "create"), userHandler.Create)
	protected.Get("/users", permMW.Require("au_users", "read"), userHandler.List)
	protected.Get("/users/:id", permMW.Require("au_users", "read"), userHandler.Get)
	protected.Put("/users/:id", permMW.Require("au_users", "update"), userHandler.Update)

	protected.Get("/roles", permMW.Require("au_roles", "read"), roleHandler.List)
	protected.Post("/roles", permMW.Require("au_roles", "create"), roleHandler.Create)
	protected.Put("/roles/:id", permMW.Require("au_roles", "update"), roleHandler.Update)
	protected.Put("/roles/:id/permissions", permMW.Require("au_role_permissions", "update"), roleHandler.UpdatePermissions)
}
