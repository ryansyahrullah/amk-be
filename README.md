# AMK-BE

Backend monolith sederhana berbasis [Fiber](https://github.com/gofiber/fiber) dan MySQL. Kode dibagi ke tiga modul besar:

- **auth** – autentikasi, manajemen user, role, dan permission.
- **hcgs** – data pegawai (Human Capital & General Services).
- **fat** – jurnal umum modul Finance & Accounting.

## Menjalankan aplikasi

```bash
cp .env.example .env
# Edit .env bila perlu lalu jalankan dependency
docker compose up -d db mailhog

# Setelah MySQL siap, jalankan API dari host
go run ./cmd/api
```

Pastikan database MySQL bernama `amk_db` sudah tersedia. Perintah `docker compose` di atas akan otomatis membuat database beserta user `amk`/`amk123` dan menyambungkannya ke port 3306 host Anda sehingga `DB_HOST=127.0.0.1` di `.env` dapat langsung dipakai. Bila ingin menjalankan MySQL sendiri (bukan dari compose) cukup pastikan kredensial `.env` sesuai dengan server tersebut.

### Menyesuaikan environment

1. **Salin `.env.example`** ke `.env` lalu ubah nilai `DB_*` jika kredensial MySQL Anda berbeda.
2. **Gunakan Docker Compose** untuk dependency default:
   ```bash
   docker compose up -d db mailhog
   # tunggu healthcheck mysql hijau (docker compose ps)
   ```
3. **Jalankan API** memakai `go run ./cmd/api` atau `CompileDaemon` favorit Anda.

> **Catatan koneksi database**
>
> `.env.example` sekarang menyiapkan `DB_HOST=127.0.0.1` agar cocok untuk workflow lokal (host -> container MySQL). Bila Anda menjalankan API di dalam container compose yang sama, override variabel tersebut menjadi `db` karena nama service MySQL-nya memang `db`.

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

## Menjalankan pengujian

Seluruh paket saat ini belum memiliki unit test spesifik, namun Anda dapat memastikan dependensi terpasang dengan menjalankan:

```bash
go test ./...
```

Perintah tersebut hanya membutuhkan koneksi internet saat pertama kali untuk mengunduh modul Go.
