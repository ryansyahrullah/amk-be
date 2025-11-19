package model

import "time"

// JurnalUmum menyimpan catatan transaksi sederhana modul FAT.
type JurnalUmum struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ReferenceNo string    `gorm:"size:50;uniqueIndex" json:"reference_no"`
	Description string    `gorm:"size:255" json:"description"`
	AccountName string    `gorm:"size:120" json:"account_name"`
	Amount      float64   `json:"amount"`
	TransDate   time.Time `json:"trans_date"`
	CreatedByID uint      `json:"created_by_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName memastikan sinkron dengan tabel fa_jurnal_umum.
func (JurnalUmum) TableName() string {
	return "fa_jurnal_umum"
}
