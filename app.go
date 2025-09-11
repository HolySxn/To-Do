package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/HolySxn/To-Do/internal"
	"github.com/HolySxn/To-Do/internal/db"
)

// App struct
type App struct {
	ctx context.Context
	db  *db.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	// Get database connection string from environment variable
	// Default to a local PostgreSQL instance
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:password@localhost:5432/todo_db?sslmode=disable"
	}

	// Initialize database connection
	database, err := db.NewDB(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &App{
		db: database,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// shutdown is called when the app is shutting down
func (a *App) shutdown(ctx context.Context) {
	if a.db != nil {
		a.db.Close()
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// List CRUD operations
func (a *App) CreateList(title string) (*internal.List, error) {
	return a.db.CreateList(a.ctx, title)
}

func (a *App) GetList(id string) (*internal.List, error) {
	return a.db.GetList(a.ctx, id)
}

func (a *App) GetAllLists() ([]internal.List, error) {
	return a.db.GetAllLists(a.ctx)
}

func (a *App) UpdateList(id string, title string) (*internal.List, error) {
	return a.db.UpdateList(a.ctx, id, title)
}

func (a *App) DeleteList(id string) error {
	return a.db.DeleteList(a.ctx, id)
}

// Task CRUD operations
func (a *App) CreateTask(listID string, taskName, description string) (*internal.Task, error) {
	return a.db.CreateTask(a.ctx, listID, taskName, description)
}

func (a *App) GetTask(id string) (*internal.Task, error) {
	return a.db.GetTask(a.ctx, id)
}

func (a *App) GetTasksByListID(listID string) ([]internal.Task, error) {
	return a.db.GetTasksByListID(a.ctx, listID)
}

func (a *App) GetAllTasks() ([]internal.Task, error) {
	return a.db.GetAllTasks(a.ctx)
}

func (a *App) UpdateTask(id string, taskName, description string, completed bool) (*internal.Task, error) {
	return a.db.UpdateTask(a.ctx, id, taskName, description, completed)
}

func (a *App) ToggleTaskCompletion(id string) (*internal.Task, error) {
	return a.db.ToggleTaskCompletion(a.ctx, id)
}

func (a *App) DeleteTask(id string) error {
	return a.db.DeleteTask(a.ctx, id)
}

// SubTask CRUD operations
func (a *App) CreateSubTask(taskID string, subTaskName, description string) (*internal.SubTask, error) {
	return a.db.CreateSubTask(a.ctx, taskID, subTaskName, description)
}

func (a *App) GetSubTask(id string) (*internal.SubTask, error) {
	return a.db.GetSubTask(a.ctx, id)
}

func (a *App) GetSubTasksByTaskID(taskID string) ([]internal.SubTask, error) {
	return a.db.GetSubTasksByTaskID(a.ctx, taskID)
}

func (a *App) GetAllSubTasks() ([]internal.SubTask, error) {
	return a.db.GetAllSubTasks(a.ctx)
}

func (a *App) UpdateSubTask(id string, subTaskName, description string, completed bool) (*internal.SubTask, error) {
	return a.db.UpdateSubTask(a.ctx, id, subTaskName, description, completed)
}

func (a *App) ToggleSubTaskCompletion(id string) (*internal.SubTask, error) {
	return a.db.ToggleSubTaskCompletion(a.ctx, id)
}

func (a *App) DeleteSubTask(id string) error {
	return a.db.DeleteSubTask(a.ctx, id)
}
