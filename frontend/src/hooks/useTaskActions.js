import { useState } from 'react'
import { 
  CreateList, 
  UpdateList, 
  DeleteList,
  CreateTask,
  UpdateTask,
  DeleteTask,
  ToggleTaskCompletion,
  CreateSubTask,
  UpdateSubTask,
  DeleteSubTask,
  ToggleSubTaskCompletion
} from '../../wailsjs/go/main/App'

export function useTaskActions() {
  const [sortStates, setSortStates] = useState({})

  // List actions
  const createList = async (setLists, setTaskLists) => {
    const listName = prompt('Enter list name:')
    if (!listName || !listName.trim()) return

    try {
      const newList = await CreateList(listName.trim())
      
      if (newList) {
        const listWithVisibility = {
          ...newList,
          visible: true,
          count: 0
        }
        
        const taskListWithVisibility = {
          ...newList,
          tasks: [],
          completedCollapsed: true,
          settingsOpen: false,
          visible: true
        }
        
        setLists(prev => [...prev, listWithVisibility])
        setTaskLists(prev => [...prev, taskListWithVisibility])
      }
    } catch (error) {
      alert('Failed to create list. Please try again.')
    }
  }

  const renameList = async (listId, setLists, setTaskLists) => {
    const list = setTaskLists.find(l => l.id === listId)
    if (!list) return

    const newName = prompt('Enter new list name:', list.title)
    if (!newName || !newName.trim()) return

    try {
      const updatedList = await UpdateList(listId, newName.trim())
      
      if (updatedList) {
        setLists(lists => lists.map(l => 
          l.id === listId 
            ? { ...l, title: updatedList.title }
            : l
        ))
        
        setTaskLists(lists => lists.map(l => 
          l.id === listId 
            ? { ...l, title: updatedList.title, settingsOpen: false }
            : l
        ))
      }
    } catch (error) {
      alert('Failed to rename list. Please try again.')
    }
  }

  const deleteList = async (listId, setLists, setTaskLists) => {
    const list = setTaskLists.find(l => l.id === listId)
    if (!list) return

    const confirmDelete = window.confirm(`Are you sure you want to delete the list "${list.title}"? This will also delete all tasks and subtasks in this list.`)
    if (!confirmDelete) return

    try {
      await DeleteList(listId)
      
      setLists(lists => lists.filter(l => l.id !== listId))
      setTaskLists(taskLists => taskLists.filter(l => l.id !== listId))
    } catch (error) {
      alert('Failed to delete list. Please try again.')
    }
  }

  // Task actions
  const addTask = async (listId, setTaskLists, setLists) => {
    const newTaskText = prompt('Enter new task:')
    if (newTaskText && newTaskText.trim()) {
      try {
        const newTask = await CreateTask(listId, newTaskText.trim())
        const taskWithUI = {
          ...newTask,
          text: newTask.task_name,
          subtasks: [],
          settingsOpen: false
        }
        
        setTaskLists(lists => lists.map(list => {
          if (list.id !== listId) return list
          
          const newTasks = [taskWithUI, ...list.tasks]
          
          // Apply current sort if exists
          const currentSort = sortStates[listId]
          if (currentSort && currentSort.type) {
            let sortedTasks = [...newTasks]
            
            if (currentSort.type === 'alphabetical') {
              sortedTasks.sort((a, b) => {
                const result = a.text.localeCompare(b.text)
                return currentSort.ascending ? result : -result
              })
            } else if (currentSort.type === 'date') {
              sortedTasks.sort((a, b) => {
                const result = new Date(a.created_at) - new Date(b.created_at)
                return currentSort.ascending ? result : -result
              })
            }
            
            return { ...list, tasks: sortedTasks }
          }
          
          return { ...list, tasks: newTasks }
        }))
        
        // Update list count
        setLists(lists => lists.map(list => 
          list.id === listId 
            ? { ...list, count: list.count + 1 }
            : list
        ))
      } catch (error) {
        alert('Failed to create task')
      }
    }
  }

  const toggleTask = async (listId, taskId, setTaskLists) => {
    try {
      const updatedTask = await ToggleTaskCompletion(taskId)
      
      setTaskLists(lists => lists.map(list => 
        list.id === listId 
          ? { 
              ...list, 
              tasks: list.tasks.map(task => 
                task.id === taskId 
                  ? { ...task, completed: updatedTask.completed }
                  : task
              )
            }
          : list
      ))
    } catch (error) {
      alert('Failed to toggle task')
    }
  }

  const deleteTask = async (listId, taskId, setTaskLists, setLists) => {
    if (confirm('Are you sure you want to delete this task?')) {
      try {
        await DeleteTask(taskId)
        
        setTaskLists(lists => lists.map(list => 
          list.id === listId 
            ? {
                ...list,
                tasks: list.tasks.filter(task => task.id !== taskId)
              }
            : list
        ))
        
        // Update list count
        setLists(lists => lists.map(list => 
          list.id === listId 
            ? { ...list, count: Math.max(0, list.count - 1) }
            : list
        ))
      } catch (error) {
        alert('Failed to delete task')
      }
    }
  }

  // Subtask actions
  const addSubtask = async (listId, taskId, setTaskLists) => {
    const subtaskText = prompt('Enter subtask:')
    if (subtaskText && subtaskText.trim()) {
      try {
        const newSubtask = await CreateSubTask(taskId, subtaskText.trim())
        const subtaskWithUI = {
          ...newSubtask,
          text: newSubtask.subtask_name,
          completed: newSubtask.completed
        }
        
        setTaskLists(lists => lists.map(list => 
          list.id === listId 
            ? {
                ...list,
                tasks: list.tasks.map(task => 
                  task.id === taskId 
                    ? {
                        ...task,
                        subtasks: [...task.subtasks, subtaskWithUI],
                        settingsOpen: false
                      }
                    : task
                )
              }
            : list
        ))
      } catch (error) {
        alert('Failed to create subtask')
      }
    }
  }

  const toggleSubtask = async (listId, taskId, subtaskId, setTaskLists) => {
    try {
      const updatedSubtask = await ToggleSubTaskCompletion(subtaskId)
      
      setTaskLists(lists => lists.map(list => 
        list.id === listId 
          ? {
              ...list,
              tasks: list.tasks.map(task => 
                task.id === taskId 
                  ? {
                      ...task,
                      subtasks: task.subtasks.map(subtask => 
                        subtask.id === subtaskId 
                          ? { ...subtask, completed: updatedSubtask.completed }
                          : subtask
                      )
                    }
                  : task
              )
            }
          : list
      ))
    } catch (error) {
      alert('Failed to toggle subtask')
    }
  }

  const deleteSubtask = async (listId, taskId, subtaskId, setTaskLists) => {
    try {
      await DeleteSubTask(subtaskId)
      
      setTaskLists(lists => lists.map(list => 
        list.id === listId 
          ? {
              ...list,
              tasks: list.tasks.map(task => 
                task.id === taskId 
                  ? {
                      ...task,
                      subtasks: task.subtasks.filter(subtask => subtask.id !== subtaskId)
                    }
                  : task
              )
            }
          : list
      ))
    } catch (error) {
      alert('Failed to delete subtask')
    }
  }

  return {
    sortStates,
    setSortStates,
    createList,
    renameList,
    deleteList,
    addTask,
    toggleTask,
    deleteTask,
    addSubtask,
    toggleSubtask,
    deleteSubtask
  }
}
