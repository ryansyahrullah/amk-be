package repository

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/ryansyahrullah/amk-be/auth/model"
)

// PasswordResetRepository menangani tabel au_password_resets.
type PasswordResetRepository struct {
	db *gorm.DB
}

// NewPasswordResetRepository membuat repository baru.
func NewPasswordResetRepository(db *gorm.DB) *PasswordResetRepository {
	return &PasswordResetRepository{db: db}
}

// Create membuat token reset password baru.
func (r *PasswordResetRepository) Create(ctx context.Context, reset *model.PasswordReset) error {
	return r.db.WithContext(ctx).Create(reset).Error
}

// FindLatestValid mencari OTP terakhir yang masih berlaku.
func (r *PasswordResetRepository) FindLatestValid(ctx context.Context, userID uint, otp string) (*model.PasswordReset, error) {
	var reset model.PasswordReset
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND otp_code = ? AND used_at IS NULL AND expires_at > ?", userID, otp, time.Now()).
		Order("created_at DESC").
		First(&reset).Error
	if err != nil {
		return nil, err
	}
	return &reset, nil
}

// MarkUsed menandai OTP sebagai sudah digunakan.
func (r *PasswordResetRepository) MarkUsed(ctx context.Context, reset *model.PasswordReset) error {
	now := time.Now()
	reset.UsedAt = &now
	return r.db.WithContext(ctx).Save(reset).Error
}
