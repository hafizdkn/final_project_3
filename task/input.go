package task

type TaskCreateInput struct {
	Title       string `json:"title" biding:"required"`
	Description string `json:"description" biding:"required"`
	CategoryId  int    `json:"category_id" biding:"required"`
}

type TaskUpdateInput struct {
	Title       string `json:"title,omitempty" biding:"required"`
	Description string `json:"description,omitempty" biding:"required"`
}

type TaskUpdateStatusInput struct {
	Status string `json:"status" binding:"required"`
}

type TaskUpdateCategoryInput struct {
	CategoryID int `json:"category_id" binding:"required"`
}
