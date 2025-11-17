package handler

// File: auth/handler/user_handler.go
// ---------------------------------------------------------
// Berisi handler untuk CRUD user.
// Contoh endpoint:
//   - POST   /auth/users       -> CreateUser
//   - GET    /auth/users       -> ListUser
//   - GET    /auth/users/:id   -> GetUserByID
//   - PUT    /auth/users/:id   -> UpdateUser
//   - DELETE /auth/users/:id   -> Delete/DeactivateUser
//
// Handler di sini akan:
//   - Parse request ke DTO (CreateUserRequest, UserUpdateRequest)
//   - Panggil auth/service/user_service.go
//   - Mengembalikan response JSON ke client.
