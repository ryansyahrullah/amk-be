-- File: database/migrations/002_create_hcgs_tables.sql
-- ---------------------------------------------------------
-- Migration ini membuat tabel-tabel untuk modul HC-GS.
-- Minimal:
--   - hc_pegawai
--
-- Tabel ini akan menampung data utama pegawai yang terhubung dengan au_users.nrp.

CREATE TABLE IF NOT EXISTS hc_pegawai (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id BIGINT UNSIGNED NOT NULL,
  nrp VARCHAR(30) NOT NULL,
  full_name VARCHAR(150) NOT NULL,
  birth_place VARCHAR(100) NULL,
  birth_date DATE NULL,
  gender ENUM('L', 'P') NULL,
  department VARCHAR(120) NULL,
  division VARCHAR(120) NULL,
  position VARCHAR(120) NULL,
  grade VARCHAR(30) NULL,
  employment_status ENUM('permanent', 'contract', 'intern') NULL,
  join_date DATE NULL,
  phone_number VARCHAR(30) NULL,
  email_company VARCHAR(150) NULL,
  address TEXT NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uq_hc_pegawai_user_id (user_id),
  UNIQUE KEY uq_hc_pegawai_nrp (nrp),
  CONSTRAINT fk_hc_pegawai_user FOREIGN KEY (user_id) REFERENCES au_users (id)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
