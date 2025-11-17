package service

import (
	"context"
	"errors"
	"time"

	"github.com/ryansyahrullah/amk-be/auth/dto"
	"github.com/ryansyahrullah/amk-be/auth/repository"
	"github.com/ryansyahrullah/amk-be/pkg/utils"
)

// AuthService menangani proses autentikasi.
type AuthService struct {
	users     *repository.UserRepository
	jwtSecret string
	jwtTTL    int64
}

// NewAuthService membuat service autentikasi baru.
func NewAuthService(users *repository.UserRepository, secret string, ttlSeconds int64) *AuthService {
	return &AuthService{
		users:     users,
		jwtSecret: secret,
		jwtTTL:    ttlSeconds,
	}
}

// Login memverifikasi kredensial dan mengembalikan token.
func (s *AuthService) Login(ctx context.Context, req dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	user, err := s.users.FindByIdentifier(ctx, req.Identifier)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return nil, errors.New("user is inactive")
	}

	if err := utils.CheckPassword(user.PasswordHash, req.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(s.jwtSecret, user.ID, user.NRP, user.Role.Slug, durationFromSeconds(s.jwtTTL))
	if err != nil {
		return nil, err
	}

	resp := &dto.AuthLoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:       user.ID,
			NRP:      user.NRP,
			Email:    user.Email,
			FullName: user.FullName,
			Role: dto.RoleSummary{
				ID:          user.Role.ID,
				Name:        user.Role.Name,
				Slug:        user.Role.Slug,
				Description: user.Role.Description,
			},
			IsActive: user.IsActive,
		},
	}
	return resp, nil
}

func durationFromSeconds(seconds int64) time.Duration {
	if seconds <= 0 {
		return 12 * time.Hour
	}
	return time.Duration(seconds) * time.Second
}
