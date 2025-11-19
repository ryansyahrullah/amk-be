package fat

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/ryansyahrullah/amk-be/config"
	"github.com/ryansyahrullah/amk-be/fat/handler"
	"github.com/ryansyahrullah/amk-be/fat/repository"
	"github.com/ryansyahrullah/amk-be/fat/service"
	"github.com/ryansyahrullah/amk-be/pkg/middleware"
)

// RegisterRoutes memasang endpoint modul FAT (keuangan).
func RegisterRoutes(app *fiber.App, db *gorm.DB, cfg *config.AppConfig) {
	repo := repository.NewJurnalRepository(db)
	svc := service.NewJurnalService(repo)
	h := handler.NewJurnalHandler(svc)

	authMW := middleware.NewAuthMiddleware(cfg.JWTSecret)
	permMW := middleware.NewPermissionMiddleware(db)

	group := app.Group("/fat", authMW.Handle)
	group.Get("/jurnal-umum", permMW.Require("fa_jurnal_umum", "read"), h.List)
	group.Post("/jurnal-umum", permMW.Require("fa_jurnal_umum", "create"), h.Create)
}
