package dto

// PegawaiUpdateRequest digunakan untuk memperbarui profil pegawai.
type PegawaiUpdateRequest struct {
	FullName   string `json:"full_name"`
	Department string `json:"department"`
	Position   string `json:"position"`
	Phone      string `json:"phone"`
}
