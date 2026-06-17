CREATE TABLE links (
    id BIGSERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    original_url TEXT NOT NULL,
    slug VARCHAR(50) NOT NULL,
    click_count BIGINT DEFAULT 0 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    CONSTRAINT fk_links_users FOREIGN KEY (user_id)
        REFERENCES users (id) ON DELETE CASCADE
);