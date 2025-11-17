package dto

// UserUpdateRequest digunakan ketika admin mengubah data user.
type UserUpdateRequest struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	RoleSlug string `json:"role_slug"`
	IsActive *bool  `json:"is_active"`
}
