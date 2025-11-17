package main

// File: cmd/api/main.go
// ---------------------------------------------------------
// Ini adalah entrypoint utama aplikasi backend AMK-BE.
// Tugas file ini nantinya:
//   - Memuat konfigurasi dari file .env (APP_PORT, DB, dsb)
//   - Memanggil config.InitDB() untuk membuka koneksi MySQL via GORM
//   - Menginisialisasi web server Fiber
//   - Mendaftarkan semua route dari modul:
//       * auth (login, user, role, lupa password)
//       * hcgs (data pegawai)
//       * fat (jurnal umum / modul keuangan)
//   - Menjalankan HTTP server di port yang ditentukan APP_PORT.
// File ini akan "menghubungkan" semua modul agar menjadi satu aplikasi monolith.

func main() {
	// TODO:
	// - Load .env
	// - Panggil config.InitDB()
	// - Buat instance Fiber
	// - Register routes: auth.RegisterRoutes(app), hcgs.RegisterRoutes(app), fat.RegisterRoutes(app)
	// - app.Listen(":" + APP_PORT)
}
