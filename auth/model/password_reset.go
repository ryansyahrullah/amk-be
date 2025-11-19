package model

import "time"

// PasswordReset menyimpan OTP untuk proses lupa password.
type PasswordReset struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `json:"user_id"`
	OTPCode   string     `gorm:"size:10" json:"otp_code"`
	ExpiresAt time.Time  `json:"expires_at"`
	UsedAt    *time.Time `json:"used_at"`
	CreatedAt time.Time  `json:"created_at"`
}

// TableName memastikan penggunaan tabel au_password_resets.
func (PasswordReset) TableName() string {
	return "au_password_resets"
}
