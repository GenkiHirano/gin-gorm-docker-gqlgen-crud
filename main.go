package main

import (
	"github.com/GenkiHirano/gin-gorm-docker-gqlgen-crud/infra"
	"github.com/GenkiHirano/gin-gorm-docker-gqlgen-crud/model"
)

func main() {
	db := infra.NewDB()
	// defer db.Close()

	db.AutoMigrate(&model.User{})
	db.Create(&model.TestUsers)
}
