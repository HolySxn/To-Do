package db

import (
	"context"
	"fmt"

	server "github.com/HolySxn/To-Do/internal"
	"github.com/jackc/pgx/v5"
)

// CRUD operations for Lists
func (d *DB) CreateList(ctx context.Context, title string) (*server.List, error) {
	// Get the next position for the new list
	var maxPosition int
	query := `SELECT COALESCE(MAX(position), 0) FROM lists`
	err := d.Pool.QueryRow(ctx, query).Scan(&maxPosition)
	if err != nil {
		return nil, fmt.Errorf("failed to get max position: %w", err)
	}

	query = `INSERT INTO lists (title, position) VALUES ($1, $2) RETURNING id, title, position, created_at, updated_at`

	var list server.List
	err = d.Pool.QueryRow(ctx, query, title, maxPosition+1).Scan(
		&list.ID,
		&list.Title,
		&list.Position,
		&list.CreatedAt,
		&list.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create list: %w", err)
	}

	return &list, nil
}

func (d *DB) GetList(ctx context.Context, id string) (*server.List, error) {
	query := `SELECT id, title, position, created_at, updated_at FROM lists WHERE id = $1`

	var list server.List
	err := d.Pool.QueryRow(ctx, query, id).Scan(
		&list.ID,
		&list.Title,
		&list.Position,
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
	query := `SELECT id, title, position, created_at, updated_at FROM lists ORDER BY position ASC`

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
			&list.Position,
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
	query := `UPDATE lists SET title = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2 RETURNING id, title, position, created_at, updated_at`

	var list server.List
	err := d.Pool.QueryRow(ctx, query, title, id).Scan(
		&list.ID,
		&list.Title,
		&list.Position,
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

// ReorderLists updates the positions of lists based on the provided order
func (d *DB) ReorderLists(ctx context.Context, listIDs []string) error {
	if len(listIDs) == 0 {
		return nil
	}

	// Start a transaction to ensure atomicity
	tx, err := d.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Update positions for each list
	for i, listID := range listIDs {
		query := `UPDATE lists SET position = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`
		_, err := tx.Exec(ctx, query, i+1, listID)
		if err != nil {
			return fmt.Errorf("failed to update list position: %w", err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
