package model

import "time"

// Pegawai mewakili data personal pegawai yang terkait dengan user.
type Pegawai struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	NRP        string    `gorm:"size:20;uniqueIndex" json:"nrp"`
	FullName   string    `gorm:"size:120" json:"full_name"`
	Department string    `gorm:"size:120" json:"department"`
	Position   string    `gorm:"size:120" json:"position"`
	Phone      string    `gorm:"size:30" json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName memastikan sinkron dengan tabel hc_pegawai.
func (Pegawai) TableName() string {
	return "hc_pegawai"
}
