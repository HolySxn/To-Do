import React from 'react'

function Sidebar({ 
  sidebarOpen, 
  lists, 
  listsCollapsed, 
  onToggleLists, 
  onToggleListVisibility, 
  onCreateList 
}) {
  return (
    <div className={`sidebar ${sidebarOpen ? 'sidebar-open' : ''}`}>
      <button className="sidebar-create-list" onClick={onCreateList}>+ Create</button>
      
      {/* Collapsible Lists Section */}
      <div className="lists-section">
        <div className="lists-header" onClick={onToggleLists}>
          <span className="lists-title">Lists</span>
          <span className={`collapse-arrow ${listsCollapsed ? 'collapsed' : 'expanded'}`}>▼</span>
        </div>
        
        <div className={`lists-container ${listsCollapsed ? 'lists-collapsed' : 'lists-expanded'}`}>
          {lists.map((list) => (
            <div key={list.id} className="list-item" onClick={() => onToggleListVisibility(list.id)}>
              <div className="checkbox-container">
                <div className={`checkbox ${list.visible ? 'checked' : ''}`}>
                  {list.visible && <span className="checkmark">✓</span>}
                </div>
                <span className="list-name">{list.title}</span>
              </div>
              <span className="list-count">{list.count}</span>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

export default Sidebar
