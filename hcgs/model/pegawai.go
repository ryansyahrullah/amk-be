package hcgs

// File: hcgs/routes.go
// ---------------------------------------------------------
// File ini mendaftarkan semua route untuk modul HC-GS (Human Capital / General Services).
// Contoh route:
//   - GET /hcgs/pegawai/me      -> handler.GetMe (pegawai lihat data diri sendiri)
//   - (nanti) CRUD data pegawai untuk admin HC-GS.
//
// routes.go akan dipanggil dari main.go, misalnya:
//   hcgs.RegisterRoutes(app)
