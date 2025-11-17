package dto

// PasswordForgotRequest adalah payload untuk forgot password.
type PasswordForgotRequest struct {
	Email string `json:"email"`
}
