package dto

// PermissionRequest mendeskripsikan hak akses role pada satu tabel.
type PermissionRequest struct {
	TableName string `json:"table_name"`
	CanCreate bool   `json:"can_create"`
	CanRead   bool   `json:"can_read"`
	CanUpdate bool   `json:"can_update"`
	CanDelete bool   `json:"can_delete"`
}
