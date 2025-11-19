package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/ryansyahrullah/amk-be/auth"
	authModel "github.com/ryansyahrullah/amk-be/auth/model"
	"github.com/ryansyahrullah/amk-be/config"
	"github.com/ryansyahrullah/amk-be/database/seeders"
	fatModel "github.com/ryansyahrullah/amk-be/fat/model"
	hcModel "github.com/ryansyahrullah/amk-be/hcgs/model"
)

func main() {
	_ = godotenv.Load()

	// init DB
	config.InitDB()

	// (kalau kamu punya fungsi khusus buat AppConfig, pakai itu.
	// Di sini contoh simple: load dari env)
	cfg := &config.AppConfig{
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTTTL:    72 * time.Hour,

		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPPort: 1025, // kalau mau, parse dari env
		SMTPUser: os.Getenv("SMTP_USER"),
		SMTPPass: os.Getenv("SMTP_PASS"),
		SMTPFrom: os.Getenv("SMTP_FROM"),
	}

	// AutoMigrate
	if err := config.DB.AutoMigrate(
		&authModel.Role{},
		&authModel.User{},
		&authModel.RolePermission{},
		&authModel.PasswordReset{},
		&hcModel.Pegawai{},
		&fatModel.JurnalUmum{},
	); err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}

	// Seeder (kalau pakai)
	if err := seeders.Run(); err != nil {
		log.Fatalf("Seeding error: %v", err)
	}

	app := fiber.New()

	// ⚠️ PENTING: DAFTARIN ROUTE AUTH DI SINI
	auth.RegisterRoutes(app, config.DB, cfg)

	// route test
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "AMK-BE API running"})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
