package model

import "time"

// RolePermission mendefinisikan hak akses role terhadap tabel tertentu.
type RolePermission struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	RoleID    uint      `json:"role_id"`
	Table     string    `gorm:"column:table_name;size:120" json:"table_name"`
	CanCreate bool      `json:"can_create"`
	CanRead   bool      `json:"can_read"`
	CanUpdate bool      `json:"can_update"`
	CanDelete bool      `json:"can_delete"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName memastikan penggunaan tabel au_role_permissions.
func (RolePermission) TableName() string {
	return "au_role_permissions"
}
