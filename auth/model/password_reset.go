package model

// File: auth/model/password_reset.go
// ---------------------------------------------------------
// File ini akan menyimpan definisi struct PasswordReset.
// Struct ini dipetakan ke tabel:
//   - Nama tabel: au_password_resets
//   - Kolom: id, user_id, otp_code, expires_at, used_at, created_at
//
// Tujuan tabel ini:
//   - Menyimpan kode OTP untuk proses lupa kata sandi.
//   - Mencatat kapan OTP dibuat, kapan expired, dan kapan sudah digunakan.
//   - Menghubungkan OTP dengan user tertentu (user_id -> au_users.id).
//
// File ini digunakan oleh:
//   - auth/repository/password_reset_repository.go (membuat dan mengambil OTP)
//   - auth/service/password_service.go (logic lupa password & reset password)
//   - auth/handler/password_handler.go (endpoint /auth/forgot-password dan /auth/reset-password).
