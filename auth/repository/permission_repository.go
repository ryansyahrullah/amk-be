package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/ryansyahrullah/amk-be/auth/model"
)

// PermissionRepository menangani tabel au_role_permissions.
type PermissionRepository struct {
	db *gorm.DB
}

// NewPermissionRepository membuat repository baru.
func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

// FindByRoleAndTable mencari permission berdasarkan role dan nama tabel.
func (r *PermissionRepository) FindByRoleAndTable(ctx context.Context, roleID uint, table string) (*model.RolePermission, error) {
	var perm model.RolePermission
	if err := r.db.WithContext(ctx).Where("role_id = ? AND table_name = ?", roleID, table).First(&perm).Error; err != nil {
		return nil, err
	}
	return &perm, nil
}

// ListByRole mengembalikan seluruh permission untuk role tertentu.
func (r *PermissionRepository) ListByRole(ctx context.Context, roleID uint) ([]model.RolePermission, error) {
	var perms []model.RolePermission
	if err := r.db.WithContext(ctx).Where("role_id = ?", roleID).Find(&perms).Error; err != nil {
		return nil, err
	}
	return perms, nil
}

// Save melakukan create atau update permission.
func (r *PermissionRepository) Save(ctx context.Context, perm *model.RolePermission) error {
	return r.db.WithContext(ctx).Save(perm).Error
}
