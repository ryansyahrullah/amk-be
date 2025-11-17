package repository

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"github.com/ryansyahrullah/amk-be/auth/model"
)

// UserRepository menyimpan operasi database untuk tabel au_users.
type UserRepository struct {
	db *gorm.DB
}

// UserFilter menyediakan opsi pencarian user.
type UserFilter struct {
	Query      string
	RoleSlug   string
	OnlyActive *bool
}

// NewUserRepository membuat instance repository baru.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByID mencari user berdasarkan ID.
func (r *UserRepository) FindByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Preload("Role").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByIdentifier mencari user berdasarkan nrp atau email.
func (r *UserRepository) FindByIdentifier(ctx context.Context, identifier string) (*model.User, error) {
	normalized := strings.TrimSpace(identifier)
	var user model.User
	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("nrp = ? OR email = ?", normalized, normalized).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail mencari user berdasarkan email.
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).Preload("Role").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Create menyimpan user baru.
func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// Update memperbarui data user.
func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// List mengembalikan daftar user sesuai filter.
func (r *UserRepository) List(ctx context.Context, filter UserFilter) ([]model.User, error) {
	var users []model.User
	query := r.db.WithContext(ctx).Model(&model.User{}).Preload("Role")

	if filter.Query != "" {
		like := "%" + filter.Query + "%"
		query = query.Where("nrp LIKE ? OR email LIKE ? OR full_name LIKE ?", like, like, like)
	}

	if filter.RoleSlug != "" {
		query = query.Joins("JOIN au_roles ON au_roles.id = au_users.role_id").Where("au_roles.slug = ?", filter.RoleSlug)
	}

	if filter.OnlyActive != nil {
		query = query.Where("au_users.is_active = ?", *filter.OnlyActive)
	}

	if err := query.Order("au_users.created_at DESC").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
