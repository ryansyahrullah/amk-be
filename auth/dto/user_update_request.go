package dto

// File: auth/dto/user_update_request.go
// ---------------------------------------------------------
// DTO untuk request update data user (misalnya update role atau status aktif).
// Digunakan di endpoint seperti PUT/PATCH /auth/users/:id.
// Field yang umum: email, full_name, role_slug, is_active.
//
// Dipakai di:
//   - auth/handler/user_handler.go
//   - auth/service/user_service.go
