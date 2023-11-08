package main

import (
	"todo-backend/config"
	_ "todo-backend/docs"
	handler "todo-backend/internal/adapter/handler"
	todoPgRepo "todo-backend/internal/adapter/repository/todo/postgres"
	todoService "todo-backend/internal/core/todo"
	"todo-backend/routes"
	"todo-backend/store/postgres"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	gotenv.Load()
}

// @title			Todo API
// @version		1.0
// @description	This is a todo API server.
func main() {
	// Environment config
	appConfig := config.AppConfig()

	r := gin.Default()
	postgres.InitDB(appConfig.DatabaseUrl)
	db := postgres.GetDB()
	postgres.MigrageDB()
	defer postgres.CloseDB()

	todoPgRepo := todoPgRepo.NewTodoPgRepository(db)
	todoService := todoService.NewTodoService(todoPgRepo)
	todoHandler := handler.NewTodoHTTPHandler(todoService)

	routes.NewTodoRoute(r, todoHandler)

	// Swagger documentation route
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run(":8080")
}
