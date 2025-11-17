package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/ryansyahrullah/amk-be/auth"
	"github.com/ryansyahrullah/amk-be/config"
	"github.com/ryansyahrullah/amk-be/fat"
	"github.com/ryansyahrullah/amk-be/hcgs"
)

func main() {
	cfg := config.MustLoad()
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	app := fiber.New()

	auth.RegisterRoutes(app, db, cfg)
	hcgs.RegisterRoutes(app, db, cfg)
	fat.RegisterRoutes(app, db, cfg)

	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatalf("fiber server stopped: %v", err)
	}
}
