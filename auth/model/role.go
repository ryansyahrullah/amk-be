package model

// File: auth/model/role.go
// ---------------------------------------------------------
// File ini akan menyimpan definisi struct Role untuk modul AUTH.
// Struct Role di sini akan dipetakan ke tabel database:
//   - Nama tabel: au_roles
//   - Kolom penting: id, name, slug, description, created_at, updated_at
//
// Role digunakan untuk menentukan:
//   - Jenis pengguna (superadmin, pegawai, admin_hcgs, admin_fat, direktur, dll)
//   - Keterkaitan dengan permission (lihat file permission.go)
//   - Relasi dengan User (lihat user.go --> field RoleID).
//
// File ini akan sering dipakai oleh:
//   - auth/repository/role_repository.go   (operasi database terkait role)
//   - auth/service/role_service.go         (logic bisnis role & permission)
//   - pkg/middleware/permission_middleware.go (cek hak akses berdasarkan role).
