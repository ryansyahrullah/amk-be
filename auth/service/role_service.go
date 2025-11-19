package service

import (
	"context"

	"github.com/ryansyahrullah/amk-be/auth/dto"
	"github.com/ryansyahrullah/amk-be/auth/model"
	"github.com/ryansyahrullah/amk-be/auth/repository"
)

// RoleService fokus pada pengaturan role dan permission.
type RoleService struct {
	roles       *repository.RoleRepository
	permissions *repository.PermissionRepository
}

// NewRoleService membuat RoleService baru.
func NewRoleService(roles *repository.RoleRepository, permissions *repository.PermissionRepository) *RoleService {
	return &RoleService{roles: roles, permissions: permissions}
}

// ListRoles mengembalikan seluruh role yang ada.
func (s *RoleService) ListRoles(ctx context.Context) ([]model.Role, error) {
	return s.roles.ListAll(ctx)
}

// CreateRole menambahkan role baru.
func (s *RoleService) CreateRole(ctx context.Context, name, slug, description string) (*model.Role, error) {
	role := &model.Role{Name: name, Slug: slug, Description: description}
	if err := s.roles.Create(ctx, role); err != nil {
		return nil, err
	}
	return role, nil
}

// UpdateRole memperbarui nama/desk role.
func (s *RoleService) UpdateRole(ctx context.Context, id uint, name, description string) (*model.Role, error) {
	role, err := s.roles.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if name != "" {
		role.Name = name
	}
	if description != "" {
		role.Description = description
	}
	if err := s.roles.Update(ctx, role); err != nil {
		return nil, err
	}
	return role, nil
}

// UpsertPermissions menyimpan konfigurasi permission per role.
func (s *RoleService) UpsertPermissions(ctx context.Context, roleID uint, req []dto.PermissionRequest) ([]model.RolePermission, error) {
	results := make([]model.RolePermission, 0, len(req))
	for _, item := range req {
		perm := model.RolePermission{
			RoleID:    roleID,
			Table:     item.TableName,
			CanCreate: item.CanCreate,
			CanRead:   item.CanRead,
			CanUpdate: item.CanUpdate,
			CanDelete: item.CanDelete,
		}
		if err := s.permissions.Save(ctx, &perm); err != nil {
			return nil, err
		}
		results = append(results, perm)
	}
	return results, nil
}
