package service

import "goToDoList/internal/repository"

type TaskService struct {
	repo repository.TaskRepository
}
