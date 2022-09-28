package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	DBName     string
	Username   string
	Password   string
	Connection *gorm.DB
}

func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		DBName:   c.DBName,
		Username: c.Username,
		Password: c.Password,
	})
}

func newDB(d *DB) *DB {
	dsn := d.Username + ":" + d.Password + "@tcp(127.0.0.1:3306)/" + d.DBName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err.Error())
	}

	d.Connection = db
	return d
}

func (db *DB) Begin() *gorm.DB {
    return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
    return db.Connection
}
