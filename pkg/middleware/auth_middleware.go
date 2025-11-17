package middleware

// File: pkg/middleware/auth_middleware.go
// ---------------------------------------------------------
// Middleware untuk autentikasi JWT.
// Tugasnya:
//   - Membaca header Authorization: Bearer <token>
//   - Mem-parse dan memverifikasi JWT menggunakan pkg/utils/jwt.go
//   - Jika valid, menyimpan data penting ke context Fiber, misalnya:
//       * user_id
//       * nrp
//       * role_slug
//   - Jika tidak valid / tidak ada token, balikan 401 Unauthorized.
//
// Middleware ini akan digunakan di route yang butuh login:
//   - /hcgs/pegawai/me
//   - /fat/jurnal-umum
//   - dll.
