package handler

// File: auth/handler/auth_handler.go
// ---------------------------------------------------------
// Berisi HTTP handler untuk proses login (dan mungkin logout).
// Contoh endpoint yang di-handle:
//   - POST /auth/login
//
// Alur umum:
//   - Parse body JSON ke dto.LoginRequest
//   - Panggil auth/service/auth_service.go untuk validasi user & password
//   - Jika sukses, balikan token JWT + data user dalam format dto.AuthLoginResponse.
//
// File ini menjadi penghubung antara dunia HTTP (Fiber) dengan logic di auth_service.
