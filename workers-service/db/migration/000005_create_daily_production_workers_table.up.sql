CREATE TABLE IF NOT EXISTS daily_production_workers (
    id SERIAL PRIMARY KEY,
    daily_production_id INT NOT NULL,
    worker_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0
);