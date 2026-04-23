package repository

import "goToDoList/internal/model"

type TaskRepository interface {
	createTask(title model.Task) (int, error)
	getAll() []model.Task
	getByID(id int) (model.Task, bool)
	updateTask(id int, updatedTask model.Task) bool
	deleteTask(id int) bool
}
