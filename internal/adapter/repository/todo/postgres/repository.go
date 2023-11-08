package repository

import (
	core "todo-backend/internal/core/todo"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoPgRepository struct {
	db *gorm.DB
}

func NewTodoPgRepository(db *gorm.DB) *TodoPgRepository {
	return &TodoPgRepository{
		db: db,
	}
}

func (r *TodoPgRepository) GetTodos() ([]*core.Todo, error) {
	var todos []*core.Todo
	err := r.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoPgRepository) SaveTodo(todo *core.TodoRequest) (*core.Todo, error) {
	newTodo := &Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Image:       todo.Image,
		Status:      todo.Status,
	}

	// Save the newTodo to the database
	err := r.db.Create(newTodo).Error
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &core.Todo{
		ID:          newTodo.ID.String(),
		Title:       newTodo.Title,
		Description: newTodo.Description,
		Image:       newTodo.Image,
		Status:      newTodo.Status,
	}, nil
}

func (r *TodoPgRepository) UpdateTodo(id string, todo *core.TodoRequest) (*core.TodoRequest, error) {
	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	err = r.db.Model(&Todo{}).Where("id = ?", ID).Updates(&Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Image:       todo.Image,
		Status:      todo.Status,
	}).Error
	if err != nil {
		return nil, err
	}

	return todo, nil
}
