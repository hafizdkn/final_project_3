package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	FullName  string    `json:"full_name" gorm:"not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID       int    `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	FullName string `json:"full_name,omitempty"`
}

func (UserResponse) TableName() string {
	return "users"
}
