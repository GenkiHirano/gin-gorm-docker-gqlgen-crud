package main

import (
	"github.com/GenkiHirano/gin-gorm-docker-gqlgen-crud/infra"
	"github.com/GenkiHirano/gin-gorm-docker-gqlgen-crud/model"
)

func main() {
	db := infra.NewDB()
	db.Create(&model.TestUsers)
}
