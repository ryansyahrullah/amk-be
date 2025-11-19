package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/ryansyahrullah/amk-be/hcgs/model"
)

// PegawaiRepository menangani operasi database untuk hc_pegawai.
type PegawaiRepository struct {
	db *gorm.DB
}

// NewPegawaiRepository membuat repository baru.
func NewPegawaiRepository(db *gorm.DB) *PegawaiRepository {
	return &PegawaiRepository{db: db}
}

// Create membuat data pegawai baru.
func (r *PegawaiRepository) Create(ctx context.Context, pegawai *model.Pegawai) error {
	return r.db.WithContext(ctx).Create(pegawai).Error
}

// Update memperbarui data pegawai.
func (r *PegawaiRepository) Update(ctx context.Context, pegawai *model.Pegawai) error {
	return r.db.WithContext(ctx).Save(pegawai).Error
}

// FindByUserID mengambil data pegawai berdasarkan user_id.
func (r *PegawaiRepository) FindByUserID(ctx context.Context, userID uint) (*model.Pegawai, error) {
	var pegawai model.Pegawai
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&pegawai).Error; err != nil {
		return nil, err
	}
	return &pegawai, nil
}
