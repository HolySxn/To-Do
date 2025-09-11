package db

import (
	"context"
	"fmt"

	"github.com/HolySxn/To-Do/internal"
	"github.com/jackc/pgx/v5"
)

// CRUD operations for Tasks
func (d *DB) CreateTask(ctx context.Context, listID, taskName, description string) (*internal.Task, error) {
	query := `INSERT INTO tasks (list_id, task_name, description) VALUES ($1, $2, $3) RETURNING id, list_id, task_name, description, completed, created_at, updated_at`

	var task internal.Task
	err := d.Pool.QueryRow(ctx, query, listID, taskName, description).Scan(
		&task.ID,
		&task.ListID,
		&task.TaskName,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return &task, nil
}

func (d *DB) GetTask(ctx context.Context, id string) (*internal.Task, error) {
	query := `SELECT * FROM tasks WHERE id = $1`

	var task internal.Task
	err := d.Pool.QueryRow(ctx, query, id).Scan(
		&task.ID,
		&task.ListID,
		&task.TaskName,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("task with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return &task, nil
}

func (d *DB) GetTasksByListID(ctx context.Context, listID string) ([]internal.Task, error) {
	query := `SELECT * FROM tasks WHERE list_id = $1 ORDER BY created_at DESC`

	rows, err := d.Pool.Query(ctx, query, listID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}
	defer rows.Close()

	var tasks []internal.Task
	for rows.Next() {
		var task internal.Task
		err := rows.Scan(
			&task.ID,
			&task.ListID,
			&task.TaskName,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (d *DB) GetAllTasks(ctx context.Context) ([]internal.Task, error) {
	query := `SELECT * FROM tasks ORDER BY created_at DESC`

	rows, err := d.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}
	defer rows.Close()

	var tasks []internal.Task
	for rows.Next() {
		var task internal.Task
		err := rows.Scan(
			&task.ID,
			&task.ListID,
			&task.TaskName,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (d *DB) UpdateTask(ctx context.Context, id, taskName, description string, completed bool) (*internal.Task, error) {
	query := `UPDATE tasks SET task_name = $1, description = $2, completed = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4 RETURNING id, list_id, task_name, description, completed, created_at, updated_at`

	var task internal.Task
	err := d.Pool.QueryRow(ctx, query, taskName, description, completed, id).Scan(
		&task.ID,
		&task.ListID,
		&task.TaskName,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("task with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return &task, nil
}

func (d *DB) ToggleTaskCompletion(ctx context.Context, id string) (*internal.Task, error) {
	query := `UPDATE tasks SET completed = NOT completed, updated_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING id, list_id, task_name, description, completed, created_at, updated_at`

	var task internal.Task
	err := d.Pool.QueryRow(ctx, query, id).Scan(
		&task.ID,
		&task.ListID,
		&task.TaskName,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("task with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to toggle task completion: %w", err)
	}

	return &task, nil
}

func (d *DB) DeleteTask(ctx context.Context, id string) error {
	query := `DELETE FROM tasks WHERE id = $1`

	result, err := d.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("task with id %s not found", id)
	}

	return nil
}
