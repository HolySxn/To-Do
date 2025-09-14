import React from 'react'

function SubTaskItem({ 
  subtask, 
  listId, 
  taskId, 
  onToggleSubtask, 
  onDeleteSubtask 
}) {
  return (
    <div className="subtask-item">
      <div className="subtask-radio" onClick={() => onToggleSubtask(listId, taskId, subtask.id)}>
        <div className={`subtask-button ${subtask.completed ? 'checked' : ''}`}>
          {subtask.completed && <span className="subtask-checkmark">✓</span>}
        </div>
      </div>
      <span className={`subtask-text ${subtask.completed ? 'completed' : ''}`}>{subtask.text}</span>
      <button className="subtask-delete" onClick={() => onDeleteSubtask(listId, taskId, subtask.id)}>
        ×
      </button>
    </div>
  )
}

export default SubTaskItem
