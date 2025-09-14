import { useState, useEffect } from 'react'
import { 
  GetAllLists, 
  GetTasksByListID,
  GetSubTasksByTaskID
} from '../../wailsjs/go/main/App'

export function useTodoData() {
  const [lists, setLists] = useState([])
  const [taskLists, setTaskLists] = useState([])
  const [loading, setLoading] = useState(true)

  const loadData = async () => {
    try {
      setLoading(true)
      
      // Load lists
      const listsData = await GetAllLists()
      
      // Check if listsData is null or undefined
      if (!listsData) {
        setLoading(false)
        return
      }
      
      const listsWithVisibility = listsData.map(list => ({
        ...list,
        visible: true,
        count: 0
      }))
      setLists(listsWithVisibility)

      // Load tasks for each list
      const taskListsData = []
      for (const list of listsData) {
        try {
          const tasks = await GetTasksByListID(list.id)
          
          // Check if tasks is null or undefined
          if (!tasks) {
            taskListsData.push({
              ...list,
              tasks: [],
              completedCollapsed: true,
              settingsOpen: false,
              visible: true
            })
            continue
          }
          
          const tasksWithSubtasks = await Promise.all(
            tasks.map(async (task) => {
              try {
                const subtasks = await GetSubTasksByTaskID(task.id)
                
                // Check if subtasks is null or undefined
                if (!subtasks) {
                  return {
                    ...task,
                    text: task.task_name || 'Untitled Task',
                    subtasks: [],
                    settingsOpen: false
                  }
                }
                
                return {
                  ...task,
                  text: task.task_name || 'Untitled Task', // Map task_name to text for frontend compatibility
                  subtasks: subtasks.map(subtask => ({
                    ...subtask,
                    text: subtask.subtask_name || 'Untitled Subtask', // Map subtask_name to text
                    completed: subtask.completed || false
                  })),
                  settingsOpen: false
                }
              } catch (subtaskError) {
                return {
                  ...task,
                  text: task.task_name || 'Untitled Task',
                  subtasks: [],
                  settingsOpen: false
                }
              }
            })
          )
          
          taskListsData.push({
            ...list,
            tasks: tasksWithSubtasks,
            completedCollapsed: true,
            settingsOpen: false,
            visible: true
          })
        } catch (taskError) {
          // Add list with empty tasks if task loading fails
          taskListsData.push({
            ...list,
            tasks: [],
            completedCollapsed: true,
            settingsOpen: false,
            visible: true
          })
        }
      }
      
      setTaskLists(taskListsData)
      
      // Update list counts
      const updatedLists = listsWithVisibility.map(list => {
        const taskList = taskListsData.find(tl => tl.id === list.id)
        return {
          ...list,
          count: taskList ? taskList.tasks.length : 0
        }
      })
      setLists(updatedLists)
      
    } catch (error) {
      // Handle error silently or show user-friendly message
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    loadData()
  }, [])

  return {
    lists,
    setLists,
    taskLists,
    setTaskLists,
    loading,
    loadData
  }
}
