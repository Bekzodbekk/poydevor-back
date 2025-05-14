-- name: InsertUser :exec
INSERT INTO users
    (login, password)
VALUES 
    ($1, $2);

-- name: AuthLogin :one
SELECT * FROM users WHERE login = $1;