package dto

// RoleSummary merangkum informasi role di response.
type RoleSummary struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

// UserResponse mewakili struktur user pada response API.
type UserResponse struct {
	ID       uint        `json:"id"`
	NRP      string      `json:"nrp"`
	Email    string      `json:"email"`
	FullName string      `json:"full_name"`
	Role     RoleSummary `json:"role"`
	IsActive bool        `json:"is_active"`
}
