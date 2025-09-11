package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDB(connectionString string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db := &DB{Pool: pool}

	// Initialize tables
	if err := db.initTables(); err != nil {
		return nil, fmt.Errorf("failed to initialize tables: %w", err)
	}

	return db, nil
}

func (d *DB) Close() {
	d.Pool.Close()
}

func (d *DB) initTables() error {
	queries := []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`,
		`CREATE TABLE IF NOT EXISTS lists (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			title VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS tasks (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			list_id UUID NOT NULL REFERENCES lists(id) ON DELETE CASCADE,
			task_name VARCHAR(255) NOT NULL,
			description TEXT,
			completed BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS subtasks (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			task_id UUID NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
			subtask_name VARCHAR(255) NOT NULL,
			description TEXT,
			completed BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
	}

	for _, query := range queries {
		if _, err := d.Pool.Exec(context.Background(), query); err != nil {
			return fmt.Errorf("failed to execute query: %w", err)
		}
	}

	log.Println("Database tables initialized successfully")
	return nil
}
