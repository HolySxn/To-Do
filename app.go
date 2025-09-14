package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	server "github.com/HolySxn/To-Do/internal"
	"github.com/HolySxn/To-Do/internal/config"
	"github.com/HolySxn/To-Do/internal/db"
)

// App struct
type App struct {
	ctx    context.Context
	db     *db.DB
	config *config.Config
	logger *slog.Logger
}

// NewApp creates a new App application struct
func NewApp() *App {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize structured logger with config
	logLevel := slog.LevelInfo
	switch cfg.App.LogLevel {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	logger.Info("Initializing To-Do application", "config", cfg)

	// Get database connection string from config
	connStr := cfg.GetConnectionString()
	logger.Info("Connecting to database", "host", cfg.Database.Host, "port", cfg.Database.Port, "database", cfg.Database.Name)

	// Initialize database connection
	database, err := db.NewDB(connStr)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		log.Fatalf("Failed to connect to database: %v", err)
	}

	logger.Info("Database connection established successfully")

	return &App{
		db:     database,
		config: cfg,
		logger: logger,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.logger.Info("Application started successfully")
}

// shutdown is called when the app is shutting down
func (a *App) shutdown(ctx context.Context) {
	a.logger.Info("Application shutting down")
	if a.db != nil {
		a.db.Close()
		a.logger.Info("Database connection closed")
	}
}

// List CRUD operations
func (a *App) CreateList(title string) (*server.List, error) {
	a.logger.Info("Creating new list", "title", title)
	list, err := a.db.CreateList(a.ctx, title)
	if err != nil {
		a.logger.Error("Failed to create list", "title", title, "error", err)
		return nil, err
	}
	a.logger.Info("List created successfully", "list_id", list.ID, "title", list.Title)
	return list, nil
}

func (a *App) GetList(id string) (*server.List, error) {
	a.logger.Info("Getting list", "list_id", id)
	list, err := a.db.GetList(a.ctx, id)
	if err != nil {
		a.logger.Error("Failed to get list", "list_id", id, "error", err)
		return nil, err
	}
	a.logger.Info("List retrieved successfully", "list_id", id, "title", list.Title)
	return list, nil
}

func (a *App) GetAllLists() ([]server.List, error) {
	a.logger.Info("Getting all lists")
	lists, err := a.db.GetAllLists(a.ctx)
	if err != nil {
		a.logger.Error("Failed to get all lists", "error", err)
		return nil, err
	}
	a.logger.Info("All lists retrieved successfully", "count", len(lists))
	return lists, nil
}

func (a *App) UpdateList(id string, title string) (*server.List, error) {
	a.logger.Info("Updating list", "list_id", id, "new_title", title)
	list, err := a.db.UpdateList(a.ctx, id, title)
	if err != nil {
		a.logger.Error("Failed to update list", "list_id", id, "new_title", title, "error", err)
		return nil, err
	}
	a.logger.Info("List updated successfully", "list_id", id, "title", list.Title)
	return list, nil
}

func (a *App) DeleteList(id string) error {
	a.logger.Info("Deleting list", "list_id", id)
	err := a.db.DeleteList(a.ctx, id)
	if err != nil {
		a.logger.Error("Failed to delete list", "list_id", id, "error", err)
		return err
	}
	a.logger.Info("List deleted successfully", "list_id", id)
	return nil
}

// Task CRUD operations
func (a *App) CreateTask(listID string, taskName string) (*server.Task, error) {
	a.logger.Info("Creating new task", "list_id", listID, "task_name", taskName)
	task, err := a.db.CreateTask(a.ctx, listID, taskName)
	if err != nil {
		a.logger.Error("Failed to create task", "list_id", listID, "task_name", taskName, "error", err)
		return nil, err
	}
	a.logger.Info("Task created successfully", "task_id", task.ID, "list_id", listID, "task_name", task.TaskName)
	return task, nil
}

func (a *App) GetTask(id string) (*server.Task, error) {
	a.logger.Info("Getting task", "task_id", id)
	task, err := a.db.GetTask(a.ctx, id)
	if err != nil {
		a.logger.Error("Failed to get task", "task_id", id, "error", err)
		return nil, err
	}
	a.logger.Info("Task retrieved successfully", "task_id", id, "task_name", task.TaskName)
	return task, nil
}

func (a *App) GetTasksByListID(listID string) ([]server.Task, error) {
	a.logger.Info("Getting tasks by list ID", "list_id", listID)
	tasks, err := a.db.GetTasksByListID(a.ctx, listID)
	if err != nil {
		a.logger.Error("Failed to get tasks by list ID", "list_id", listID, "error", err)
		return nil, err
	}
	a.logger.Info("Tasks retrieved successfully", "list_id", listID, "count", len(tasks))
	return tasks, nil
}

func (a *App) GetAllTasks() ([]server.Task, error) {
	a.logger.Info("Getting all tasks")
	tasks, err := a.db.GetAllTasks(a.ctx)
	if err != nil {
		a.logger.Error("Failed to get all tasks", "error", err)
		return nil, err
	}
	a.logger.Info("All tasks retrieved successfully", "count", len(tasks))
	return tasks, nil
}

func (a *App) UpdateTask(id string, taskName string, completed bool) (*server.Task, error) {
	a.logger.Info("Updating task", "task_id", id, "task_name", taskName, "completed", completed)
	task, err := a.db.UpdateTask(a.ctx, id, taskName, completed)
	if err != nil {
		a.logger.Error("Failed to update task", "task_id", id, "task_name", taskName, "completed", completed, "error", err)
		return nil, err
	}
	a.logger.Info("Task updated successfully", "task_id", id, "task_name", task.TaskName, "completed", task.Completed)
	return task, nil
}

func (a *App) ToggleTaskCompletion(id string) (*server.Task, error) {
	a.logger.Info("Toggling task completion", "task_id", id)
	task, err := a.db.ToggleTaskCompletion(a.ctx, id)
	if err != nil {
		a.logger.Error("Failed to toggle task completion", "task_id", id, "error", err)
		return nil, err
	}
	a.logger.Info("Task completion toggled successfully", "task_id", id, "completed", task.Completed)
	return task, nil
}

func (a *App) DeleteTask(id string) error {
	a.logger.Info("Deleting task", "task_id", id)
	err := a.db.DeleteTask(a.ctx, id)
	if err != nil {
		a.logger.Error("Failed to delete task", "task_id", id, "error", err)
		return err
	}
	a.logger.Info("Task deleted successfully", "task_id", id)
	return nil
}

// SubTask CRUD operations
func (a *App) CreateSubTask(taskID string, subTaskName string) (*server.SubTask, error) {
	a.logger.Info("Creating new subtask", "task_id", taskID, "subtask_name", subTaskName)
	subtask, err := a.db.CreateSubTask(a.ctx, taskID, subTaskName)
	if err != nil {
		a.logger.Error("Failed to create subtask", "task_id", taskID, "subtask_name", subTaskName, "error", err)
		return nil, err
	}
	a.logger.Info("Subtask created successfully", "subtask_id", subtask.ID, "task_id", taskID, "subtask_name", subtask.SubTaskName)
	return subtask, nil
}

func (a *App) GetSubTask(id string) (*server.SubTask, error) {
	a.logger.Info("Getting subtask", "subtask_id", id)
	subtask, err := a.db.GetSubTask(a.ctx, id)
	if err != nil {
		a.logger.Error("Failed to get subtask", "subtask_id", id, "error", err)
		return nil, err
	}
	a.logger.Info("Subtask retrieved successfully", "subtask_id", id, "subtask_name", subtask.SubTaskName)
	return subtask, nil
}

func (a *App) GetSubTasksByTaskID(taskID string) ([]server.SubTask, error) {
	a.logger.Info("Getting subtasks by task ID", "task_id", taskID)
	subtasks, err := a.db.GetSubTasksByTaskID(a.ctx, taskID)
	if err != nil {
		a.logger.Error("Failed to get subtasks by task ID", "task_id", taskID, "error", err)
		return nil, err
	}
	a.logger.Info("Subtasks retrieved successfully", "task_id", taskID, "count", len(subtasks))
	return subtasks, nil
}

func (a *App) GetAllSubTasks() ([]server.SubTask, error) {
	a.logger.Info("Getting all subtasks")
	subtasks, err := a.db.GetAllSubTasks(a.ctx)
	if err != nil {
		a.logger.Error("Failed to get all subtasks", "error", err)
		return nil, err
	}
	a.logger.Info("All subtasks retrieved successfully", "count", len(subtasks))
	return subtasks, nil
}

func (a *App) UpdateSubTask(id string, subTaskName string, completed bool) (*server.SubTask, error) {
	a.logger.Info("Updating subtask", "subtask_id", id, "subtask_name", subTaskName, "completed", completed)
	subtask, err := a.db.UpdateSubTask(a.ctx, id, subTaskName, completed)
	if err != nil {
		a.logger.Error("Failed to update subtask", "subtask_id", id, "subtask_name", subTaskName, "completed", completed, "error", err)
		return nil, err
	}
	a.logger.Info("Subtask updated successfully", "subtask_id", id, "subtask_name", subtask.SubTaskName, "completed", subtask.Completed)
	return subtask, nil
}

func (a *App) ToggleSubTaskCompletion(id string) (*server.SubTask, error) {
	a.logger.Info("Toggling subtask completion", "subtask_id", id)
	subtask, err := a.db.ToggleSubTaskCompletion(a.ctx, id)
	if err != nil {
		a.logger.Error("Failed to toggle subtask completion", "subtask_id", id, "error", err)
		return nil, err
	}
	a.logger.Info("Subtask completion toggled successfully", "subtask_id", id, "completed", subtask.Completed)
	return subtask, nil
}

func (a *App) DeleteSubTask(id string) error {
	a.logger.Info("Deleting subtask", "subtask_id", id)
	err := a.db.DeleteSubTask(a.ctx, id)
	if err != nil {
		a.logger.Error("Failed to delete subtask", "subtask_id", id, "error", err)
		return err
	}
	a.logger.Info("Subtask deleted successfully", "subtask_id", id)
	return nil
}
