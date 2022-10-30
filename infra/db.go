package infra

import (
	"fmt"

	"github.com/GenkiHirano/gin-gorm-docker-gqlgen-crud/graph/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	dsn := dbConfig()
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(model.Todo{})

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()

	if err != nil {
		fmt.Println(err)
	}

	sqlDB.Close()
}

func SeedTodo(db *gorm.DB) {
	db.Create(model.GetSeedTodo())
}
