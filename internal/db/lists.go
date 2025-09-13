package db

import (
	"context"
	"fmt"

	server "github.com/HolySxn/To-Do/internal"
	"github.com/jackc/pgx/v5"
)

// CRUD operations for Lists
func (d *DB) CreateList(ctx context.Context, title string) (*server.List, error) {
	query := `INSERT INTO lists (title) VALUES ($1) RETURNING id, title, created_at, updated_at`

	var list server.List
	err := d.Pool.QueryRow(ctx, query, title).Scan(
		&list.ID,
		&list.Title,
		&list.CreatedAt,
		&list.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create list: %w", err)
	}

	return &list, nil
}

func (d *DB) GetList(ctx context.Context, id string) (*server.List, error) {
	query := `SELECT * FROM lists WHERE id = $1`

	var list server.List
	err := d.Pool.QueryRow(ctx, query, id).Scan(
		&list.ID,
		&list.Title,
		&list.CreatedAt,
		&list.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("list with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to get list: %w", err)
	}

	return &list, nil
}

func (d *DB) GetAllLists(ctx context.Context) ([]server.List, error) {
	query := `SELECT * FROM lists ORDER BY created_at DESC`

	rows, err := d.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get lists: %w", err)
	}
	defer rows.Close()

	var lists []server.List
	for rows.Next() {
		var list server.List
		err := rows.Scan(
			&list.ID,
			&list.Title,
			&list.CreatedAt,
			&list.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan list: %w", err)
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func (d *DB) UpdateList(ctx context.Context, id string, title string) (*server.List, error) {
	query := `UPDATE lists SET title = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2 RETURNING id, title, created_at, updated_at`

	var list server.List
	err := d.Pool.QueryRow(ctx, query, title, id).Scan(
		&list.ID,
		&list.Title,
		&list.CreatedAt,
		&list.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("list with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to update list: %w", err)
	}

	return &list, nil
}

func (d *DB) DeleteList(ctx context.Context, id string) error {
	query := `DELETE FROM lists WHERE id = $1`

	result, err := d.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete list: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("list with id %s not found", id)
	}

	return nil
}
