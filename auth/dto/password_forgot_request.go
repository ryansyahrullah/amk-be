package dto

// File: auth/dto/password_forgot_request.go
// ---------------------------------------------------------
// DTO untuk request lupa kata sandi (forgot password).
// Client akan mengirim JSON seperti:
//   { "email": "user@example.com" }
//
// Endpoint terkait:
//   - POST /auth/forgot-password
//
// Handler akan menggunakan DTO ini untuk:
//   - Mencari user berdasarkan email
//   - Membuat OTP di tabel au_password_resets
//   - Mengirimkan OTP ke email user.
