package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := dbConfig()
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
