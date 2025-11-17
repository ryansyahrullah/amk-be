package main

// File: cmd/tools/hash_superadmin.go
// ---------------------------------------------------------
// Tujuan file ini:
//   - Menghasilkan bcrypt hash dari password default superadmin
//     yaitu: "qwerty@amk123".
//   - Hash ini akan digunakan di seeder:
//         database/seeders/003_seed_superadmin.sql
//     pada kolom password_hash.
//
// Kenapa perlu file ini?
//   - Supaya kita tidak menulis password plain text di seeder SQL.
//   - Bcrypt hash lebih aman dan sesuai dengan cara sistem login
//     memvalidasi password.
//
// Cara pakai:
//   1. Pastikan sudah install package bcrypt:
//        go get golang.org/x/crypto/bcrypt
//   2. Jalankan perintah dari root project (amk-be):
//        go run cmd/tools/hash_superadmin.go
//   3. Program akan mencetak hash di terminal, contoh:
//        Bcrypt hash untuk password superadmin: $2a$10$....
//   4. Copy hash tersebut, lalu paste ke file
//        database/seeders/003_seed_superadmin.sql
//      menggantikan teks: <BCRYPT_HASH_DI_SINI>
//   5. Setelah itu, file ini boleh kamu hapus kalau sudah tidak diperlukan.
//
// Catatan:
//   - File ini TIDAK dijalankan oleh aplikasi utama.
//   - Ini murni tool kecil untuk membantu proses setup awal.

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "qwerty@amk123"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bcrypt hash untuk password superadmin:", string(hash))
}
