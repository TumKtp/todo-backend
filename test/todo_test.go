package todo

import (
	"testing"
	todoError "todo-backend/errors"
	core "todo-backend/internal/core/todo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) SaveTodo(todo *core.TodoRequest) (*core.Todo, error) {
	args := m.Called(todo)
	return args.Get(0).(*core.Todo), args.Error(1)
}

func (m *MockTodoRepository) UpdateTodo(id string, todo *core.TodoRequest) (*core.Todo, error) {
	args := m.Called(id, todo)
	return args.Get(0).(*core.Todo), args.Error(1)
}

func (m *MockTodoRepository) GetTodos(sort, title, description string) ([]*core.Todo, error) {
	args := m.Called(sort, title, description)
	return args.Get(0).([]*core.Todo), args.Error(1)
}

func TestCreateNewTodo_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockTodoRepository)
	service := core.NewTodoService(mockRepo)
	inputTodo := &core.TodoRequest{
		Title:  "Sample Todo",
		Status: "IN_PROGRESS",
	}

	expectedTodo := &core.Todo{
		Title:  inputTodo.Title,
		Status: inputTodo.Status,
	}

	mockRepo.On("SaveTodo", inputTodo).Return(expectedTodo, nil)

	// Act
	resultTodo, err := service.CreateNewTodo(inputTodo)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, resultTodo)
	assert.Equal(t, expectedTodo, resultTodo)

	// Assert that the expected method was called on the mock repository
	mockRepo.AssertExpectations(t)
}

func TestCreateNewTodo_InvalidStatus(t *testing.T) {
	// Arrange
	mockRepo := new(MockTodoRepository)
	service := core.NewTodoService(mockRepo)
	inputTodo := &core.TodoRequest{
		Title:  "Sample Todo",
		Status: "INVALID_STATUS",
	}

	// Act
	resultTodo, err := service.CreateNewTodo(inputTodo)

	// Assert
	assert.Nil(t, resultTodo)
	assert.Equal(t, todoError.InvalidStatus, err)

	// Assert that the SaveTodo method was not called on the mock repository
	mockRepo.AssertNotCalled(t, "SaveTodo")
}

func TestUpdateTodo_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockTodoRepository)
	service := core.NewTodoService(mockRepo)
	inputTodo := &core.TodoRequest{
		Title:  "Sample Todo",
		Status: "IN_PROGRESS",
	}

	expectedTodo := &core.Todo{
		Title:  inputTodo.Title,
		Status: inputTodo.Status,
	}

	mockRepo.On("UpdateTodo", "1", inputTodo).Return(expectedTodo, nil)

	// Act
	resultTodo, err := service.UpdateTodo("1", inputTodo)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, resultTodo)
	assert.Equal(t, expectedTodo, resultTodo)

	// Assert that the expected method was called on the mock repository
	mockRepo.AssertExpectations(t)
}

func TestUpdateTodo_InvalidStatus(t *testing.T) {
	// Arrange
	mockRepo := new(MockTodoRepository)
	service := core.NewTodoService(mockRepo)
	inputTodo := &core.TodoRequest{
		Title:  "Sample Todo",
		Status: "INVALID_STATUS",
	}

	// Act
	resultTodo, err := service.UpdateTodo("1", inputTodo)

	// Assert
	assert.Nil(t, resultTodo)
	assert.Equal(t, todoError.InvalidStatus, err)

	// Assert that the SaveTodo method was not called on the mock repository
	mockRepo.AssertNotCalled(t, "UpdateTodo")
}

func TestListTodos_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockTodoRepository)
	service := core.NewTodoService(mockRepo)

	expectedTodos := []*core.Todo{
		{
			Title:  "Sample Todo 1",
			Status: "IN_PROGRESS",
		},
		{
			Title:  "Sample Todo 2",
			Status: "COMPLETED",
		},
	}

	mockRepo.On("GetTodos", "", "", "").Return(expectedTodos, nil)

	// Act
	resultTodos, err := service.ListTodos("", "", "")

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, resultTodos)
	assert.Equal(t, expectedTodos, resultTodos)

	// Assert that the expected method was called on the mock repository
	mockRepo.AssertExpectations(t)
}
