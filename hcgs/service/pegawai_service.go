package service

import (
	"context"

	"github.com/ryansyahrullah/amk-be/hcgs/dto"
	"github.com/ryansyahrullah/amk-be/hcgs/model"
	"github.com/ryansyahrullah/amk-be/hcgs/repository"
)

// PegawaiService menangani logic profil pegawai.
type PegawaiService struct {
	repo *repository.PegawaiRepository
}

// NewPegawaiService membuat service baru.
func NewPegawaiService(repo *repository.PegawaiRepository) *PegawaiService {
	return &PegawaiService{repo: repo}
}

// GetProfile mengambil profil pegawai berdasarkan user id.
func (s *PegawaiService) GetProfile(ctx context.Context, userID uint) (*dto.PegawaiResponse, error) {
	pegawai, err := s.repo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return mapPegawaiToDTO(pegawai), nil
}

// UpdateProfile memperbarui data pegawai.
func (s *PegawaiService) UpdateProfile(ctx context.Context, userID uint, req dto.PegawaiUpdateRequest) (*dto.PegawaiResponse, error) {
	pegawai, err := s.repo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if req.FullName != "" {
		pegawai.FullName = req.FullName
	}
	if req.Department != "" {
		pegawai.Department = req.Department
	}
	if req.Position != "" {
		pegawai.Position = req.Position
	}
	if req.Phone != "" {
		pegawai.Phone = req.Phone
	}

	if err := s.repo.Update(ctx, pegawai); err != nil {
		return nil, err
	}

	return mapPegawaiToDTO(pegawai), nil
}

func mapPegawaiToDTO(p *model.Pegawai) *dto.PegawaiResponse {
	return &dto.PegawaiResponse{
		ID:         p.ID,
		UserID:     p.UserID,
		NRP:        p.NRP,
		FullName:   p.FullName,
		Department: p.Department,
		Position:   p.Position,
		Phone:      p.Phone,
	}
}
