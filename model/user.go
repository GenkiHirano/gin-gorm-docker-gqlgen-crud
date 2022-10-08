package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Age int
}

var TestUsers = []User{
	{Name: "yamada", Age: 20},
	{Name: "tanaka", Age: 21},
	{Name: "sato", Age: 22},
}
