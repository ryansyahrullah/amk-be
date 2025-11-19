package dto

import "time"

// JurnalResponse mewakili data jurnal umum yang dikirim ke client.
type JurnalResponse struct {
	ID          uint      `json:"id"`
	ReferenceNo string    `json:"reference_no"`
	Description string    `json:"description"`
	AccountName string    `json:"account_name"`
	Amount      float64   `json:"amount"`
	TransDate   time.Time `json:"trans_date"`
}
