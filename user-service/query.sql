-- name: InsertUser :exec
INSERT INTO users
    (login, password)
VALUES 
    ($1, $2);

-- name: Login :one
SELECT id, login, password FROM users
WHERE login = $1;