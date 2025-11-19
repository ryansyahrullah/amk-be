package model

import "time"

// Role mendeskripsikan peran pengguna dalam sistem.
type Role struct {
	ID          uint             `gorm:"primaryKey" json:"id"`
	Name        string           `gorm:"size:80" json:"name"`
	Slug        string           `gorm:"size:80;uniqueIndex" json:"slug"`
	Description string           `gorm:"size:255" json:"description"`
	Permissions []RolePermission `json:"permissions"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

// TableName memastikan penggunaan tabel au_roles.
func (Role) TableName() string {
	return "au_roles"
}
