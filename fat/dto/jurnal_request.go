package dto

import "time"

// JurnalRequest digunakan untuk membuat entri jurnal umum baru.
type JurnalRequest struct {
	ReferenceNo string    `json:"reference_no"`
	Description string    `json:"description"`
	AccountName string    `json:"account_name"`
	Amount      float64   `json:"amount"`
	TransDate   time.Time `json:"trans_date"`
}
