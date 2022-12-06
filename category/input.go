package category

type CategoryCreateInput struct {
	Type string `json:"type" biding:"required"`
}

type CategoryUpdateInput struct {
	Type string `json:"type,omitempty" biding:"required"`
}
