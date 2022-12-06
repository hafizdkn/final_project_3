package task

import (
	"gorm.io/gorm"
)

type IRepository interface {
	CreateTask(task Task) (Task, error)
	UpdateTask(task Task) (Task, error)
	GetTaskById(id int) (Task, error)
	GetTaskByUserId(id int) (Task, error)
	GetTasks() ([]Task, error)
	DeleteTask(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetTaskByUserId(id int) (Task, error) {
	var task Task

	if err := r.db.Debug().Where("user_id = ?", id).First(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) CreateTask(task Task) (Task, error) {
	if err := r.db.Debug().Create(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) GetTasks() ([]Task, error) {
	taks := make([]Task, 0)

	if err := r.db.Preload("User").Find(&taks).Error; err != nil {
		return taks, err
	}

	return taks, nil
}

func (r *repository) GetTaskById(id int) (Task, error) {
	var task Task

	if err := r.db.Debug().Where("id = ?", id).First(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) UpdateTask(task Task) (Task, error) {
	if err := r.db.Debug().Save(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) DeleteTask(id int) error {
	_, err := r.GetTaskById(id)
	if err != nil {
		return err
	}

	if err := r.db.Debug().Delete(&Task{}, id).Error; err != nil {
		return err
	}

	return nil
}
