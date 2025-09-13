# ToDo Application

A modern, responsive to-do application built with Wails (Go + React) and PostgreSQL.

## Features

### Frontend
- **Modern UI**: Clean, responsive design with smooth animations
- **Sidebar Navigation**: Collapsible sidebar with menu toggle
- **Horizontal Scroll**: Lists are displayed in a horizontal scrollable container
- **Task Management**: 
  - Create, read, update, and delete tasks
  - Toggle task completion with checkboxes
  - Separate completed tasks section (collapsible)
  - Delete tasks with confirmation
- **List Management**:
  - Create new lists
  - View all lists in sidebar
  - Add tasks to specific lists

### Backend
- **PostgreSQL Database**: Robust data storage with proper relationships
- **String-based IDs**: Simple string identifiers for all entities
- **CRUD Operations**: Complete Create, Read, Update, Delete functionality
- **Cascading Deletes**: Deleting a list removes all its tasks
- **Automatic Timestamps**: Created and updated timestamps for all records

## Database Schema

### Lists
- `id` (TEXT PRIMARY KEY)
- `title` (VARCHAR)
- `created_at` (TIMESTAMP)
- `updated_at` (TIMESTAMP)

### Tasks
- `id` (TEXT PRIMARY KEY)
- `list_id` (TEXT FOREIGN KEY)
- `task_name` (VARCHAR)
- `description` (TEXT)
- `completed` (BOOLEAN)
- `created_at` (TIMESTAMP)
- `updated_at` (TIMESTAMP)

### SubTasks
- `id` (TEXT PRIMARY KEY)
- `task_id` (TEXT FOREIGN KEY)
- `subtask_name` (VARCHAR)
- `description` (TEXT)
- `completed` (BOOLEAN)
- `created_at` (TIMESTAMP)
- `updated_at` (TIMESTAMP)

## Getting Started

### Prerequisites
- Go 1.25.1 or later
- Node.js and npm
- PostgreSQL database
- Wails v2

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd to-do
   ```

2. **Set up the database**
   - Create a PostgreSQL database named `todo_db`
   - Update the connection string in `app.go` if needed
   - Default connection: `postgres://postgres:password@localhost:5432/todo_db?sslmode=disable`

3. **Install dependencies**
   ```bash
   # Backend dependencies
   go mod tidy
   
   # Frontend dependencies
   cd frontend
   npm install
   cd ..
   ```

4. **Build and run**
   ```bash
   # Development mode
   wails dev
   
   # Build for production
   wails build
   ```

## Usage

### Creating Lists
1. Click the menu button (☰) to open the sidebar
2. Click "+ Create New List" button
3. Enter a list name and click "Create"

### Adding Tasks
1. In any list, click "+ Add a task"
2. Enter task name and optional description
3. Click "Add" to create the task

### Managing Tasks
- **Complete a task**: Click the checkbox next to the task
- **Delete a task**: Click the "×" button next to the task
- **View completed tasks**: Click on "Completed (X)" to expand/collapse

### Navigation
- **Sidebar**: Contains list of all lists and create new list button
- **Main area**: Horizontal scrollable list of all your lists
- **Responsive**: Works on desktop and mobile devices

## API Functions

The application exposes the following functions to the frontend:

### Lists
- `CreateList(title string) (*List, error)`
- `GetList(id string) (*List, error)`
- `GetAllLists() ([]List, error)`
- `UpdateList(id string, title string) (*List, error)`
- `DeleteList(id string) error`

### Tasks
- `CreateTask(listID string, taskName, description string) (*Task, error)`
- `GetTask(id string) (*Task, error)`
- `GetTasksByListID(listID string) ([]Task, error)`
- `GetAllTasks() ([]Task, error)`
- `UpdateTask(id string, taskName, description string, completed bool) (*Task, error)`
- `ToggleTaskCompletion(id string) (*Task, error)`
- `DeleteTask(id string) error`

### SubTasks
- `CreateSubTask(taskID string, subTaskName, description string) (*SubTask, error)`
- `GetSubTask(id string) (*SubTask, error)`
- `GetSubTasksByTaskID(taskID string) ([]SubTask, error)`
- `GetAllSubTasks() ([]SubTask, error)`
- `UpdateSubTask(id string, subTaskName, description string, completed bool) (*SubTask, error)`
- `ToggleSubTaskCompletion(id string) (*SubTask, error)`
- `DeleteSubTask(id string) error`

## Architecture

### Backend Structure
```
internal/
├── models.go          # Data models (List, Task, SubTask)
├── server.go          # Server utilities
└── db/
    ├── db.go          # Database connection and initialization
    ├── lists.go       # List CRUD operations
    ├── tasks.go       # Task CRUD operations
    └── subtasks.go    # SubTask CRUD operations
```

### Frontend Structure
```
frontend/src/
├── App.jsx            # Main application component
├── App.css            # Application styles
└── main.jsx           # Application entry point
```

## Technologies Used

- **Backend**: Go, PostgreSQL, pgx/v5
- **Frontend**: React, CSS3
- **Framework**: Wails v2
- **Database**: PostgreSQL with TEXT primary keys

## License

This project is licensed under the MIT License.