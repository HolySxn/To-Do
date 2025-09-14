import React, { useState } from 'react'
import './App.css'
import Header from './components/Header'
import Sidebar from './components/Sidebar'
import TaskList from './components/TaskList'
import LoadingScreen from './components/LoadingScreen'
import { useTodoData } from './hooks/useTodoData'
import { useTaskActions } from './hooks/useTaskActions'

function App() {
  const [sidebarOpen, setSidebarOpen] = useState(false)
  const [listsCollapsed, setListsCollapsed] = useState(true)
  
  const { lists, setLists, taskLists, setTaskLists, loading } = useTodoData()
  const { 
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
  } = useTaskActions()

  // UI state handlers
  const toggleSidebar = () => setSidebarOpen(!sidebarOpen)
  const toggleLists = () => setListsCollapsed(!listsCollapsed)

  const toggleListVisibility = (listId) => {
    setLists(lists.map(list => 
      list.id === listId ? { ...list, visible: !list.visible } : list
    ))
    
    setTaskLists(taskLists.map(list => 
      list.id === listId ? { ...list, visible: !list.visible } : list
    ))
  }

  const toggleSettings = (listId, e) => {
    e.stopPropagation()
    setTaskLists(lists => lists.map(list => 
      list.id === listId 
        ? { ...list, settingsOpen: !list.settingsOpen }
        : { ...list, settingsOpen: false }
    ))
  }

  const closeAllSettings = () => {
    setTaskLists(lists => lists.map(list => ({ ...list, settingsOpen: false })))
  }

  const sortTasksBy = (listId, type) => {
    setTaskLists(lists => lists.map(list => {
      if (list.id !== listId) return list
      
      const currentSort = sortStates[listId] || { type: null, ascending: true }
      
      let ascending = true
      if (currentSort.type === type) {
        ascending = !currentSort.ascending
      } else {
        ascending = type === 'date' ? false : true
      }
      
      setSortStates(prev => ({
        ...prev,
        [listId]: { type, ascending }
      }))
      
      let sortedTasks = [...list.tasks]
      
      if (type === 'alphabetical') {
        sortedTasks.sort((a, b) => {
          const result = a.text.localeCompare(b.text)
          return ascending ? result : -result
        })
      } else if (type === 'date') {
        sortedTasks.sort((a, b) => {
          const result = new Date(a.created_at) - new Date(b.created_at)
          return ascending ? result : -result
        })
      }
      
      return { ...list, tasks: sortedTasks, settingsOpen: false }
    }))
  }

  const deleteAllCompleted = (listId) => {
    setTaskLists(lists => lists.map(list => 
      list.id === listId 
        ? { ...list, tasks: list.tasks.filter(task => !task.completed), settingsOpen: false }
        : list
    ))
  }

  const toggleCompletedTasks = (listId) => {
    setTaskLists(lists => lists.map(list => 
      list.id === listId 
        ? { ...list, completedCollapsed: !list.completedCollapsed }
        : list
    ))
  }

  const toggleTaskSettings = (listId, taskId, e) => {
    e.stopPropagation()
    setTaskLists(lists => lists.map(list => 
      list.id === listId 
        ? {
            ...list,
            tasks: list.tasks.map(task => 
              task.id === taskId 
                ? { ...task, settingsOpen: !task.settingsOpen }
                : { ...task, settingsOpen: false }
            )
          }
        : list
    ))
  }

  const closeAllTaskSettings = () => {
    setTaskLists(lists => lists.map(list => ({
      ...list,
      tasks: list.tasks.map(task => ({ ...task, settingsOpen: false }))
    })))
  }

  const moveTaskToList = (fromListId, taskId, toListId) => {
    if (fromListId === toListId) return

    const sourceList = taskLists.find(list => list.id === fromListId)
    const taskToMove = sourceList.tasks.find(task => task.id === taskId)

    setTaskLists(lists => lists.map(list => {
      if (list.id === fromListId) {
        return {
          ...list,
          tasks: list.tasks.filter(task => task.id !== taskId)
        }
      } else if (list.id === toListId) {
        return {
          ...list,
          tasks: [
            ...list.tasks,
            { ...taskToMove, settingsOpen: false }
          ]
        }
      }
      return list
    }))
  }

  // Wrapper functions for actions
  const handleCreateList = () => createList(setLists, setTaskLists)
  const handleRenameList = (listId) => renameList(listId, setLists, setTaskLists)
  const handleDeleteList = (listId) => deleteList(listId, setLists, setTaskLists)
  const handleAddTask = (listId) => addTask(listId, setTaskLists, setLists)
  const handleToggleTask = (listId, taskId) => toggleTask(listId, taskId, setTaskLists)
  const handleDeleteTask = (listId, taskId) => deleteTask(listId, taskId, setTaskLists, setLists)
  const handleAddSubtask = (listId, taskId) => addSubtask(listId, taskId, setTaskLists)
  const handleToggleSubtask = (listId, taskId, subtaskId) => toggleSubtask(listId, taskId, subtaskId, setTaskLists)
  const handleDeleteSubtask = (listId, taskId, subtaskId) => deleteSubtask(listId, taskId, subtaskId, setTaskLists)

  if (loading) {
    return <LoadingScreen onToggleSidebar={toggleSidebar} />
  }

  return (
    <div className="app">
      <Header onToggleSidebar={toggleSidebar} />

      <Sidebar
        sidebarOpen={sidebarOpen}
        lists={lists}
        listsCollapsed={listsCollapsed}
        onToggleLists={toggleLists}
        onToggleListVisibility={toggleListVisibility}
        onCreateList={handleCreateList}
      />

      <div className={`main-content ${sidebarOpen ? 'main-content-shifted' : ''}`} onClick={() => { closeAllSettings(); closeAllTaskSettings(); }}>
        <div className="content-area">
          {taskLists.filter(list => list.visible).map((list) => (
            <TaskList
              key={list.id}
              list={list}
              taskLists={taskLists}
              sortStates={sortStates}
              onToggleSettings={toggleSettings}
              onSortTasksBy={sortTasksBy}
              onRenameList={handleRenameList}
              onDeleteAllCompleted={deleteAllCompleted}
              onDeleteList={handleDeleteList}
              onAddTask={handleAddTask}
              onToggleTask={handleToggleTask}
              onToggleCompletedTasks={toggleCompletedTasks}
              onToggleTaskSettings={toggleTaskSettings}
              onAddSubtask={handleAddSubtask}
              onDeleteTask={handleDeleteTask}
              onMoveTaskToList={moveTaskToList}
              onToggleSubtask={handleToggleSubtask}
              onDeleteSubtask={handleDeleteSubtask}
            />
          ))}
        </div>
      </div>
    </div>
  )
}

export default App