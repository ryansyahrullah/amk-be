package repository

// File: auth/repository/permission_repository.go
// ---------------------------------------------------------
// File ini fokus untuk berinteraksi dengan tabel au_role_permissions.
// Fungsi-fungsi di sini akan membantu:
//   - Mencari permission berdasarkan role_id dan table_name
//   - Menyimpan konfigurasi permission baru
//   - Mengupdate permission
//
// Repository ini akan digunakan oleh:
//   - auth/service/role_service.go (untuk pengaturan hak akses per role)
//   - pkg/middleware/permission_middleware.go (mengecek apakah role boleh create/read/update/delete).
