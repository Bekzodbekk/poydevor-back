package postgres

import (
	"database/sql"
	"fmt"
	"workers-service/internal/pkg/config"
	"workers-service/storage"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	target := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
	)

	db, err := sql.Open("postgres", target)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func NewQueries(db *sql.DB) *storage.Queries {
	return storage.New(db)
}
