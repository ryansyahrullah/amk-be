package service

import (
	"context"
	"errors"

	"github.com/ryansyahrullah/amk-be/auth/dto"
	"github.com/ryansyahrullah/amk-be/auth/model"
	"github.com/ryansyahrullah/amk-be/auth/repository"
	hcgsModel "github.com/ryansyahrullah/amk-be/hcgs/model"
	hcgsRepo "github.com/ryansyahrullah/amk-be/hcgs/repository"
	"github.com/ryansyahrullah/amk-be/pkg/utils"
)

// UserService menangani logic bisnis terkait user.
type UserService struct {
	users   *repository.UserRepository
	roles   *repository.RoleRepository
	pegawai *hcgsRepo.PegawaiRepository
}

// NewUserService membuat service baru.
func NewUserService(users *repository.UserRepository, roles *repository.RoleRepository, pegawai *hcgsRepo.PegawaiRepository) *UserService {
	return &UserService{users: users, roles: roles, pegawai: pegawai}
}

// CreateUser membuat akun user baru beserta data pegawai.
func (s *UserService) CreateUser(ctx context.Context, req dto.UserCreateRequest) (*dto.UserResponse, error) {
	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	role, err := s.roles.FindBySlug(ctx, req.RoleSlug)
	if err != nil {
		return nil, err
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		NRP:          req.NRP,
		Email:        req.Email,
		FullName:     req.FullName,
		PasswordHash: hash,
		RoleID:       role.ID,
		IsActive:     true,
	}

	if err := s.users.Create(ctx, user); err != nil {
		return nil, err
	}

	pegawai := &hcgsModel.Pegawai{
		UserID:     user.ID,
		NRP:        req.NRP,
		FullName:   req.FullName,
		Department: req.Department,
		Position:   req.Position,
	}

	if err := s.pegawai.Create(ctx, pegawai); err != nil {
		return nil, err
	}

	user.Role = *role
	return mapUserToDTO(user), nil
}

// UpdateUser memperbarui data user.
func (s *UserService) UpdateUser(ctx context.Context, id uint, req dto.UserUpdateRequest) (*dto.UserResponse, error) {
	user, err := s.users.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.RoleSlug != "" {
		role, err := s.roles.FindBySlug(ctx, req.RoleSlug)
		if err != nil {
			return nil, err
		}
		user.RoleID = role.ID
		user.Role = *role
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := s.users.Update(ctx, user); err != nil {
		return nil, err
	}

	return mapUserToDTO(user), nil
}

// ListUsers menampilkan seluruh user sesuai filter.
func (s *UserService) ListUsers(ctx context.Context, filter repository.UserFilter) ([]dto.UserResponse, error) {
	users, err := s.users.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.UserResponse, 0, len(users))
	for _, u := range users {
		copy := u
		responses = append(responses, *mapUserToDTO(&copy))
	}
	return responses, nil
}

// GetUser mengambil user berdasarkan ID.
func (s *UserService) GetUser(ctx context.Context, id uint) (*dto.UserResponse, error) {
	user, err := s.users.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapUserToDTO(user), nil
}

func mapUserToDTO(user *model.User) *dto.UserResponse {
	return &dto.UserResponse{
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
	}
}
