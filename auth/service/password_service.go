package service

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"errors"
	"strings"
	"time"

	"github.com/ryansyahrullah/amk-be/auth/dto"
	"github.com/ryansyahrullah/amk-be/auth/model"
	"github.com/ryansyahrullah/amk-be/auth/repository"
	"github.com/ryansyahrullah/amk-be/pkg/mail"
	"github.com/ryansyahrullah/amk-be/pkg/utils"
)

// PasswordService mengurus alur lupa dan reset password.
type PasswordService struct {
	users  *repository.UserRepository
	resets *repository.PasswordResetRepository
	mailer *mail.Mailer
	otpTTL time.Duration
}

// NewPasswordService membuat PasswordService baru.
func NewPasswordService(users *repository.UserRepository, resets *repository.PasswordResetRepository, mailer *mail.Mailer, otpTTL time.Duration) *PasswordService {
	return &PasswordService{users: users, resets: resets, mailer: mailer, otpTTL: otpTTL}
}

// ForgotPassword membuat OTP dan mengirim ke email user.
func (s *PasswordService) ForgotPassword(ctx context.Context, req dto.PasswordForgotRequest) error {
	user, err := s.users.FindByEmail(ctx, strings.TrimSpace(req.Email))
	if err != nil {
		return err
	}

	otp := generateOTP()
	reset := &model.PasswordReset{
		UserID:    user.ID,
		OTPCode:   otp,
		ExpiresAt: time.Now().Add(s.otpTTL),
	}

	if err := s.resets.Create(ctx, reset); err != nil {
		return err
	}

	body := "Kode OTP lupa password Anda: " + otp
	return s.mailer.Send(user.Email, "Kode OTP AMK", body)
}

// ResetPassword memvalidasi OTP dan mengganti password user.
func (s *PasswordService) ResetPassword(ctx context.Context, req dto.PasswordResetRequest) error {
	user, err := s.users.FindByEmail(ctx, strings.TrimSpace(req.Email))
	if err != nil {
		return err
	}

	reset, err := s.resets.FindLatestValid(ctx, user.ID, req.OTP)
	if err != nil {
		return errors.New("invalid or expired otp")
	}

	hash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	user.PasswordHash = hash

	if err := s.users.Update(ctx, user); err != nil {
		return err
	}

	if err := s.resets.MarkUsed(ctx, reset); err != nil {
		return err
	}

	return nil
}

func generateOTP() string {
	b := make([]byte, 5)
	if _, err := rand.Read(b); err != nil {
		return "000000"
	}
	encoded := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(b)
	if len(encoded) > 6 {
		encoded = encoded[:6]
	}
	return strings.ToUpper(encoded)
}
