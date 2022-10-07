package infra

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := dbConfig()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}

func dbConfig() string {
	loadEnv()
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := "tcp(db:3306)"
	dbName := os.Getenv("DB_NAME")

	dsn := user + ":" + pass + "@" + host + "/" + dbName

	return dsn
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込みに失敗しました: %v", err)
	}
}
