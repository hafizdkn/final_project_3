package user

import (
	"time"

	"final_project_3/database/models"
)

type User models.User

type UserResponse struct {
	ID        int        `json:"id,omitempty"`
	FullName  string     `json:"full_name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Age       int        `json:"age,omitempty"`
	Token     string     `json:"token,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
