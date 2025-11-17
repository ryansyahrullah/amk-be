# AMK-BE

Backend monolith sederhana berbasis [Fiber](https://github.com/gofiber/fiber) dan MySQL. Kode dibagi ke tiga modul besar:

- **auth** – autentikasi, manajemen user, role, dan permission.
- **hcgs** – data pegawai (Human Capital & General Services).
- **fat** – jurnal umum modul Finance & Accounting.

## Menjalankan aplikasi

```bash
cp .env.example .env # sesuaikan konfigurasi
export $(grep -v '^#' .env | xargs) # opsional ketika tidak menggunakan docker
GO111MODULE=on go run cmd/api/main.go
```

Pastikan database MySQL sudah tersedia sesuai variabel `DB_*` di `.env`.

## Variabel lingkungan penting

| Variabel | Deskripsi |
| --- | --- |
| `APP_PORT` | Port HTTP Fiber (default `3000`). |
| `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME` | Konfigurasi MySQL. |
| `JWT_SECRET`, `JWT_TTL` | Secret dan masa berlaku token JWT. |
| `SMTP_HOST`, `SMTP_PORT`, `SMTP_USER`, `SMTP_PASS`, `SMTP_FROM` | Kredensial SMTP untuk OTP lupa password. |

## Dokumentasi API

### 1. Autentikasi & Manajemen User (`/auth`)

| Method & Path | Deskripsi |
| --- | --- |
| `POST /auth/login` | Login menggunakan `identifier` (NRP/email) dan `password`. Response berupa token JWT dan data user. |
| `POST /auth/forgot-password` | Kirim OTP ke email user. Body: `{ "email": "..." }`. |
| `POST /auth/reset-password` | Ganti password menggunakan OTP. Body: `{ "email": "...", "otp": "ABC123", "new_password": "..." }`. |
| `GET /auth/users` | (Perlu JWT + permission `au_users:read`) Menampilkan daftar user dengan query optional `?q=` dan `?role=`. |
| `POST /auth/users` | (Perlu `au_users:create`) Membuat user sekaligus data pegawai. Body mengikuti `auth/dto/user_create_request.go`. |
| `GET /auth/users/:id` | (Perlu `au_users:read`) Detail user. |
| `PUT /auth/users/:id` | (Perlu `au_users:update`) Update email, nama, role, atau status aktif. |
| `GET /auth/roles` | (Perlu `au_roles:read`) Daftar role. |
| `POST /auth/roles` | (Perlu `au_roles:create`) Membuat role baru. |
| `PUT /auth/roles/:id` | (Perlu `au_roles:update`) Update nama/desk role. |
| `PUT /auth/roles/:id/permissions` | (Perlu `au_role_permissions:update`) Atur hak akses CRUD per tabel. Body berupa array `PermissionRequest`. |

Semua endpoint bertanda *perlu permission* wajib mengirim header `Authorization: Bearer <token>` dari endpoint login.

### 2. Modul HC-GS (`/hcgs`)

| Method & Path | Deskripsi |
| --- | --- |
| `GET /hcgs/pegawai/me` | Mengambil profil pegawai terkait token JWT aktif. |
| `PUT /hcgs/pegawai/me` | Memperbarui profil sendiri. Membutuhkan permission `hc_pegawai:update`. |

### 3. Modul FAT (`/fat`)

| Method & Path | Deskripsi |
| --- | --- |
| `GET /fat/jurnal-umum` | Menampilkan jurnal umum terbaru (opsional `?limit=`). Requires permission `fa_jurnal_umum:read`. |
| `POST /fat/jurnal-umum` | Membuat jurnal umum baru (body sesuai `fat/dto/jurnal_request.go`). Requires permission `fa_jurnal_umum:create`. |

## Struktur Direktori Singkat

```
cmd/api        # entrypoint fiber server
config         # loader .env + koneksi database
pkg            # utilitas (hash, jwt, middleware, mailer)
auth           # modul autentikasi & user management
hcgs           # modul data pegawai
fat            # modul jurnal umum
```

Untuk detail struktur payload lihat berkas DTO pada masing-masing modul.
