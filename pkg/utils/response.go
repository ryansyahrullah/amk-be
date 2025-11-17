package utils

// File: pkg/utils/response.go
// ---------------------------------------------------------
// (Opsional) Helper untuk standarisasi bentuk response JSON.
// Misalnya menyediakan fungsi:
//   - Success(c *fiber.Ctx, data interface{})
//   - Error(c *fiber.Ctx, status int, message string)
//
// Tujuannya agar semua response API konsisten dan rapi.
