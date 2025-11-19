-- File: database/seeders/002_seed_permissions.sql
-- ---------------------------------------------------------
-- Isi default permission untuk setiap role terhadap tabel penting.
--       * pegawai: read-only di hc_pegawai (data diri sendiri via endpoint /me)
--       * admin_hcgs: full CRUD di hc_pegawai
--       * admin_fat: full CRUD di fa_jurnal_umum, read di hc_pegawai
--       * direktur: read-only di hc_pegawai dan fa_jurnal_umum

INSERT INTO au_role_permissions (role_id, table_name, can_create, can_read, can_update, can_delete, created_at, updated_at)
SELECT r.id, 'hc_pegawai', 1, 1, 1, 1, NOW(), NOW()
FROM au_roles r
WHERE r.slug = 'superadmin'
ON DUPLICATE KEY UPDATE
  can_create = VALUES(can_create),
  can_read = VALUES(can_read),
  can_update = VALUES(can_update),
  can_delete = VALUES(can_delete),
  updated_at = VALUES(updated_at);

INSERT INTO au_role_permissions (role_id, table_name, can_create, can_read, can_update, can_delete, created_at, updated_at)
SELECT r.id, 'fa_jurnal_umum', 1, 1, 1, 1, NOW(), NOW()
FROM au_roles r
WHERE r.slug = 'superadmin'
ON DUPLICATE KEY UPDATE
  can_create = VALUES(can_create),
  can_read = VALUES(can_read),
  can_update = VALUES(can_update),
  can_delete = VALUES(can_delete),
  updated_at = VALUES(updated_at);

INSERT INTO au_role_permissions (role_id, table_name, can_create, can_read, can_update, can_delete, created_at, updated_at)
SELECT r.id, 'hc_pegawai', 1, 1, 1, 1, NOW(), NOW()
FROM au_roles r
WHERE r.slug = 'admin_hcgs'
ON DUPLICATE KEY UPDATE
  can_create = VALUES(can_create),
  can_read = VALUES(can_read),
  can_update = VALUES(can_update),
  can_delete = VALUES(can_delete),
  updated_at = VALUES(updated_at);

INSERT INTO au_role_permissions (role_id, table_name, can_create, can_read, can_update, can_delete, created_at, updated_at)
SELECT r.id, 'hc_pegawai', 0, 1, 0, 0, NOW(), NOW()
FROM au_roles r
WHERE r.slug = 'pegawai'
ON DUPLICATE KEY UPDATE
  can_create = VALUES(can_create),
  can_read = VALUES(can_read),
  can_update = VALUES(can_update),
  can_delete = VALUES(can_delete),
  updated_at = VALUES(updated_at);

INSERT INTO au_role_permissions (role_id, table_name, can_create, can_read, can_update, can_delete, created_at, updated_at)
SELECT r.id, 'hc_pegawai', 0, 1, 0, 0, NOW(), NOW()
FROM au_roles r
WHERE r.slug = 'direktur'
ON DUPLICATE KEY UPDATE
  can_create = VALUES(can_create),
  can_read = VALUES(can_read),
  can_update = VALUES(can_update),
  can_delete = VALUES(can_delete),
  updated_at = VALUES(updated_at);

INSERT INTO au_role_permissions (role_id, table_name, can_create, can_read, can_update, can_delete, created_at, updated_at)
SELECT r.id, 'hc_pegawai', 0, 1, 0, 0, NOW(), NOW()
FROM au_roles r
WHERE r.slug = 'admin_fat'
ON DUPLICATE KEY UPDATE
  can_create = VALUES(can_create),
  can_read = VALUES(can_read),
  can_update = VALUES(can_update),
  can_delete = VALUES(can_delete),
  updated_at = VALUES(updated_at);

INSERT INTO au_role_permissions (role_id, table_name, can_create, can_read, can_update, can_delete, created_at, updated_at)
SELECT r.id, 'fa_jurnal_umum', 1, 1, 1, 1, NOW(), NOW()
FROM au_roles r
WHERE r.slug = 'admin_fat'
ON DUPLICATE KEY UPDATE
  can_create = VALUES(can_create),
  can_read = VALUES(can_read),
  can_update = VALUES(can_update),
  can_delete = VALUES(can_delete),
  updated_at = VALUES(updated_at);

INSERT INTO au_role_permissions (role_id, table_name, can_create, can_read, can_update, can_delete, created_at, updated_at)
SELECT r.id, 'fa_jurnal_umum', 0, 1, 0, 0, NOW(), NOW()
FROM au_roles r
WHERE r.slug = 'direktur'
ON DUPLICATE KEY UPDATE
  can_create = VALUES(can_create),
  can_read = VALUES(can_read),
  can_update = VALUES(can_update),
  can_delete = VALUES(can_delete),
  updated_at = VALUES(updated_at);
