package infra

import (
	"fmt"

	"github.com/GenkiHirano/gin-gorm-docker-gqlgen-crud/graph/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := dbConfig()
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	// TODO: model.Todoも追加
	db.AutoMigrate(model.User{})

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()

	if err != nil {
		fmt.Println(err)
	}

    defer sqlDB.Close()
}
