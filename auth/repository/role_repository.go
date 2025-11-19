package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/ryansyahrullah/amk-be/auth/model"
)

// RoleRepository menangani query untuk tabel au_roles.
type RoleRepository struct {
	db *gorm.DB
}

// NewRoleRepository membuat instance baru.
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

// FindByID mencari role berdasarkan ID.
func (r *RoleRepository) FindByID(ctx context.Context, id uint) (*model.Role, error) {
	var role model.Role
	if err := r.db.WithContext(ctx).First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// FindBySlug mencari role berdasarkan slug.
func (r *RoleRepository) FindBySlug(ctx context.Context, slug string) (*model.Role, error) {
	var role model.Role
	if err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// Create menambahkan role baru.
func (r *RoleRepository) Create(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

// Update memperbarui role.
func (r *RoleRepository) Update(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Save(role).Error
}

// ListAll mengembalikan seluruh role.
func (r *RoleRepository) ListAll(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	if err := r.db.WithContext(ctx).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
