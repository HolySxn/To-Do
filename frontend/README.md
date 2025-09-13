# Simple ToDo Frontend

A clean, minimal React frontend for the ToDo Wails application.

## Structure

```
src/
├── App.jsx          # Main application component
├── App.css          # Main styles
└── main.jsx         # React entry point
```

## Features

- **Wails Integration**: Ready to use Wails backend functions
- **Loading States**: Handles Wails connection loading
- **Error Handling**: Shows errors and offline mode warnings
- **Responsive Design**: Works on desktop and mobile
- **Clean CSS**: Minimal, modern styling

## Available Wails Functions

The app has access to all backend functions through `window.go.main.App`:

- `CreateList(title)` - Create a new list
- `GetAllLists()` - Get all lists
- `CreateTask(listID, taskName, description)` - Create a new task
- `GetTasksByListID(listID)` - Get tasks for a list
- `ToggleTaskCompletion(taskID)` - Toggle task completion
- `DeleteTask(taskID)` - Delete a task
- And more...

## Usage Example

```javascript
// Get Wails functions
const getWailsFunctions = () => {
  if (!window.go || !window.go.main || !window.go.main.App) {
    throw new Error('Wails not ready')
  }
  return window.go.main.App
}

// Use a function
const { CreateList } = getWailsFunctions()
const newList = await CreateList('My New List')
```

## Development

1. Install dependencies: `npm install`
2. Start dev server: `npm run dev`
3. Build for production: `npm run build`

## Building with Wails

1. Run `wails build` from the project root
2. The built app will be in `build/bin/`

## Customization

This is a minimal starting point. You can:
- Add your own components
- Create custom hooks
- Implement your own UI/UX
- Add state management (Redux, Zustand, etc.)
- Use any React libraries you prefer

The Wails backend is fully functional and ready to use!
