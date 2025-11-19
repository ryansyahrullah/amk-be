-- File: database/migrations/001_create_auth_tables.sql
-- ---------------------------------------------------------
-- Migration ini berisi perintah SQL untuk membuat tabel-tabel modul AUTH:
--   - au_roles
--   - au_users
--   - au_role_permissions
--   - au_password_resets
--
-- Tujuan:
--   - Menyiapkan struktur awal database untuk autentikasi & otorisasi.

CREATE TABLE IF NOT EXISTS au_roles (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  slug VARCHAR(50) NOT NULL,
  description TEXT NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uq_au_roles_slug (slug)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS au_users (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  nrp VARCHAR(30) NOT NULL,
  email VARCHAR(150) NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  full_name VARCHAR(150) NOT NULL,
  role_id BIGINT UNSIGNED NOT NULL,
  is_active TINYINT(1) NOT NULL DEFAULT 1,
  last_login_at DATETIME NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uq_au_users_nrp (nrp),
  UNIQUE KEY uq_au_users_email (email),
  KEY idx_au_users_role_id (role_id),
  CONSTRAINT fk_au_users_role FOREIGN KEY (role_id) REFERENCES au_roles (id)
    ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS au_role_permissions (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  role_id BIGINT UNSIGNED NOT NULL,
  table_name VARCHAR(100) NOT NULL,
  can_create TINYINT(1) NOT NULL DEFAULT 0,
  can_read TINYINT(1) NOT NULL DEFAULT 0,
  can_update TINYINT(1) NOT NULL DEFAULT 0,
  can_delete TINYINT(1) NOT NULL DEFAULT 0,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uq_au_role_permissions_role_table (role_id, table_name),
  KEY idx_au_role_permissions_role_id (role_id),
  CONSTRAINT fk_au_role_permissions_role FOREIGN KEY (role_id) REFERENCES au_roles (id)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS au_password_resets (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id BIGINT UNSIGNED NOT NULL,
  token CHAR(64) NOT NULL,
  expires_at DATETIME NOT NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uq_au_password_resets_token (token),
  KEY idx_au_password_resets_user_id (user_id),
  CONSTRAINT fk_au_password_resets_user FOREIGN KEY (user_id) REFERENCES au_users (id)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
