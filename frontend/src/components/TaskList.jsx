import React from 'react'
import TaskItem from './TaskItem'

function TaskList({ 
  list, 
  taskLists, 
  sortStates, 
  onToggleSettings, 
  onSortTasksBy, 
  onRenameList, 
  onDeleteAllCompleted, 
  onDeleteList, 
  onAddTask, 
  onToggleTask, 
  onToggleCompletedTasks, 
  onToggleTaskSettings, 
  onAddSubtask, 
  onDeleteTask, 
  onMoveTaskToList, 
  onToggleSubtask, 
  onDeleteSubtask 
}) {
  const activeTasks = list.tasks.filter(task => !task.completed)
  const completedTasks = list.tasks.filter(task => task.completed)

  return (
    <div className="task-list-container">
      <div className="task-list-header">
        <h2>{list.title}</h2>
        <div className="header-actions">
          <button className="settings-button" onClick={(e) => onToggleSettings(list.id, e)}>
            <span className="three-dots">⋯</span>
          </button>
          {list.settingsOpen && (
            <div className="settings-menu" onClick={(e) => e.stopPropagation()}>
              <div className="settings-section">
                <div className="settings-section-header">
                  <span>Sort by</span>
                </div>
                <div className="settings-submenu">
                  <div className="settings-subitem" onClick={() => onSortTasksBy(list.id, 'alphabetical')}>
                    <span>Alphabetically</span>
                    <span className="sort-arrow">
                      {sortStates[list.id]?.type === 'alphabetical' 
                        ? (sortStates[list.id]?.ascending ? '↑' : '↓')
                        : ''
                      }
                    </span>
                  </div>
                  <div className="settings-subitem" onClick={() => onSortTasksBy(list.id, 'date')}>
                    <span>Date</span>
                    <span className="sort-arrow">
                      {sortStates[list.id]?.type === 'date' 
                        ? (sortStates[list.id]?.ascending ? '↑' : '↓')
                        : ''
                      }
                    </span>
                  </div>
                </div>
              </div>
              <div className="settings-divider"></div>
              <div className="settings-item" onClick={() => onRenameList(list.id)}>
                Rename
              </div>
              <div className="settings-item" onClick={() => onDeleteAllCompleted(list.id)}>
                Delete all completed
              </div>
              <div className="settings-divider"></div>
              <div className="settings-item settings-danger" onClick={() => onDeleteList(list.id)}>
                Delete list
              </div>
            </div>
          )}
        </div>
      </div>
      
      {/* Add Task Button */}
      <div className="add-task-button" onClick={() => onAddTask(list.id)}>
        <div className="add-task-radio">
          <div className="add-task-circle">+</div>
        </div>
        <span className="add-task-text">Add a task</span>
      </div>
      
      <div className="task-list">
        {/* Active Tasks */}
        {activeTasks.map((task) => (
          <TaskItem
            key={task.id}
            task={task}
            listId={list.id}
            taskLists={taskLists}
            onToggleTask={onToggleTask}
            onToggleTaskSettings={onToggleTaskSettings}
            onAddSubtask={onAddSubtask}
            onDeleteTask={onDeleteTask}
            onMoveTaskToList={onMoveTaskToList}
            onToggleSubtask={onToggleSubtask}
            onDeleteSubtask={onDeleteSubtask}
          />
        ))}
        
        {/* Completed Tasks Section */}
        {completedTasks.length > 0 && (
          <div className="completed-section">
            <div className="completed-header" onClick={() => onToggleCompletedTasks(list.id)}>
              <span className="completed-title">Completed Tasks</span>
              <span className={`completed-arrow ${list.completedCollapsed ? 'collapsed' : 'expanded'}`}>▼</span>
            </div>
            
            <div className={`completed-container ${list.completedCollapsed ? 'completed-collapsed' : 'completed-expanded'}`}>
              {completedTasks.map((task) => (
                <TaskItem
                  key={task.id}
                  task={task}
                  listId={list.id}
                  taskLists={taskLists}
                  onToggleTask={onToggleTask}
                  onToggleTaskSettings={onToggleTaskSettings}
                  onAddSubtask={onAddSubtask}
                  onDeleteTask={onDeleteTask}
                  onMoveTaskToList={onMoveTaskToList}
                  onToggleSubtask={onToggleSubtask}
                  onDeleteSubtask={onDeleteSubtask}
                />
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  )
}

export default TaskList
