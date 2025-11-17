package model

import "time"

// User mewakili akun yang dapat login ke aplikasi.
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	NRP          string    `gorm:"size:20;uniqueIndex" json:"nrp"`
	Email        string    `gorm:"size:120;uniqueIndex" json:"email"`
	PasswordHash string    `gorm:"size:255" json:"-"`
	FullName     string    `gorm:"size:120" json:"full_name"`
	RoleID       uint      `json:"role_id"`
	Role         Role      `json:"role"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TableName memastikan GORM menggunakan nama tabel au_users.
func (User) TableName() string {
	return "au_users"
}
