package task

import (
	"time"

	"final_project_3/database/models"
)

type Task models.Task

type TaskResponse struct {
	ID          int        `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Status      bool       `json:"status"`
	Description string     `json:"description,omitempty"`
	UserID      int        `json:"user_id,omitempty"`
	CategoryID  int        `json:"category_id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
