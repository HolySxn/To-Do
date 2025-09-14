import React from 'react'
import {
  DndContext,
  closestCenter,
  KeyboardSensor,
  PointerSensor,
  useSensor,
  useSensors,
} from '@dnd-kit/core'
import {
  arrayMove,
  SortableContext,
  sortableKeyboardCoordinates,
  verticalListSortingStrategy,
} from '@dnd-kit/sortable'
import {
  useSortable,
} from '@dnd-kit/sortable'
import { CSS } from '@dnd-kit/utilities'

// Sortable List Item Component
function SortableListItem({ list, onToggleListVisibility }) {
  const {
    attributes,
    listeners,
    setNodeRef,
    transform,
    transition,
    isDragging,
  } = useSortable({ id: list.id })

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
    opacity: isDragging ? 0.5 : 1,
  }

  return (
    <div
      ref={setNodeRef}
      style={style}
      {...attributes}
      {...listeners}
      className="list-item"
      data-sortable="true"
      data-dragging={isDragging}
      onClick={() => onToggleListVisibility(list.id)}
    >
      <div className="checkbox-container">
        <div className={`checkbox ${list.visible ? 'checked' : ''}`}>
          {list.visible && <span className="checkmark">✓</span>}
        </div>
        <span className="list-name">{list.title}</span>
      </div>
      <span className="list-count">{list.count}</span>
    </div>
  )
}

function Sidebar({ 
  sidebarOpen, 
  lists, 
  listsCollapsed, 
  onToggleLists, 
  onToggleListVisibility, 
  onCreateList,
  onReorderLists
}) {
  const sensors = useSensors(
    useSensor(PointerSensor),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    })
  )

  const handleDragEnd = (event) => {
    const { active, over } = event

    if (active.id !== over.id) {
      const oldIndex = lists.findIndex((list) => list.id === active.id)
      const newIndex = lists.findIndex((list) => list.id === over.id)

      const reorderedLists = arrayMove(lists, oldIndex, newIndex)
      onReorderLists(reorderedLists.map(list => list.id))
    }
  }
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
          <DndContext
            sensors={sensors}
            collisionDetection={closestCenter}
            onDragEnd={handleDragEnd}
          >
            <SortableContext items={lists.map(list => list.id)} strategy={verticalListSortingStrategy}>
              {lists.map((list) => (
                <SortableListItem
                  key={list.id}
                  list={list}
                  onToggleListVisibility={onToggleListVisibility}
                />
              ))}
            </SortableContext>
          </DndContext>
        </div>
      </div>
    </div>
  )
}

export default Sidebar
