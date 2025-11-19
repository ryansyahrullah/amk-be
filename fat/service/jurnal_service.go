package service

import (
	"context"
	"time"

	"github.com/ryansyahrullah/amk-be/fat/dto"
	"github.com/ryansyahrullah/amk-be/fat/model"
	"github.com/ryansyahrullah/amk-be/fat/repository"
)

// JurnalService menangani logic transaksi jurnal umum.
type JurnalService struct {
	repo *repository.JurnalRepository
}

// NewJurnalService membuat service baru.
func NewJurnalService(repo *repository.JurnalRepository) *JurnalService {
	return &JurnalService{repo: repo}
}

// Create membuat transaksi baru.
func (s *JurnalService) Create(ctx context.Context, req dto.JurnalRequest, createdBy uint) (*dto.JurnalResponse, error) {
	jurnal := &model.JurnalUmum{
		ReferenceNo: req.ReferenceNo,
		Description: req.Description,
		AccountName: req.AccountName,
		Amount:      req.Amount,
		TransDate:   req.TransDate,
		CreatedByID: createdBy,
	}

	if jurnal.TransDate.IsZero() {
		jurnal.TransDate = time.Now()
	}

	if err := s.repo.Create(ctx, jurnal); err != nil {
		return nil, err
	}

	return mapJurnalToDTO(jurnal), nil
}

// List menampilkan daftar jurnal.
func (s *JurnalService) List(ctx context.Context, limit int) ([]dto.JurnalResponse, error) {
	items, err := s.repo.List(ctx, limit)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.JurnalResponse, 0, len(items))
	for _, item := range items {
		copy := item
		responses = append(responses, *mapJurnalToDTO(&copy))
	}
	return responses, nil
}

func mapJurnalToDTO(j *model.JurnalUmum) *dto.JurnalResponse {
	return &dto.JurnalResponse{
		ID:          j.ID,
		ReferenceNo: j.ReferenceNo,
		Description: j.Description,
		AccountName: j.AccountName,
		Amount:      j.Amount,
		TransDate:   j.TransDate,
	}
}
