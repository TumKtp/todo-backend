package todo

import (
	todoError "todo-backend/errors"
)

type todoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) *todoService {
	return &todoService{
		repo: repo,
	}
}

func (s *todoService) ListTodos(sort, title, description string) ([]*Todo, error) {
	result, err := s.repo.GetTodos(sort, title, description)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *todoService) CreateNewTodo(todo *TodoRequest) (*Todo, error) {
	if todo.Status != "IN_PROGRESS" && todo.Status != "COMPLETED" {
		return nil, todoError.InvalidStatus
	}
	result, err := s.repo.SaveTodo(todo)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *todoService) UpdateTodo(id string, todo *TodoRequest) (*Todo, error) {
	if todo.Status != "IN_PROGRESS" && todo.Status != "COMPLETED" {
		return nil, todoError.InvalidStatus
	}
	result, err := s.repo.UpdateTodo(id, todo)
	if err != nil {
		return nil, err
	}

	return result, nil
}
