-- File: database/seeders/003_seed_superadmin.sql
-- ---------------------------------------------------------
-- Seeder ini membuat 1 user superadmin default untuk sistem.
-- Tujuannya:
--   - Memberikan akun awal dengan akses penuh ke seluruh sistem,
--     sehingga admin bisa mengelola user, role, dan permission
--     langsung dari aplikasi tanpa perlu akses database lagi.
--
-- Data user superadmin yang dibuat:
--   - NRP      : SA-0001  (NRP khusus, bukan pegawai biasa)
--   - Email    : superadmin123@amk.com
--   - Password : qwerty@amk123  (disimpan dalam bentuk bcrypt hash)
--   - Role     : superadmin (diambil dari tabel au_roles)
--
-- Catatan:
--   - User ini TIDAK dibuatkan data di tabel hc_pegawai agar
--     jelas bahwa dia bukan pegawai, tetapi akun root sistem.
--   - String '<BCRYPT_HASH_DI_SINI>' harus diganti dengan hash
--     bcrypt asli yang dihasilkan oleh tool Go di:
--       cmd/tools/hash_superadmin.go
--
-- Cara pakai seeder ini:
--   - Pastikan tabel au_roles sudah terisi role 'superadmin'
--     (lihat 001_seed_roles.sql)
--   - Generate hash password, lalu ganti placeholder di bawah
--   - Jalankan seeder ini ke database amk_db.

INSERT INTO au_users (nrp, email, password_hash, full_name, role_id, is_active, created_at, updated_at)
SELECT
  'SA-0001' AS nrp,
  'superadmin123@amk.com' AS email,
  '<BCRYPT_HASH_DI_SINI>' AS password_hash,
  'Super Admin Root' AS full_name,
  id AS role_id,
  1 AS is_active,
  NOW() AS created_at,
  NOW() AS updated_at
FROM au_roles
WHERE slug = 'superadmin';
