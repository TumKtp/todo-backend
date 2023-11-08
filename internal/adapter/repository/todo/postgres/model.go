package repository

import (
	core "todo-backend/internal/core/todo"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string    `gorm:"type:varchar(100);not null"`
	Description string
	Image       string
	Status      core.Status `gorm:"type:todo_status;default:'IN_PROGRESS';not null"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}
