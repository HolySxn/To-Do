import React from 'react'
import SubTaskItem from './SubTaskItem'

function TaskItem({ 
  task, 
  listId, 
  taskLists, 
  onToggleTask, 
  onToggleTaskSettings, 
  onAddSubtask, 
  onDeleteTask, 
  onMoveTaskToList,
  onToggleSubtask,
  onDeleteSubtask
}) {
  return (
    <div className="task-container">
      <div className={`task-item ${task.completed ? 'completed' : ''}`}>
        <div className="task-radio" onClick={() => onToggleTask(listId, task.id)}>
          <div className={`radio-button ${task.completed ? 'checked' : ''}`}>
            {task.completed && <span className="radio-checkmark">✓</span>}
          </div>
        </div>
        <span className={`task-text ${task.completed ? 'completed-text' : ''}`}>{task.text}</span>
        <div className="task-settings">
          <button className="task-settings-button" onClick={(e) => onToggleTaskSettings(listId, task.id, e)}>
            <span className="task-settings-icon">⋯</span>
          </button>
          {task.settingsOpen && (
            <div className="task-settings-menu" onClick={(e) => e.stopPropagation()}>
              <div className="task-settings-item" onClick={() => onAddSubtask(listId, task.id)}>
                Add subtask
              </div>
              <div className="task-settings-item task-settings-danger" onClick={() => onDeleteTask(listId, task.id)}>
                Delete task
              </div>
              <div className="task-settings-divider"></div>
              <div className="task-settings-section">
                <div className="task-settings-section-header">
                  Move to list
                </div>
                <div className="task-settings-submenu">
                  {taskLists.filter(otherList => otherList.id !== listId).map(otherList => (
                    <div key={otherList.id} className="task-settings-subitem" onClick={() => onMoveTaskToList(listId, task.id, otherList.id)}>
                      {otherList.title}
                    </div>
                  ))}
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
      
      {/* Subtasks */}
      {task.subtasks && task.subtasks.length > 0 && (
        <div className="subtasks-list">
          {(task.subtasks || []).map((subtask) => (
            <SubTaskItem
              key={subtask.id}
              subtask={subtask}
              listId={listId}
              taskId={task.id}
              onToggleSubtask={onToggleSubtask}
              onDeleteSubtask={onDeleteSubtask}
            />
          ))}
        </div>
      )}
    </div>
  )
}

export default TaskItem
