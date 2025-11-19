-- File: database/seeders/001_seed_roles.sql
-- ---------------------------------------------------------
-- Seeder untuk mengisi data awal tabel au_roles.
-- Role default yang bisa diisi:
--   - superadmin
--   - admin_hcgs
--   - admin_fat
--   - direktur
--   - pegawai
--
-- File ini membantu agar kamu punya role dasar tanpa perlu input manual.

INSERT INTO au_roles (name, slug, description, created_at, updated_at)
VALUES
  ('Super Admin', 'superadmin', 'Akses penuh ke seluruh modul', NOW(), NOW()),
  ('Admin HC-GS', 'admin_hcgs', 'Kelola data pegawai', NOW(), NOW()),
  ('Admin FAT', 'admin_fat', 'Kelola modul finance & accounting', NOW(), NOW()),
  ('Direktur', 'direktur', 'Akses laporan tingkat direktur', NOW(), NOW()),
  ('Pegawai', 'pegawai', 'Akun pegawai umum', NOW(), NOW())
ON DUPLICATE KEY UPDATE
  name = VALUES(name),
  description = VALUES(description),
  updated_at = VALUES(updated_at);
