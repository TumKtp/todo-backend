package routes

import (
	handler "todo-backend/internal/adapter/handler"

	"github.com/gin-gonic/gin"
)

func NewTodoRoute(g *gin.Engine, h *handler.TodoHTTPHandler) {
	routes := g.Group("/todos")

	routes.POST("/", h.CreateNewTodo)
	routes.GET("/", h.ListTodos)
	routes.PUT("/:id", h.UpdateTodoByID)
}
