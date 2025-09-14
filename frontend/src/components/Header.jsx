import React from 'react'

function Header({ onToggleSidebar }) {
  return (
    <div className="header">
      <div className="header-bar-button" onClick={onToggleSidebar}>
        <div className="header-bar"></div>
        <div className="header-bar"></div>
        <div className="header-bar"></div> 
      </div>
      <img className="header-icon" src="/src/assets/icons/checked.png" alt="Checked Icon"/>
      <h1 className="header-title">ToDo</h1>
    </div>
  )
}

export default Header
