package controllers

import "goPractice/task_manager/usecases"

type TaskController struct {
	taskUsecase usecases.TaksUsecase
}

func NewTaskCotroller(taskUsecase usecases.TaksUsecase) *TaskController {
	return &TaskController{taskUsecase: taskUsecase}
}
