package middleware

// File: pkg/middleware/permission_middleware.go
// ---------------------------------------------------------
// Middleware untuk mengecek hak akses (permission) berdasarkan role.
// Cara kerja (konsep):
//   - Di route kamu panggil: Permission("hc_pegawai", "read")
//   - Middleware akan:
//       * Mengambil role_slug dari context (di-set oleh JWTAuth)
//       * Mengambil role_id dari au_roles
//       * Mengecek au_role_permissions untuk role_id & table_name
//       * Melihat apakah action (create/read/update/delete) diizinkan.
//
// Jika tidak diizinkan, middleware akan mengembalikan 403 Forbidden.
//
// Middleware ini menghubungkan modul auth (role & permission) dengan
// modul lain (hcgs & fat).
