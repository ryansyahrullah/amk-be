package dto

// File: auth/dto/user_create_request.go
// ---------------------------------------------------------
// DTO untuk request pembuatan user baru dari modul AUTH.
// Body JSON ini akan dipakai untuk endpoint seperti POST /auth/users.
// Field di sini biasanya:
//   - nrp, email, full_name, password, role_slug
//
// Saat handler memproses DTO ini, service akan:
//   - Membuat record baru di au_users
//   - Sekaligus membuat record baru di hc_pegawai (otomatis).
//
// Dipakai di:
//   - auth/handler/user_handler.go
//   - auth/service/user_service.go
