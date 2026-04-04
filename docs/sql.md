-- ============================================
-- 📚 SIPALA Database Schema for MySQL
-- ============================================

-- Hapus database jika sudah ada (opsional, hati-hati!)
-- DROP DATABASE IF EXISTS sipala-database;

-- Buat database
CREATE DATABASE IF NOT EXISTS sipala-database 
    CHARACTER SET utf8mb4 
    COLLATE utf8mb4_unicode_ci;

USE sipala database;

-- ============================================
-- 🔹 Enums (Dibuat sebagai tabel referensi)
-- ============================================

-- Tabel untuk user roles
CREATE TABLE user_roles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    role_name VARCHAR(20) NOT NULL UNIQUE,
    description VARCHAR(100)
) ENGINE=InnoDB;

INSERT INTO user_roles (role_name, description) VALUES
('admin', 'Full akses ke sistem'),
('staff', 'Approve & monitoring'),
('borrower', 'Peminjam');

-- Tabel untuk borrow status
CREATE TABLE borrow_statuses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    status_name VARCHAR(20) NOT NULL UNIQUE,
    description VARCHAR(100)
) ENGINE=InnoDB;

INSERT INTO borrow_statuses (status_name, description) VALUES
('pending', 'Menunggu approval'),
('approved', 'Disetujui'),
('rejected', 'Ditolak'),
('returned', 'Sudah dikembalikan');

-- ============================================
-- 👤 Tabel Users
-- ============================================

CREATE TABLE users (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    FOREIGN KEY (role) REFERENCES user_roles(role_name),
    INDEX idx_email (email),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB;

-- Trigger untuk memastikan UUID generate jika NULL
DELIMITER //
CREATE TRIGGER before_insert_users
BEFORE INSERT ON users
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
END//
DELIMITER ;

-- ============================================
-- 🏷️ Tabel Categories
-- ============================================

CREATE TABLE categories (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_name (name)
) ENGINE=InnoDB;

-- Trigger untuk UUID
DELIMITER //
CREATE TRIGGER before_insert_categories
BEFORE INSERT ON categories
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
END//
DELIMITER ;

-- ============================================
-- 🧰 Tabel Tools
-- ============================================

CREATE TABLE tools (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    name VARCHAR(255) NOT NULL,
    category_id CHAR(36),
    stock INT NOT NULL DEFAULT 0 CHECK (stock >= 0),
    description TEXT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL,
    INDEX idx_category_id (category_id),
    INDEX idx_deleted_at (deleted_at),
    INDEX idx_name (name)
) ENGINE=InnoDB;

-- Trigger untuk UUID
DELIMITER //
CREATE TRIGGER before_insert_tools
BEFORE INSERT ON tools
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
END//
DELIMITER ;

-- ============================================
-- 📥 Tabel Borrow Transactions (Header)
-- ============================================

CREATE TABLE borrow_transactions (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    borrower_id CHAR(36) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    borrow_date TIMESTAMP NOT NULL,
    due_date TIMESTAMP NOT NULL,
    approved_by CHAR(36) NULL,
    approved_at TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (borrower_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (status) REFERENCES borrow_statuses(status_name),
    FOREIGN KEY (approved_by) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT chk_due_date CHECK (due_date >= borrow_date),
    
    INDEX idx_borrower_id (borrower_id),
    INDEX idx_status (status),
    INDEX idx_approved_by (approved_by),
    INDEX idx_borrow_date (borrow_date)
) ENGINE=InnoDB;

-- Trigger untuk UUID
DELIMITER //
CREATE TRIGGER before_insert_borrow_transactions
BEFORE INSERT ON borrow_transactions
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
END//
DELIMITER ;

-- ============================================
-- 📦 Tabel Borrow Transaction Items (Detail)
-- ============================================

CREATE TABLE borrow_transaction_items (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    borrow_transaction_id CHAR(36) NOT NULL,
    tool_id CHAR(36) NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (borrow_transaction_id) REFERENCES borrow_transactions(id) ON DELETE CASCADE,
    FOREIGN KEY (tool_id) REFERENCES tools(id) ON DELETE CASCADE,
    
    INDEX idx_borrow_transaction_id (borrow_transaction_id),
    INDEX idx_tool_id (tool_id)
) ENGINE=InnoDB;

-- Trigger untuk UUID
DELIMITER //
CREATE TRIGGER before_insert_borrow_transaction_items
BEFORE INSERT ON borrow_transaction_items
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
END//
DELIMITER ;

-- ============================================
-- 📤 Tabel Return Transactions
-- ============================================

CREATE TABLE return_transactions (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    borrow_transaction_id CHAR(36) NOT NULL UNIQUE,
    returned_at TIMESTAMP NOT NULL,
    late_days INT NOT NULL DEFAULT 0 CHECK (late_days >= 0),
    fine_amount DECIMAL(10, 2) NOT NULL DEFAULT 0.00 CHECK (fine_amount >= 0),
    processed_by CHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (borrow_transaction_id) REFERENCES borrow_transactions(id) ON DELETE CASCADE,
    FOREIGN KEY (processed_by) REFERENCES users(id) ON DELETE CASCADE,
    
    INDEX idx_borrow_transaction_id (borrow_transaction_id),
    INDEX idx_processed_by (processed_by)
) ENGINE=InnoDB;

-- Trigger untuk UUID
DELIMITER //
CREATE TRIGGER before_insert_return_transactions
BEFORE INSERT ON return_transactions
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
END//
DELIMITER ;

-- ============================================
-- 📜 Tabel Activity Logs
-- ============================================

CREATE TABLE activity_logs (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    user_id CHAR(36),
    action VARCHAR(50) NOT NULL,
    entity VARCHAR(50) NOT NULL,
    entity_id CHAR(36) NULL,
    description TEXT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    
    INDEX idx_user_id (user_id),
    INDEX idx_action (action),
    INDEX idx_entity (entity),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB;

-- Trigger untuk UUID
DELIMITER //
CREATE TRIGGER before_insert_activity_logs
BEFORE INSERT ON activity_logs
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
END//
DELIMITER ;

-- REFRESH TOKENS

CREATE TABLE refresh_tokens (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    user_id CHAR(36),
    token TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    revoked_at TIMESTAMP NULL DEFAULT NULL,
    
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    INDEX idx_user_id (user_id),
    INDEX idx_revoked_at (revoked_at),
) ENGINE=InnoDB;

-- Trigger untuk UUID
DELIMITER //
CREATE TRIGGER before_insert_tokens
BEFORE INSERT ON refresh_tokens
FOR EACH ROW
BEGIN
    IF NEW.id IS NULL THEN
        SET NEW.id = UUID();
    END IF;
END//
DELIMITER ;

-- ============================================
-- 🔥 Views untuk Mempermudah Query
-- ============================================

-- View untuk melihat transaksi dengan detail borrower
CREATE VIEW view_borrow_transactions AS
SELECT 
    bt.id,
    bt.borrower_id,
    u.name AS borrower_name,
    u.email AS borrower_email,
    bt.status,
    bt.borrow_date,
    bt.due_date,
    bt.approved_by,
    approver.name AS approved_by_name,
    bt.approved_at,
    bt.created_at,
    bt.updated_at
FROM borrow_transactions bt
LEFT JOIN users u ON bt.borrower_id = u.id
LEFT JOIN users approver ON bt.approved_by = approver.id;

-- View untuk melihat items dalam transaksi
CREATE VIEW view_borrow_items AS
SELECT 
    bti.id,
    bti.borrow_transaction_id,
    bti.tool_id,
    t.name AS tool_name,
    c.name AS category_name,
    bti.quantity,
    bti.created_at
FROM borrow_transaction_items bti
JOIN tools t ON bti.tool_id = t.id
LEFT JOIN categories c ON t.category_id = c.id;

-- View untuk melihat summary pengembalian
CREATE VIEW view_return_summary AS
SELECT 
    rt.id,
    rt.borrow_transaction_id,
    bt.borrower_id,
    u.name AS borrower_name,
    bt.borrow_date,
    bt.due_date,
    rt.returned_at,
    rt.late_days,
    rt.fine_amount,
    rt.processed_by,
    processor.name AS processed_by_name,
    rt.created_at
FROM return_transactions rt
JOIN borrow_transactions bt ON rt.borrow_transaction_id = bt.id
LEFT JOIN users u ON bt.borrower_id = u.id
LEFT JOIN users processor ON rt.processed_by = processor.id;

-- ============================================
-- ✅ Selesai!
-- ============================================

-- Cek semua tabel yang dibuat
SHOW TABLES;