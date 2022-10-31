package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID   int `gorm:"AUTO_INCREMENT"`
	Name string
	Text string
	Done bool
}

func GetSeedTodo() []Todo {
	return []Todo{
		{Name: "Todo1", Text: "Todo1のタスクです", Done: false},
		{Name: "Todo2", Text: "Todo2のタスクです", Done: false},
		{Name: "Todo3", Text: "Todo3のタスクです", Done: true},
	}
}
