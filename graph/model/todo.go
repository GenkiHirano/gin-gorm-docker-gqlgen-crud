package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID   int
	Name string
	Text string
	Done bool
}

func GetSeedTodo() []Todo {
	return []Todo{
		{ID: 1, Name: "Todo1", Text: "Todo1のタスクです。", Done: false},
		{ID: 2, Name: "Todo2", Text: "Todo2のタスクです。", Done: false},
		{ID: 3, Name: "Todo3", Text: "Todo3のタスクです。完了しています。", Done: true},
	}
}
