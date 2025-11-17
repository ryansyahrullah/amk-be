package dto

// File: auth/dto/auth_login_request.go
// ---------------------------------------------------------
// DTO (Data Transfer Object) untuk request login.
// Struktur di file ini akan mewakili body JSON yang dikirim client saat login,
// misalnya:
//   {
//     "identifier": "12345"   // bisa nrp atau email
//     "password": "rahasia"
//   }
//
// File ini akan digunakan oleh:
//   - auth/handler/auth_handler.go untuk parsing body request login dari Fiber.
