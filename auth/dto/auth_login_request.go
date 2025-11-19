package dto

// AuthLoginRequest mewakili body untuk proses login.
type AuthLoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
