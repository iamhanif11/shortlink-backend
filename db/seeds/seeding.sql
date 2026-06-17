-- ====================================================
-- 1. DUMMY DATA UNTUK TABEL USERS
-- ====================================================
-- Password di bawah ini menggunakan contoh teks biasa. 
-- Pada aplikasi asli, teks ini harus berupa hash Bcrypt[cite: 70, 80].
INSERT INTO users (email, password, name, photo, job) VALUES
('hanif@gmail.com.com', 'Password123', 'M. Hanif Irfan', 'https://api.dicebear.com/7.x/avataaars/svg?seed=Hanif', 'Fullstack Developer'),
('budi@gmail.com.com', 'Password123', 'Budi Santoso', 'https://api.dicebear.com/7.x/avataaars/svg?seed=Budi', 'Social Media Specialist'),
('siti@gmail.com.com', 'Password123', 'Siti Aminah', 'https://api.dicebear.com/7.x/avataaars/svg?seed=Siti', 'Graphic Designer'),
('koda36@gmail.com.com', 'Password123', 'Koda 36', 'https://api.dicebear.com/7.x/avataaars/svg?seed=Admin', 'DevOps Engineer');


-- ====================================================
-- 2. DUMMY DATA UNTUK TABEL LINKS
-- ====================================================
-- Menggunakan ID user 1, 2, dan 3 yang baru saja dibuat di atas[cite: 74].
INSERT INTO links (user_id, original_url, slug, click_count, created_at, deleted_at) VALUES
-- User 1 (Hanif) - Link Aktif
(1, 'https://github.com/hanif/project-fullstack-golang', 'go-bootcamp', 150, NOW() - INTERVAL '5 days', NULL),
(1, 'https://w3schools.com/go/default.asp', 'w3-go', 42, NOW() - INTERVAL '3 days', NULL),
(1, 'https://tailwindcss.com/docs/installation', 'tw-docs', 12, NOW(), NULL),
-- User 1 (Hanif) - Contoh Link yang di-Soft Delete 
(1, 'https://example.com/expired-portfolio-v1', 'port-old', 5, NOW() - INTERVAL '10 days', NOW() - INTERVAL '1 day'),

-- User 2 (Budi) - Link Aktif dengan Random Slug [cite: 21]
(2, 'https://www.youtube.com/watch?v=dQw4w9WgXcQ', 'aB3x9K', 2400, NOW() - INTERVAL '30 days', NULL),
(2, 'https://canva.com/design/template-social-media', 'canva-sm', 89, NOW() - INTERVAL '12 days', NULL),
(2, 'https://instagram.com/socialvit.id', 'ig-svit', 512, NOW() - INTERVAL '7 days', NULL),

-- User 3 (Siti) - Link Aktif & Soft Delete
(3, 'https://behance.net/siti-design-portfolio', 'behance', 310, NOW() - INTERVAL '15 days', NULL),
(3, 'https://pinterest.com/ideas-graphic-design', 'pin-idea', 75, NOW() - INTERVAL '4 days', NULL),
(3, 'https://example.com/deleted-test-link', 'test-del', 0, NOW() - INTERVAL '1 day', NOW());