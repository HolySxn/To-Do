import React from 'react'
import Header from './Header'

function LoadingScreen({ onToggleSidebar }) {
  return (
    <div className="app">
      <Header onToggleSidebar={onToggleSidebar} />
      <div className="main-content">
        <div className="content-area" style={{ justifyContent: 'center', alignItems: 'center', height: '100%' }}>
          <div style={{ textAlign: 'center' }}>
            <h2>Loading...</h2>
            <p>Please wait while we load your tasks.</p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default LoadingScreen
