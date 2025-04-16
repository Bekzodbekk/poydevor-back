CREATE TABLE IF NOT EXISTS send_blocks (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    count_blocks INT NOT NULL,
    address VARCHAR(255) NOT NULL,
    load_price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0
);