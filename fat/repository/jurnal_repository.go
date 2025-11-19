package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/ryansyahrullah/amk-be/fat/model"
)

// JurnalRepository menyimpan operasi database untuk jurnal umum.
type JurnalRepository struct {
	db *gorm.DB
}

// NewJurnalRepository membuat repository baru.
func NewJurnalRepository(db *gorm.DB) *JurnalRepository {
	return &JurnalRepository{db: db}
}

// Create menyimpan jurnal umum baru.
func (r *JurnalRepository) Create(ctx context.Context, jurnal *model.JurnalUmum) error {
	return r.db.WithContext(ctx).Create(jurnal).Error
}

// List mengembalikan daftar jurnal umum (limit optional).
func (r *JurnalRepository) List(ctx context.Context, limit int) ([]model.JurnalUmum, error) {
	var items []model.JurnalUmum
	query := r.db.WithContext(ctx).Order("trans_date DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if err := query.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
