INSERT INTO categories (name, description, created_at, updated_at) VALUES
('Sayuran', 'Kategori untuk berbagai jenis sayuran', NOW(), NULL),
('Protein', 'Kategori untuk sumber protein seperti daging, telur, dan kacang-kacangan', NOW(), NULL),
('Buah', 'Kategori untuk berbagai jenis buah-buahan', NOW(), NULL),
('Snack', 'Kategori untuk makanan ringan', NOW(), NULL);

INSERT INTO products (name, price, stock, category_id, created_at, updated_at) VALUES
('Bayam Segar', 8000.00, 50, 1, NOW(), NULL),
('Wortel Organik', 12000.00, 40, 1, NOW(), NULL),
('Dada Ayam Fillet', 35000.00, 25, 2, NOW(), NULL),
('Telur Ayam Kampung', 28000.00, 60, 2, NOW(), NULL),
('Apel Fuji', 30000.00, 35, 3, NOW(), NULL),
('Pisang Cavendish', 18000.00, 45, 3, NOW(), NULL),
('Keripik Kentang', 15000.00, 70, 4, NOW(), NULL),
('Coklat Batang', 22000.00, 55, 4, NOW(), NULL),
('Brokoli Segar', 20000.00, 30, 1, NOW(), NULL),
('Susu Protein', 45000.00, 20, 2, NOW(), NULL);