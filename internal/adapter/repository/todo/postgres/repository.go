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

func (r *TodoPgRepository) GetTodos(sort, title, description string) ([]*core.Todo, error) {
	var todos []*core.Todo
	query := r.db.Model(&Todo{})
	switch sort {
	case "title":
		query = query.Order("title")
	case "date":
		query = query.Order("created_at")
	case "status":
		query = query.Order("status")
	default:
		query = query.Order("id")
	}

	query = query.Where("title LIKE ?", "%"+title+"%")
	query = query.Where("description LIKE ?", "%"+description+"%")

	// Execute the query and retrieve todos
	err := query.Find(&todos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
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
		CreatedAt:   newTodo.CreatedAt,
	}, nil
}

func (r *TodoPgRepository) UpdateTodo(id string, todo *core.TodoRequest) (*core.Todo, error) {
	newTodo := &Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Image:       todo.Image,
		Status:      todo.Status,
	}

	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	err = r.db.Model(&Todo{}).Where("id = ?", ID).Updates(newTodo).Error
	if err != nil {
		return nil, err
	}

	return &core.Todo{
		ID:          id,
		Title:       newTodo.Title,
		Description: newTodo.Description,
		Image:       newTodo.Image,
		Status:      newTodo.Status,
	}, nil
}
