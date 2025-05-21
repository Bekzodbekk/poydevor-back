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



-- name: EndDayDataMonthlyReport :many
WITH worker_blocks AS (
    -- Get all production IDs where the worker participated
    SELECT 
        dp.id,
        dpw.daily_production_id,
        dp.date,
        dp.count_blocks
    FROM 
        daily_production_workers dpw
    JOIN 
        daily_production dp ON dpw.daily_production_id = dp.id
    WHERE 
        dpw.worker_id = CAST($1 AS INTEGER)
        AND dpw.deleted_at = 0
        AND dp.deleted_at = 0
        AND EXTRACT(YEAR FROM dp.date) = CAST($2 AS INTEGER)
        AND EXTRACT(MONTH FROM dp.date) = CAST($3 AS INTEGER)
),
production_stats AS (
    -- Count workers for each production
    SELECT 
        dpw.daily_production_id,
        COUNT(dpw.worker_id) AS worker_count
    FROM 
        daily_production_workers dpw
    WHERE 
        dpw.deleted_at = 0
    GROUP BY 
        dpw.daily_production_id
)
SELECT 
    wb.daily_production_id,
    wb.date,
    wb.count_blocks AS total_blocks,
    ps.worker_count,
    ROUND(wb.count_blocks::NUMERIC / ps.worker_count, 1) AS worker_share,
    ROUND((wb.count_blocks::NUMERIC / ps.worker_count) * 600, 1) AS worker_payment
FROM 
    worker_blocks wb
JOIN 
    production_stats ps ON wb.daily_production_id = ps.daily_production_id
ORDER BY 
    wb.date;

-- name: LoadBlocksDataMonthlyReport :many
WITH worker_payments AS (
    SELECT
        sb.id,
        lp.worker_id,
        sb.date,
        sb.count_blocks AS total_blocks,
        sb.address,
        COUNT(DISTINCT lp2.worker_id) AS worker_count,
        ROUND(sb.count_blocks::NUMERIC / COUNT(DISTINCT lp2.worker_id), 1) AS blocks_per_worker,
        sb.load_price AS price_per_block,
        ROUND((sb.count_blocks::NUMERIC / COUNT(DISTINCT lp2.worker_id)) * sb.load_price, 1) AS payment
    FROM 
        load_production lp
    JOIN 
        send_blocks sb ON lp.send_block_id = sb.id
    JOIN 
        load_production lp2 ON lp.send_block_id = lp2.send_block_id
    WHERE 
        lp.worker_id = CAST($1 AS INTEGER)
        AND EXTRACT(YEAR FROM sb.date) = CAST($2 AS INTEGER)
        AND EXTRACT(MONTH FROM sb.date) = CAST($3 AS INTEGER)
    GROUP BY 
        lp.worker_id, sb.id, sb.date, sb.count_blocks, sb.address, sb.load_price
)
SELECT 
    id as send_block_id,
    worker_id,
    date,
    address,
    total_blocks,
    worker_count,
    ROUND(blocks_per_worker, 1) AS blocks_per_worker,  -- Soddalashtirilgan
    price_per_block,
    ROUND(payment, 1) AS payment,  -- Soddalashtirilgan
    ROUND(SUM(payment) OVER (), 1) AS total_payment  -- To'g'ri hisoblash
FROM 
    worker_payments;


-- name: PaidMonthlyData :many
SELECT 
    worker_id,
    date,
    paid_price,
    created_at
FROM 
    paid_monthly
WHERE 
    worker_id = CAST($1 AS INTEGER)
    AND EXTRACT(YEAR FROM date) = CAST($2 AS INTEGER)
    AND EXTRACT(MONTH FROM date) = CAST($3 AS INTEGER)
    AND deleted_at = 0
ORDER BY 
    date;

-- name: AddPaidMonthly :exec
INSERT INTO paid_monthly
    (worker_id, date, paid_price)
VALUES 
    ($1, $2, $3);


-- name: UpdateWorkers :exec
UPDATE workers 
SET first_name = $2, last_name = $3, phone = $4
WHERE id = $1;

-- name: DeleteWorkers :exec
UPDATE workers
SET deleted_at = $2
WHERE id = $1;

-- name: GetDailyProductionWorkersNameById :many
SELECT 
    w.id,
    w.first_name,
    w.last_name
FROM 
    daily_production_workers dpw
JOIN 
    workers w ON dpw.worker_id = w.id
WHERE 
    dpw.daily_production_id = $1 
    AND dpw.deleted_at = 0
    AND w.deleted_at = 0;
