package model

// File: auth/model/user.go
// ---------------------------------------------------------
// File ini akan menyimpan definisi struct User untuk modul AUTH.
// Struct User di sini akan dipetakan ke tabel database:
//   - Nama tabel: au_users
//   - Kolom penting: id, nrp, email, password_hash, full_name, role_id, is_active, created_at, updated_at
//
// User mewakili akun yang bisa login ke sistem:
//   - Pegawai (role: pegawai)
//   - Admin HC-GS (role: admin_hcgs)
//   - Admin FAT (role: admin_fat)
//   - Direktur (role: direktur)
//   - Superadmin (role: superadmin)
//
// File ini terhubung dengan:
//   - Role (lihat role.go) melalui role_id
//   - hcgs/model/pegawai.go, karena saat membuat user baru, otomatis juga dibuat data pegawai.
//   - auth/repository/user_repository.go untuk operasi DB (CRUD user)
//   - auth/service/auth_service.go (proses login, cek password, buat token).
