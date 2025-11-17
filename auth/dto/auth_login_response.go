package dto

// File: auth/dto/auth_login_response.go
// ---------------------------------------------------------
// DTO untuk response login.
// Struktur di file ini akan menentukan bentuk JSON yang dikembalikan ke client
// setelah login berhasil, misalnya:
//   {
//     "token": "jwt-token-...",
//     "user": {
//       "id": 1,
//       "nrp": "...",
//       "email": "...",
//       "full_name": "...",
//       "role": "pegawai"
//     }
//   }
//
// Dipakai di:
//   - auth/handler/auth_handler.go setelah proses login sukses.
