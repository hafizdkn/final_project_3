package models

import "time"

type Category struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Type      string         `json:"type" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Task      []TaskResponse `json:"Task"`
}
