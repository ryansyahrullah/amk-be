package service

// File: auth/service/role_service.go
// ---------------------------------------------------------
// Service ini fokus mengatur role dan permission.
// Fungsinya:
//   - Menambah role baru (selain role default)
//   - Mengupdate role (nama, deskripsi)
//   - Menghapus role (kalau dibolehkan)
//   - Mengatur permission per role (role X boleh akses tabel Y dengan hak C/R/U/D tertentu).
//
// Menggunakan:
//   - auth/repository/role_repository.go
//   - auth/repository/permission_repository.go
