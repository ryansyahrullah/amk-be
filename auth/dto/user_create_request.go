package dto

// UserCreateRequest dipakai untuk endpoint pembuatan user baru.
type UserCreateRequest struct {
	NRP        string `json:"nrp"`
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	Password   string `json:"password"`
	RoleSlug   string `json:"role_slug"`
	Department string `json:"department"`
	Position   string `json:"position"`
}
