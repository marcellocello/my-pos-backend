-- 000001_init_schema.sql
-- Inisialisasi skema basis data MyPOS Toko Kelontong

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    role ENUM('owner', 'cashier') NOT NULL DEFAULT 'cashier',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Tabel satuan produk (pcs, bal, dus, karton, dll)
CREATE TABLE IF NOT EXISTS units (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE COMMENT 'Nama satuan, contoh: pcs, bal, dus',
    description VARCHAR(100) NULL COMMENT 'Keterangan tambahan satuan'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Data satuan umum toko kelontong
INSERT INTO units (name, description) VALUES
    ('pcs',    'Piece / Satuan Buah'),
    ('bal',    'Bal / Bale (isi banyak pcs)'),
    ('dus',    'Dus / Karton Besar'),
    ('karton', 'Karton'),
    ('lusin',  'Lusin (12 pcs)'),
    ('kodi',   'Kodi (20 pcs)'),
    ('kg',     'Kilogram'),
    ('gram',   'Gram'),
    ('liter',  'Liter'),
    ('ml',     'Mililiter'),
    ('botol',  'Botol'),
    ('kaleng', 'Kaleng'),
    ('pack',   'Pack / Paket'),
    ('sachet', 'Sachet'),
    ('lembar', 'Lembar'),
    ('roll',   'Roll / Gulungan')
ON DUPLICATE KEY UPDATE name = name;

CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    category_id INT NULL,
    unit_id INT NOT NULL DEFAULT 1 COMMENT 'Referensi ke tabel units (default: pcs)',
    sku VARCHAR(50) UNIQUE,
    name VARCHAR(200) NOT NULL,
    cost_price DECIMAL(12, 2) NOT NULL DEFAULT 0.00,
    sell_price DECIMAL(12, 2) NOT NULL DEFAULT 0.00,
    stock INT NOT NULL DEFAULT 0,
    min_stock_alert INT NOT NULL DEFAULT 5,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL,
    FOREIGN KEY (unit_id) REFERENCES units(id),
    INDEX idx_product_name (name),
    INDEX idx_product_sku (sku),
    INDEX idx_product_stock (stock)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    invoice_number VARCHAR(50) NOT NULL UNIQUE,
    user_id INT NOT NULL,
    total_amount DECIMAL(12, 2) NOT NULL DEFAULT 0.00,
    paid_amount DECIMAL(12, 2) NOT NULL DEFAULT 0.00,
    change_amount DECIMAL(12, 2) NOT NULL DEFAULT 0.00,
    payment_method ENUM('cash', 'qris') NOT NULL DEFAULT 'cash',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    INDEX idx_transaction_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS transaction_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    unit_cost_price DECIMAL(12, 2) NOT NULL,
    unit_sell_price DECIMAL(12, 2) NOT NULL,
    subtotal DECIMAL(12, 2) NOT NULL,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id),
    INDEX idx_item_product (product_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS stock_logs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    quantity_change INT NOT NULL,
    type ENUM('sale', 'restock', 'adjustment') NOT NULL,
    reference_id VARCHAR(50) NULL,
    notes TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    INDEX idx_stock_log_product (product_id),
    INDEX idx_stock_log_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
