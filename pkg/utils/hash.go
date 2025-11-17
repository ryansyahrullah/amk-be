package utils

// File: pkg/utils/hash.go
// ---------------------------------------------------------
// Utility untuk hashing dan verifikasi password menggunakan bcrypt.
//
// Fungsi yang umum:
//   - HashPassword(password string) (string, error)
//   - CheckPassword(hash, plain string) error
//
// Dipakai di:
//   - auth/service/auth_service.go (cek password saat login)
//   - auth/service/user_service.go (saat membuat user baru)
//   - auth/service/password_service.go (saat reset password).
