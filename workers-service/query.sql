-- name: AddWorker :exec
INSERT INTO workers
    (first_name, last_name, phone)
VALUES 
    ($1, $2, $3);


-- name: GetWorkers :many
SELECT 
    id, first_name, last_name, phone, created_at, updated_at, deleted_at
FROM
    workers
WHERE
    deleted_at = 0;

-- name: EndDay :one
INSERT INTO daily_production
    (date, count_blocks)
VALUES 
    ($1, $2)
RETURNING id;

-- name: EndDayWorkers :exec
INSERT INTO daily_production_workers
    (daily_production_id, worker_id)
VALUES 
    ($1, $2);

-- name: SendBlocks :one
INSERT INTO send_blocks
    (date, count_blocks, address, load_price)
VALUES 
    ($1, $2, $3, $4)
RETURNING id;

-- name: LoadBlockWorkers :exec
INSERT INTO load_production
    (send_block_id, worker_id)
VALUES 
    ($1, $2);

