package dto

// PegawaiResponse mewakili data pegawai yang dikembalikan ke client.
type PegawaiResponse struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	NRP        string `json:"nrp"`
	FullName   string `json:"full_name"`
	Department string `json:"department"`
	Position   string `json:"position"`
	Phone      string `json:"phone"`
}
