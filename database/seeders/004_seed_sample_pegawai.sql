-- File: database/seeders/004_seed_sample_pegawai.sql
-- ---------------------------------------------------------
-- Seeder ini membuat contoh user pegawai beserta data hc_pegawai yang
-- terhubung langsung melalui kolom user_id.

INSERT INTO au_users (nrp, email, password_hash, full_name, role_id, is_active, created_at, updated_at)
SELECT
  'PGW-0001' AS nrp,
  'pegawai001@amk.com' AS email,
  '$2a$10$RruagWdarzpeDVFa.MbFQ.C6jyiw.yT86rDF5AF7LXmWMbf2SaxP2' AS password_hash,
  'Pegawai Contoh Pertama' AS full_name,
  r.id AS role_id,
  1 AS is_active,
  NOW() AS created_at,
  NOW() AS updated_at
FROM au_roles r
WHERE r.slug = 'pegawai'
ON DUPLICATE KEY UPDATE
  email = VALUES(email),
  password_hash = VALUES(password_hash),
  full_name = VALUES(full_name),
  role_id = VALUES(role_id),
  is_active = VALUES(is_active),
  updated_at = VALUES(updated_at);

INSERT INTO hc_pegawai (
  user_id,
  nrp,
  full_name,
  birth_place,
  birth_date,
  gender,
  department,
  division,
  position,
  grade,
  employment_status,
  join_date,
  phone_number,
  email_company,
  address,
  created_at,
  updated_at)
SELECT
  u.id,
  u.nrp,
  u.full_name,
  'Bandung' AS birth_place,
  '1995-07-04' AS birth_date,
  'L' AS gender,
  'Human Capital' AS department,
  'HC-GS' AS division,
  'Staff HC' AS position,
  '7A' AS grade,
  'permanent' AS employment_status,
  '2020-02-17' AS join_date,
  '+62-812-0000-0001' AS phone_number,
  CONCAT('pegawai.', u.nrp, '@amk.com') AS email_company,
  'Jl. Mawar No. 10, Bandung' AS address,
  NOW() AS created_at,
  NOW() AS updated_at
FROM au_users u
WHERE u.nrp = 'PGW-0001'
ON DUPLICATE KEY UPDATE
  full_name = VALUES(full_name),
  birth_place = VALUES(birth_place),
  birth_date = VALUES(birth_date),
  gender = VALUES(gender),
  department = VALUES(department),
  division = VALUES(division),
  position = VALUES(position),
  grade = VALUES(grade),
  employment_status = VALUES(employment_status),
  join_date = VALUES(join_date),
  phone_number = VALUES(phone_number),
  email_company = VALUES(email_company),
  address = VALUES(address),
  updated_at = VALUES(updated_at);
