package db

// File: pkg/db/mysql.go
// ---------------------------------------------------------
// (Opsional) File ini bisa digunakan jika kamu ingin memisahkan logika
// koneksi database dari config/config.go.
//
// Misalnya di sini disediakan fungsi:
//   - NewMySQLConnection() -> *gorm.DB
//
// Saat ini, kamu bisa memilih:
//   - Pakai config.InitDB() saja, atau
//   - Memindahkan sebagian logic koneksi DB ke file ini.
