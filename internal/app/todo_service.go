package app

import (
	"github.com/NemishGorasiya/Go-Todo/internal/db"
	"github.com/NemishGorasiya/Go-Todo/internal/model"
)

func ListTodos() []model.Todo {
	todos, _ := db.GetAllTodos()
	return todos
}

func CreateTodo(todo model.Todo) model.Todo {
	_ = db.CreateTodo(&todo)
	return todo
}

func UpdateTodo(id uint, todo model.Todo) bool {
	return db.UpdateTodo(id, todo) == nil
}

func DeleteTodo(id uint) bool {
	return db.DeleteTodo(id) == nil
}
