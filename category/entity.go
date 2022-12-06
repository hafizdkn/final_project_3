package category

import (
	"time"

	"final_project_3/database/models"
)

type Category models.Category

type CategoryResponse struct {
	ID                int        `json:"id,omitempty"`
	Type              string     `json:"type,omitempty"`
	SoldProductAmount int        `json:"sold_product_amount,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}
