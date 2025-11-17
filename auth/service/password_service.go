package service

// File: auth/service/password_service.go
// ---------------------------------------------------------
// Service ini mengurus alur lupa password dan reset password.
// Tugasnya:
//   - Saat lupa password:
//       * Cek email user
//       * Generate OTP
//       * Simpan ke au_password_resets
//       * Meminta pkg/mail/mailer.go untuk mengirim email OTP
//   - Saat reset password:
//       * Validasi OTP (belum expired, belum dipakai)
//       * Ganti password di au_users (hash baru)
//       * Menandai OTP sebagai "used".
//
// Menggunakan:
//   - auth/repository/user_repository.go
//   - auth/repository/password_reset_repository.go
//   - pkg/utils/hash.go
//   - pkg/mail/mailer.go
