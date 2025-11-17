package utils

import "github.com/gofiber/fiber/v2"

// Success mengembalikan response standar ketika operasi berhasil.
func Success(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"data":    data,
		"message": "",
	})
}

// Created merespons ketika resource berhasil dibuat.
func Created(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"data":    data,
		"message": "",
	})
}

// Error mengembalikan pesan kesalahan standar.
func Error(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}
