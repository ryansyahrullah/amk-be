package hcgs

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/ryansyahrullah/amk-be/config"
	"github.com/ryansyahrullah/amk-be/hcgs/handler"
	"github.com/ryansyahrullah/amk-be/hcgs/repository"
	"github.com/ryansyahrullah/amk-be/hcgs/service"
	"github.com/ryansyahrullah/amk-be/pkg/middleware"
)

// RegisterRoutes mengikat endpoint modul HC-GS.
func RegisterRoutes(app *fiber.App, db *gorm.DB, cfg *config.AppConfig) {
	repo := repository.NewPegawaiRepository(db)
	svc := service.NewPegawaiService(repo)
	h := handler.NewPegawaiHandler(svc)

	authMW := middleware.NewAuthMiddleware(cfg.JWTSecret)
	permMW := middleware.NewPermissionMiddleware(db)

	group := app.Group("/hcgs", authMW.Handle)
	group.Get("/pegawai/me", h.GetMe)
	group.Put("/pegawai/me", permMW.Require("hc_pegawai", "update"), h.UpdateMe)
}
