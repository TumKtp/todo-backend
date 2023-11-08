package handler

import (
	"net/http"
	core "todo-backend/internal/core/todo"

	"github.com/gin-gonic/gin"
)

type TodoHTTPHandler struct {
	todoService core.TodoService
}

func NewTodoHTTPHandler(todoService core.TodoService) *TodoHTTPHandler {
	return &TodoHTTPHandler{
		todoService: todoService,
	}
}

// @Summary Create a new todo
// @Description Create a new todo with the specified details
// @Accept json
// @Produce json
// @Param input body core.TodoRequest true "Todo details"
// @Success 200 {object} core.Todo
// @Router /todos [post]
func (h *TodoHTTPHandler) CreateNewTodo(c *gin.Context) {
	var todo core.TodoRequest
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to create the todo
	createdTodo, err := h.todoService.CreateNewTodo(&core.TodoRequest{
		Title:       todo.Title,
		Description: todo.Description,
		Image:       todo.Image,
		Status:      todo.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	// Return the created todo as the response
	c.JSON(http.StatusOK, createdTodo)
}

// @Summary List todos
// @Description Retrieve a list of todos with optional sorting and searching
// @Accept json
// @Produce json
// @Param title query string false "Task title for searching"
// @Param description query string false "Task description for searching"
// @Param sort query string false "Sort todos by title, date, or status (title/date/status)"
// @Success 200 {array} core.Todo
// @Router /todos [get]
func (h *TodoHTTPHandler) ListTodos(c *gin.Context) {
	sort := c.Query("sort")
	title := c.Query("title")
	description := c.Query("description")
	todos, err := h.todoService.ListTodos(sort, title, description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, todos)
}

// @Summary Update a todo
// @Description Update an existing todo with the specified ID
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param input body core.TodoRequest true "Task details"
// @Success 200 {object} core.Todo
// @Router /todos/{id} [put]
func (h *TodoHTTPHandler) UpdateTodoByID(c *gin.Context) {
	var todo core.TodoRequest
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Call the service to update the todo
	id := c.Param("id")
	updatedTodo, err := h.todoService.UpdateTodo(id, &core.TodoRequest{
		Title:       todo.Title,
		Description: todo.Description,
		Image:       todo.Image,
		Status:      todo.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// Return the updated todo as the response
	c.JSON(http.StatusOK, updatedTodo)
}
