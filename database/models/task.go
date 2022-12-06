package models

import "time"

type Task struct {
	ID          int          `json:"id" gorm:"primaryKey"`
	Title       string       `json:"title" gorm:"not null"`
	Description string       `json:"description" gorm:"not null"`
	Status      bool         `json:"status" gorm:"not null;type:boolean"`
	UserID      int          `json:"user_id" gorm:"not null"`
	CategoryID  int          `json:"category_id"`
	CreatedAt   *time.Time   `json:"created_at"`
	UpdatedAt   *time.Time   `json:"updated_at"`
	User        UserResponse `json:"User"`
}

type TaskResponse struct {
	ID          int        `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	UserID      int        `json:"user_id"`
	CategoryID  int        `json:"category_id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (TaskResponse) TableName() string {
	return "tasks"
}
