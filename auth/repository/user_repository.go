package repository

// File: auth/repository/user_repository.go
// ---------------------------------------------------------
// File ini berisi operasi database untuk tabel au_users.
// Contoh fungsi:
//   - FindByID(id)
//   - FindByNRP(nrp)
//   - FindByEmail(email)
//   - FindByIdentifier(identifier nrp/ email)
//   - Create(user)
//   - Update(user)
//   - ListWithFilter(...)
//
//
// Repository ini akan dipakai oleh:
//   - auth/service/auth_service.go (login)
//   - auth/service/user_service.go (CRUD user)
//   - auth/service/password_service.go (cari user berdasarkan email).
