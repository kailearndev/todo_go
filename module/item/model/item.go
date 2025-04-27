package model

import (
	"errors"
	"todo-api/common"
)

var (
	ErrTitleIsBlank = errors.New("title cannot blank")
)

type TodoItem struct {

	// Image
	// json javascript object notasion
	common.SQLModel
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      *ItemsStatus `json:"status"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreate struct {

	// Image
	// json javascript object notasion
	Id          int          `json:"-" gorm:"column:id;"`
	Title       string       `json:"title" gorm:"column:title;"`
	Description string       `json:"description" gorm:"column:description;"`
	Status      *ItemsStatus `json:"status"`
}

func (TodoItemCreate) TableName() string {
	return TodoItem{}.TableName()
}

// update
type TodoItemUpdate struct {

	// Image
	// json javascript object notasion

	Title       *string      `json:"title" gorm:"column:title;"`
	Description *string      `json:"description" gorm:"column:description;"`
	Status      *ItemsStatus `json:"status"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
