package config

// File: config/config.go
// ---------------------------------------------------------
// Package config menyimpan hal-hal terkait konfigurasi global aplikasi.
// File ini nantinya akan:
//   - Membaca variabel environment (DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME, dll)
//   - Membuat koneksi database MySQL menggunakan GORM
//   - Menyediakan variabel global seperti:
//       * DB *gorm.DB  --> digunakan oleh repository di modul auth, hcgs, dan fat
//   - (Opsional) Menyimpan konfigurasi lain seperti mode aplikasi (APP_ENV), dsb.
//
// File ini menjadi "jembatan" antara environment (.env, docker-compose) dengan kode Go.
