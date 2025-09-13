import {useState} from 'react'
import './App.css'

function App() {
  const [activeSection, setActiveSection] = useState('dashboard')
  const [sidebarOpen, setSidebarOpen] = useState(false)
  const [listsCollapsed, setListsCollapsed] = useState(true)
  const [lists, setLists] = useState([
    { id: 1, name: 'Work Tasks', count: 5, visible: true },
    { id: 2, name: 'Personal', count: 3, visible: true },
    { id: 3, name: 'Shopping', count: 8, visible: false },
    { id: 4, name: 'Learning', count: 2, visible: true },
    { id: 5, name: 'Home Projects', count: 4, visible: true }
  ])

  const [taskLists, setTaskLists] = useState([
    {
      id: 1,
      name: 'Work Tasks',
      tasks: [
        { id: 1, text: 'Complete project proposal', completed: false },
        { id: 2, text: 'Review team feedback', completed: false },
        { id: 3, text: 'Schedule meeting with client', completed: false },
        { id: 4, text: 'Update documentation', completed: true },
        { id: 5, text: 'Prepare presentation slides', completed: true }
      ],
      completedCollapsed: true,
      settingsOpen: false
    },
    {
      id: 2,
      name: 'Personal Tasks',
      tasks: [
        { id: 6, text: 'Buy groceries', completed: false },
        { id: 7, text: 'Call dentist', completed: true },
        { id: 8, text: 'Plan weekend trip', completed: false }
      ],
      completedCollapsed: true,
      settingsOpen: false
    },
    {
      id: 3,
      name: 'Shopping List',
      tasks: [
        { id: 9, text: 'Buy milk', completed: false },
        { id: 10, text: 'Buy bread', completed: false },
        { id: 11, text: 'Buy eggs', completed: true },
        { id: 12, text: 'Buy vegetables', completed: false },
        { id: 13, text: 'Buy meat', completed: false }
      ],
      completedCollapsed: true,
      settingsOpen: false
    },
    {
      id: 4,
      name: 'Learning Goals',
      tasks: [
        { id: 14, text: 'Read React documentation', completed: false },
        { id: 15, text: 'Practice TypeScript', completed: false },
        { id: 16, text: 'Build a side project', completed: false }
      ],
      completedCollapsed: true,
      settingsOpen: false
    }
  ])

  const toggleSidebar = () => {
    setSidebarOpen(!sidebarOpen)
  }

  const closeSidebar = () => {
    setSidebarOpen(false)
  }

  const toggleLists = () => {
    setListsCollapsed(!listsCollapsed)
  }

  const toggleListVisibility = (listId) => {
    setLists(lists.map(list => 
      list.id === listId ? { ...list, visible: !list.visible } : list
    ))
  }


  const toggleSettings = (listId, e) => {
    e.stopPropagation()
    setTaskLists(lists => lists.map(list => 
      list.id === listId 
        ? { ...list, settingsOpen: !list.settingsOpen }
        : { ...list, settingsOpen: false } // Close other lists' settings
    ))
  }

  const closeAllSettings = () => {
    setTaskLists(lists => lists.map(list => ({ ...list, settingsOpen: false })))
  }

  const sortTasksBy = (listId, type) => {
    setTaskLists(lists => lists.map(list => {
      if (list.id !== listId) return list
      
      let sortedTasks = [...list.tasks]
      if (type === 'alphabetical') {
        sortedTasks.sort((a, b) => a.text.localeCompare(b.text))
      } else if (type === 'date') {
        // For now, just maintain current order (could be enhanced with actual dates)
        sortedTasks = list.tasks
      } else if (type === 'priority') {
        // Place completed tasks at the end
        sortedTasks.sort((a, b) => {
          if (a.completed === b.completed) return 0
          return a.completed ? 1 : -1
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

  const renameList = (listId) => {
    const list = taskLists.find(l => l.id === listId)
    const newName = prompt('Enter new list name:', list.name)
    if (newName && newName.trim()) {
      setTaskLists(lists => lists.map(list => 
        list.id === listId 
          ? { ...list, name: newName, settingsOpen: false }
          : list
      ))
    }
  }

  const toggleTask = (listId, taskId) => {
    setTaskLists(lists => lists.map(list => 
      list.id === listId 
        ? { 
            ...list, 
            tasks: list.tasks.map(task => 
              task.id === taskId ? { ...task, completed: !task.completed } : task
            )
          }
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

  return (
    <div className="app">
      
      <div className="header">
        <div className="header-bar-button" onClick={toggleSidebar}>
          <div className="header-bar"></div>
          <div className="header-bar"></div>
          <div className="header-bar"></div> 
        </div>
        <img className="header-icon" src="/src/assets/icons/checked.png" alt="Checked Icon"/>
        <h1 className="header-title">ToDo</h1>
      </div>

      {/* Sidebar */}
      <div className={`sidebar ${sidebarOpen ? 'sidebar-open' : ''}`}>
        <button className="sidebar-create-list">+ Create</button>
        
        {/* Collapsible Lists Section */}
        <div className="lists-section">
          <div className="lists-header" onClick={toggleLists}>
            <span className="lists-title">Lists</span>
            <span className={`collapse-arrow ${listsCollapsed ? 'collapsed' : 'expanded'}`}>▼</span>
          </div>
          
          <div className={`lists-container ${listsCollapsed ? 'lists-collapsed' : 'lists-expanded'}`}>
            {lists.map((list) => (
              <div key={list.id} className="list-item" onClick={() => toggleListVisibility(list.id)}>
                <div className="checkbox-container">
                  <div className={`checkbox ${list.visible ? 'checked' : ''}`}>
                    {list.visible && <span className="checkmark">✓</span>}
                  </div>
                  <span className="list-name">{list.name}</span>
                </div>
                <span className="list-count">{list.count}</span>
              </div>
            ))}
          </div>
        </div>
      </div>

      {/* Main Content */}
      <div className={`main-content ${sidebarOpen ? 'main-content-shifted' : ''}`} onClick={closeAllSettings}>
        <div className="content-area">
          {taskLists.map((list) => (
            <div key={list.id} className="task-list-container">
              <div className="task-list-header">
                <h2>{list.name}</h2>
                <div className="header-actions">
                  <button className="settings-button" onClick={(e) => toggleSettings(list.id, e)}>
                    <span className="three-dots">⋯</span>
                  </button>
                  {list.settingsOpen && (
                    <div className="settings-menu" onClick={(e) => e.stopPropagation()}>
                      <div className="settings-section">
                        <div className="settings-section-header">
                          <span>Sort by</span>
                        </div>
                        <div className="settings-submenu">
                          <div className="settings-subitem" onClick={() => sortTasksBy(list.id, 'alphabetical')}>
                            Alphabetically
                          </div>
                          <div className="settings-subitem" onClick={() => sortTasksBy(list.id, 'date')}>
                            Date
                          </div>
                          <div className="settings-subitem" onClick={() => sortTasksBy(list.id, 'priority')}>
                            Priority
                          </div>
                        </div>
                      </div>
                      <div className="settings-divider"></div>
                      <div className="settings-item" onClick={() => renameList(list.id)}>
                        Rename
                      </div>
                      <div className="settings-item" onClick={() => deleteAllCompleted(list.id)}>
                        Delete all completed
                      </div>
                    </div>
                  )}
                </div>
              </div>
              
              <div className="task-list">
                {/* Active Tasks */}
                {list.tasks.filter(task => !task.completed).map((task) => (
                  <div key={task.id} className="task-item">
                    <div className="task-radio" onClick={() => toggleTask(list.id, task.id)}>
                      <div className={`radio-button ${task.completed ? 'checked' : ''}`}>
                        {task.completed && <span className="radio-checkmark">✓</span>}
                      </div>
                    </div>
                    <span className="task-text">{task.text}</span>
                  </div>
                ))}
                
                {/* Completed Tasks Section */}
                {list.tasks.filter(task => task.completed).length > 0 && (
                  <div className="completed-section">
                    <div className="completed-header" onClick={() => toggleCompletedTasks(list.id)}>
                      <span className="completed-title">Completed Tasks</span>
                      <span className={`completed-arrow ${list.completedCollapsed ? 'collapsed' : 'expanded'}`}>▼</span>
                    </div>
                    
                    <div className={`completed-container ${list.completedCollapsed ? 'completed-collapsed' : 'completed-expanded'}`}>
                      {list.tasks.filter(task => task.completed).map((task) => (
                        <div key={task.id} className="task-item completed">
                          <div className="task-radio" onClick={() => toggleTask(list.id, task.id)}>
                            <div className={`radio-button checked`}>
                              <span className="radio-checkmark">✓</span>
                            </div>
                          </div>
                          <span className="task-text completed-text">{task.text}</span>
                        </div>
                      ))}
                    </div>
                  </div>
                )}
              </div>
            </div>
          ))}
        </div>
      </div>


    </div>
  )
}

export default App