package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	todoPgRepo "todo-backend/internal/adapter/repository/todo/postgres"
)

var db *gorm.DB

func InitDB(connectionStr string) {
	var err error

	db, err = gorm.Open(postgres.Open(connectionStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("err connect db : ", err)
		log.Fatal(err)
	}

}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	_db, _ := db.DB()
	_db.Close()
}

func MigrageDB() {
	// Create enum type
	db.Exec("CREATE TYPE todo_status AS ENUM ('IN_PROGRESS', 'COMPLETED');")
	err := db.AutoMigrate(&todoPgRepo.Todo{})
	if err != nil {
		panic(err)
	}
}
