package db

import (
	"context"
	"fmt"

	server "github.com/HolySxn/To-Do/internal"
	"github.com/jackc/pgx/v5"
)

// CRUD operations for SubTasks
func (d *DB) CreateSubTask(ctx context.Context, taskID string, subTaskName, description string) (*server.SubTask, error) {
	query := `INSERT INTO subtasks (task_id, subtask_name, description) VALUES ($1, $2, $3) RETURNING id, task_id, subtask_name, description, completed, created_at, updated_at`

	var subTask server.SubTask
	err := d.Pool.QueryRow(ctx, query, taskID, subTaskName, description).Scan(
		&subTask.ID,
		&subTask.TaskID,
		&subTask.SubTaskName,
		&subTask.Description,
		&subTask.Completed,
		&subTask.CreatedAt,
		&subTask.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create subtask: %w", err)
	}

	return &subTask, nil
}

func (d *DB) GetSubTask(ctx context.Context, id string) (*server.SubTask, error) {
	query := `SELECT * FROM subtasks WHERE id = $1`

	var subTask server.SubTask
	err := d.Pool.QueryRow(ctx, query, id).Scan(
		&subTask.ID,
		&subTask.TaskID,
		&subTask.SubTaskName,
		&subTask.Description,
		&subTask.Completed,
		&subTask.CreatedAt,
		&subTask.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("subtask with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to get subtask: %w", err)
	}

	return &subTask, nil
}

func (d *DB) GetSubTasksByTaskID(ctx context.Context, taskID string) ([]server.SubTask, error) {
	query := `SELECT * FROM subtasks WHERE task_id = $1 ORDER BY created_at DESC`

	rows, err := d.Pool.Query(ctx, query, taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subtasks: %w", err)
	}
	defer rows.Close()

	var subTasks []server.SubTask
	for rows.Next() {
		var subTask server.SubTask
		err := rows.Scan(
			&subTask.ID,
			&subTask.TaskID,
			&subTask.SubTaskName,
			&subTask.Description,
			&subTask.Completed,
			&subTask.CreatedAt,
			&subTask.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan subtask: %w", err)
		}
		subTasks = append(subTasks, subTask)
	}

	return subTasks, nil
}

func (d *DB) GetAllSubTasks(ctx context.Context) ([]server.SubTask, error) {
	query := `SELECT * FROM subtasks ORDER BY created_at DESC`

	rows, err := d.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get subtasks: %w", err)
	}
	defer rows.Close()

	var subTasks []server.SubTask
	for rows.Next() {
		var subTask server.SubTask
		err := rows.Scan(
			&subTask.ID,
			&subTask.TaskID,
			&subTask.SubTaskName,
			&subTask.Description,
			&subTask.Completed,
			&subTask.CreatedAt,
			&subTask.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan subtask: %w", err)
		}
		subTasks = append(subTasks, subTask)
	}

	return subTasks, nil
}

func (d *DB) UpdateSubTask(ctx context.Context, id, subTaskName, description string, completed bool) (*server.SubTask, error) {
	query := `UPDATE subtasks SET subtask_name = $1, description = $2, completed = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4 RETURNING id, task_id, subtask_name, description, completed, created_at, updated_at`

	var subTask server.SubTask
	err := d.Pool.QueryRow(ctx, query, subTaskName, description, completed, id).Scan(
		&subTask.ID,
		&subTask.TaskID,
		&subTask.SubTaskName,
		&subTask.Description,
		&subTask.Completed,
		&subTask.CreatedAt,
		&subTask.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("subtask with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to update subtask: %w", err)
	}

	return &subTask, nil
}

func (d *DB) ToggleSubTaskCompletion(ctx context.Context, id string) (*server.SubTask, error) {
	query := `UPDATE subtasks SET completed = NOT completed, updated_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING id, task_id, subtask_name, description, completed, created_at, updated_at`

	var subTask server.SubTask
	err := d.Pool.QueryRow(ctx, query, id).Scan(
		&subTask.ID,
		&subTask.TaskID,
		&subTask.SubTaskName,
		&subTask.Description,
		&subTask.Completed,
		&subTask.CreatedAt,
		&subTask.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("subtask with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to toggle subtask completion: %w", err)
	}

	return &subTask, nil
}

func (d *DB) DeleteSubTask(ctx context.Context, id string) error {
	query := `DELETE FROM subtasks WHERE id = $1`

	result, err := d.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete subtask: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("subtask with id %s not found", id)
	}

	return nil
}
