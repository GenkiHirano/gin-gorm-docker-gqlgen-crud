package main

import (
	"github.com/GenkiHirano/gin-gorm-docker-gqlgen-crud/infra"
	"github.com/GenkiHirano/gin-gorm-docker-gqlgen-crud/model"
)

func main() {
	db := infra.NewDB()
	// defer db.Close()

	db.AutoMigrate(&model.User{})
	db.Create(&model.User{Name: "yamada", Age: 20})
	db.Create(&model.User{Name: "tanaka", Age: 21})
	db.Create(&model.User{Name: "sato", Age: 22})
}
