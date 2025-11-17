package utils

// File: pkg/utils/jwt.go
// ---------------------------------------------------------
// Utility untuk membuat dan membaca JWT token.
// Fungsi yang mungkin ada:
//   - GenerateToken(userID, nrp, roleSlug string) (string, error)
//   - ParseToken(tokenString string) (*JWTClaims, error)
//
// JWTClaims bisa menyimpan:
//   - user_id
//   - nrp
//   - role_slug
//   - Expiry time, IssuedAt, dll.
//
// Dipakai oleh:
//   - auth/service/auth_service.go (generate token saat login)
//   - pkg/middleware/auth_middleware.go (parse token dari request).
