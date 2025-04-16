CREATE TABLE IF NOT EXISTS load_production (
    id SERIAL PRIMARY KEY,
    send_block_id INT NOT NULL,
    worker_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0
);