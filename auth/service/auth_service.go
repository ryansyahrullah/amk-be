package service

// File: auth/service/auth_service.go
// ---------------------------------------------------------
// Service ini menangani logic utama untuk proses autentikasi (login).
// Fungsi di sini akan:
//   - Mencari user berdasarkan identifier (NRP atau email)
//   - Mengecek password dengan bcrypt
//   - Mengambil role user
//   - Membuat JWT token (menggunakan pkg/utils/jwt.go)
//   - Mengembalikan informasi user + token ke handler.
//
// Service ini menghubungkan:
//   - auth/repository/user_repository.go
//   - auth/repository/role_repository.go
//   - pkg/utils/hash.go (cek password)
//   - pkg/utils/jwt.go (generate token).
