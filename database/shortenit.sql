CREATE TABLE IF NOT EXISTS links(
    id BIGSERIAL PRIMARY KEY,
    link_id VARCHAR(11) NOT NULL,
    original_url TEXT NOT NULL,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);