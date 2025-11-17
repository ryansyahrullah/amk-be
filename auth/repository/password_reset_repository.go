package repository

// File: auth/repository/password_reset_repository.go
// ---------------------------------------------------------
// File ini menangani akses database untuk tabel au_password_resets.
// Contoh fungsi:
//   - CreateResetToken(userID, otp, expiresAt)
//   - FindLatestValidByUserAndOTP(userID, otp)
//   - MarkAsUsed(resetID)
//
// Repository ini akan dipakai oleh:
//   - auth/service/password_service.go
//   - auth/handler/password_handler.go (secara tidak langsung melalui service).
