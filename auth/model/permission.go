package model

// File: auth/model/permission.go
// ---------------------------------------------------------
// File ini akan menyimpan definisi struct RolePermission.
// Struct RolePermission dipetakan ke tabel:
//   - Nama tabel: au_role_permissions
//   - Kolom: id, role_id, table_name, can_create, can_read, can_update, can_delete, created_at, updated_at
//
// Tujuan tabel ini:
//   - Mengatur role tertentu boleh melakukan operasi CRUD apa pada tabel tertentu.
//   - Contoh:
//       * Pegawai:  can_read = 1, can_update = 0 pada table_name = "hc_pegawai"
//       * Admin HC-GS: can_create/update/delete pada "hc_pegawai"
//       * Direktur: hanya can_read pada "hc_pegawai" dan "fa_jurnal_umum"
//
// File ini sangat berkaitan dengan:
//   - auth/model/role.go (karena ada role_id)
//   - auth/repository/permission_repository.go (query DB permission)
//   - pkg/middleware/permission_middleware.go (validasi hak akses berdasarkan role & tabel).
