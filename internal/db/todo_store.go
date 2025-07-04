package db

import (
	"github.com/NemishGorasiya/Go-Todo/internal/model"
)

func GetAllTodos() ([]model.Todo, error) {
	var todos []model.Todo
	result := DB.Find(&todos)
	return todos, result.Error
}

func CreateTodo(todo *model.Todo) error {
	return DB.Create(todo).Error
}

func UpdateTodo(id uint, updated model.Todo) error {
	var todo model.Todo
	if err := DB.First(&todo, id).Error; err != nil {
		return err
	}
	todo.Title = updated.Title
	todo.Completed = updated.Completed
	return DB.Save(&todo).Error
}

func DeleteTodo(id uint) error {
	return DB.Delete(&model.Todo{}, id).Error
}
