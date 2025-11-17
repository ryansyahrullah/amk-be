package dto

// File: auth/dto/password_reset_request.go
// ---------------------------------------------------------
// DTO untuk request reset kata sandi menggunakan OTP.
// Body JSON yang diharapkan, misalnya:
//   {
//     "email": "user@example.com",
//     "otp": "123456",
//     "new_password": "passwordBaru"
//   }
//
// Endpoint terkait:
//   - POST /auth/reset-password
//
// Handler akan menggunakan DTO ini untuk:
//   - Validasi OTP
//   - Mengubah password di au_users
//   - Menandai OTP sebagai sudah digunakan (used_at terisi).
