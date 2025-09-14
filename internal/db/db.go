package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool   *pgxpool.Pool
	Logger *slog.Logger
}

func NewDB(connectionString string) (*DB, error) {
	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("Initializing database connection")

	// Create connection pool with configuration
	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		logger.Error("Failed to parse connection string", "error", err)
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}

	// Set connection pool settings
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 1

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		logger.Error("Failed to create connection pool", "error", err)
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	logger.Info("Testing database connection")
	if err := pool.Ping(context.Background()); err != nil {
		logger.Error("Failed to ping database", "error", err)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info("Database connection established successfully")

	db := &DB{
		Pool:   pool,
		Logger: logger,
	}

	// Initialize tables
	logger.Info("Initializing database tables")
	if err := db.initTables(); err != nil {
		logger.Error("Failed to initialize tables", "error", err)
		return nil, fmt.Errorf("failed to initialize tables: %w", err)
	}

	logger.Info("Database initialization completed successfully")
	return db, nil
}

func (d *DB) Close() {
	d.Logger.Info("Closing database connection")
	d.Pool.Close()
	d.Logger.Info("Database connection closed")
}

func (d *DB) initTables() error {
	queries := []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`,
		`CREATE TABLE IF NOT EXISTS lists (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			title VARCHAR(255) NOT NULL,
			position INTEGER DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS tasks (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			list_id UUID NOT NULL REFERENCES lists(id) ON DELETE CASCADE,
			task_name VARCHAR(255) NOT NULL,
			completed BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS subtasks (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			task_id UUID NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
			subtask_name VARCHAR(255) NOT NULL,
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

	return nil
}
