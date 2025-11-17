package auth

// File: auth/routes.go
// ---------------------------------------------------------
// File ini bertugas mendaftarkan semua route HTTP untuk modul AUTH.
// Contoh route yang akan di-register di sini:
//   - POST /auth/login            -> handler.AuthLogin
//   - POST /auth/users            -> handler.CreateUser
//   - GET  /auth/users            -> handler.ListUser (kalau nanti dibuat)
//   - CRUD role & permission      -> handler.Role*
//   - POST /auth/forgot-password  -> handler.ForgotPassword
//   - POST /auth/reset-password   -> handler.ResetPassword
//
// routes.go akan dipanggil dari main.go dengan sesuatu seperti:
//   auth.RegisterRoutes(app)
// sehingga struktur route auth terpusat dan rapi di satu file ini.
