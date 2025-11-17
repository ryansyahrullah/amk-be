package handler

// File: auth/handler/role_handler.go
// ---------------------------------------------------------
// Menangani HTTP endpoint terkait Role & Permission.
// Contoh endpoint:
//   - GET    /auth/roles
//   - POST   /auth/roles
//   - PUT    /auth/roles/:id
//   - DELETE /auth/roles/:id
//   - GET/PUT /auth/roles/:id/permissions
//
// Handler ini akan menggunakan auth/service/role_service.go untuk mengatur
// role baru, mengupdate, dan mengatur hak akses tabel per role.
