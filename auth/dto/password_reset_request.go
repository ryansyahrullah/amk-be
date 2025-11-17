package dto

// PasswordResetRequest adalah payload untuk mengganti kata sandi.
type PasswordResetRequest struct {
	Email       string `json:"email"`
	OTP         string `json:"otp"`
	NewPassword string `json:"new_password"`
}
