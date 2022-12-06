package task

import (
	"errors"
	"strconv"

	"final_project_3/category"
)

type Iservice interface {
	CreateTask(input TaskCreateInput, userId int) (TaskResponse, error)
	UpdateTask(input TaskUpdateInput, taskId int) (TaskResponse, error)
	UpdateStatusTask(input TaskUpdateStatusInput, taskId, currentUserId int) (TaskResponse, error)
	UpdateCategoryTask(input TaskUpdateCategoryInput, taskId, currentUserId int) (TaskResponse, error)
	DeleteTask(taskId int) error
	GetTasks() ([]Task, error)
}

type service struct {
	repository      IRepository
	serviceCategory category.IService
}

func NewCategoryServie(repository IRepository, serviceCategory category.IService) *service {
	return &service{repository: repository, serviceCategory: serviceCategory}
}

func (s *service) UpdateCategoryTask(input TaskUpdateCategoryInput, taskId, currentUserId int) (TaskResponse, error) {
	var taskResponse TaskResponse
	valUpdateCategoryId := input.CategoryID

	if currentUserId == 0 && taskId == 0 && valUpdateCategoryId == 0 {
		return taskResponse, errors.New("Error data not found")
	}

	_, err := s.serviceCategory.GetCategoryById(taskId)
	if err != nil {
		return taskResponse, err
	}

	task, err := s.GetTaskById(taskId, currentUserId)
	if err != nil {
		return taskResponse, err
	}

	task.CategoryID = valUpdateCategoryId

	updatedTask, err := s.repository.UpdateTask(task)
	if err != nil {
		return taskResponse, err
	}

	return parseToTaskResponse(updatedTask), nil
}

func (s *service) GetTaskById(taskId, currentUserId int) (Task, error) {
	task, err := s.repository.GetTaskById(taskId)
	if err != nil {
		return task, err
	}

	if task.UserID != currentUserId {
		return task, errors.New("Unauthorized")
	}

	return task, nil
}

func (s *service) UpdateStatusTask(input TaskUpdateStatusInput, taskId, currentUserId int) (TaskResponse, error) {
	var taskResponse TaskResponse

	if taskId == 0 {
		return taskResponse, errors.New("Error data not found")
	}

	task, err := s.GetTaskById(taskId, currentUserId)
	if err != nil {
		return taskResponse, err
	}

	status, err := strconv.ParseBool(input.Status)
	if err != nil {
		return taskResponse, err
	}

	task.Status = status

	updatedTask, err := s.repository.UpdateTask(task)
	if err != nil {
		return taskResponse, err
	}

	return parseToTaskResponse(updatedTask), nil
}

func (s *service) DeleteTask(taskId int) error {
	if taskId == 0 {
		return errors.New("Error data not found")
	}

	err := s.repository.DeleteTask(taskId)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateTask(input TaskUpdateInput, taskId int) (TaskResponse, error) {
	var taskResponse TaskResponse

	task, err := s.repository.GetTaskById(taskId)
	if err != nil {
		return taskResponse, err
	}

	task.Title = input.Title
	task.Description = input.Description

	updatedTask, err := s.repository.UpdateTask(task)
	if err != nil {
		return taskResponse, err
	}

	return parseToTaskResponse(updatedTask), nil
}

func (s *service) GetTasks() ([]Task, error) {
	tasks, err := s.repository.GetTasks()
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (s *service) CreateTask(input TaskCreateInput, userId int) (TaskResponse, error) {
	var task Task
	var taskResponse TaskResponse

	// melakukan pengecekan id di tabel kategori apakah sudah terbuat?
	// jika belum return err , jika sudah buat task
	_, err := s.serviceCategory.GetCategoryById(input.CategoryId)
	if err != nil {
		return taskResponse, err
	}

	task = Task{
		Title:       input.Title,
		Description: input.Description,
		CategoryID:  input.CategoryId,
		UserID:      userId,
		Status:      false,
	}

	task, err = s.repository.CreateTask(task)
	if err != nil {
		return taskResponse, err
	}

	return parseToTaskResponse(task), nil
}

func parseToTaskResponse(task Task) TaskResponse {
	taskResponse := TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Status:      task.Status,
		Description: task.Description,
		UserID:      task.UserID,
		CategoryID:  task.CategoryID,
		CreatedAt:   task.CreatedAt,
	}

	return taskResponse
}
