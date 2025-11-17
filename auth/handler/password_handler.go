package handler

// File: auth/handler/password_handler.go
// ---------------------------------------------------------
// Handler untuk endpoint lupa password dan reset password.
// Contoh endpoint:
//   - POST /auth/forgot-password
//   - POST /auth/reset-password
//
// Alur:
//   - forgot-password:
//       * Parse DTO PasswordForgotRequest
//       * Panggil password_service untuk generate OTP & kirim email
//   - reset-password:
//       * Parse DTO PasswordResetRequest
//       * Panggil password_service untuk validasi OTP & ganti password.
