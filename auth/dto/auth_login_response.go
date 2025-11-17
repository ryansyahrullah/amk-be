package dto

// AuthLoginResponse adalah bentuk response saat login berhasil.
type AuthLoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
