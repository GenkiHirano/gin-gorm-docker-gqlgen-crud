package infra

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func dbConfig() string {
	loadEnv()
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := "tcp(127.0.0.1:3306)"
	dbName := os.Getenv("DB_NAME")

	dsn := user + ":" + pass + "@" + host + "/" + dbName

	return dsn
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed to read environment variables: %v", err)
	}
}
