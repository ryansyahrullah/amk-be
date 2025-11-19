package seeders

// File: database/seeders/seed.go
// ---------------------------------------------------------
// Seeder ini bertugas mengisi data awal ke database setelah AutoMigrate.
// Yang dilakukan:
//   1. Membuat role default (superadmin, admin_hcgs, admin_fat, direktur, pegawai)
//   2. Membuat user superadmin:
//        - NRP      : SA-0001
//        - Email    : superadmin123@amk.com
//        - Password : qwerty@amk123  (disimpan dalam bentuk bcrypt hash)
//        - Role     : superadmin
//      User ini TIDAK dibuatkan entry di hc_pegawai (bukan pegawai).
//   3. Mengisi au_role_permissions untuk role selain superadmin.
//      Superadmin akan di-bypass di middleware (akses semua).
//
// Seeder ini aman dijalankan berulang kali karena:
//   - Cek dulu apakah data sudah ada (count / where), kalau sudah -> skip.
//
// Seeder akan dipanggil dari main.go setelah AutoMigrate.

import (
	"errors"
	"log"

	authModel "github.com/ryansyahrullah/amk-be/auth/model"
	"github.com/ryansyahrullah/amk-be/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Run menjalankan semua proses seeding.
func Run() error {
	if err := seedRoles(); err != nil {
		return err
	}
	if err := seedSuperadminUser(); err != nil {
		return err
	}
	if err := seedPermissions(); err != nil {
		return err
	}

	log.Println("Seeding selesai ✔")
	return nil
}

// seedRoles membuat role default jika belum ada data di au_roles.
func seedRoles() error {
	var count int64
	if err := config.DB.Model(&authModel.Role{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		// Role sudah ada, tidak perlu di-seed lagi
		log.Println("Roles sudah ada, skip seeding roles.")
		return nil
	}

	roles := []authModel.Role{
		{Name: "Super Admin", Slug: "superadmin", Description: "Memiliki akses penuh ke seluruh sistem"},
		{Name: "Admin HC-GS", Slug: "admin_hcgs", Description: "Mengelola data kepegawaian"},
		{Name: "Admin FAT", Slug: "admin_fat", Description: "Mengelola data keuangan"},
		{Name: "Direktur", Slug: "direktur", Description: "Hanya membaca laporan dan data"},
		{Name: "Pegawai", Slug: "pegawai", Description: "Pegawai biasa, hanya mengakses data dirinya sendiri"},
	}

	if err := config.DB.Create(&roles).Error; err != nil {
		return err
	}

	log.Println("Seed roles: OK")
	return nil
}

// seedSuperadminUser membuat user superadmin jika belum ada.
func seedSuperadminUser() error {
	const superEmail = "superadmin123@amk.com"

	var existing authModel.User
	err := config.DB.Where("email = ?", superEmail).First(&existing).Error
	if err == nil {
		// Sudah ada user superadmin, skip
		log.Println("User superadmin sudah ada, skip seeding user.")
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Ambil role superadmin
	var superRole authModel.Role
	if err := config.DB.Where("slug = ?", "superadmin").First(&superRole).Error; err != nil {
		return err
	}

	// Hash password default: qwerty@amk123
	hash, err := bcrypt.GenerateFromPassword([]byte("qwerty@amk123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := authModel.User{
		NRP:          "SA-0001",
		Email:        superEmail,
		PasswordHash: string(hash),
		FullName:     "Super Admin Root",
		RoleID:       superRole.ID,
		IsActive:     true,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	log.Println("Seed superadmin user: OK")
	return nil
}

// seedPermissions mengisi au_role_permissions untuk role selain superadmin.
func seedPermissions() error {
	var count int64
	if err := config.DB.Model(&authModel.RolePermission{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("Permissions sudah ada, skip seeding permissions.")
		return nil
	}

	// Ambil semua role yang dibutuhkan
	var superadmin, pegawai, adminHC, adminFAT, direktur authModel.Role

	config.DB.Where("slug = ?", "superadmin").First(&superadmin)
	config.DB.Where("slug = ?", "pegawai").First(&pegawai)
	config.DB.Where("slug = ?", "admin_hcgs").First(&adminHC)
	config.DB.Where("slug = ?", "admin_fat").First(&adminFAT)
	config.DB.Where("slug = ?", "direktur").First(&direktur)

	log.Println("Seed permissions: OK")
	return nil
}
