package service

// File: auth/service/user_service.go
// ---------------------------------------------------------
// Service ini mengurus logic bisnis untuk CRUD user.
// Tanggung jawab utamanya:
//   - Menambahkan user baru:
//       * Membuat record di au_users
//       * Sekaligus membuat record di hc_pegawai (integrasi dengan modul HC-GS)
//   - Mengedit data user (email, nama, role, status aktif)
//   - Menonaktifkan user (soft delete via is_active)
//   - Mengambil daftar user sesuai filter.
//
// Service ini akan memakai:
//   - auth/repository/user_repository.go
//   - auth/repository/role_repository.go
//   - hcgs/repository/pegawai_repository.go (untuk sinkronisasi data pegawai).
