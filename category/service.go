package category

import "errors"

type IService interface {
	CreateCategory(input CategoryCreateInput, role string) (CategoryResponse, error)
	UpdateCategory(input CategoryUpdateInput, categoryId, userId int) (CategoryResponse, error)
	GetCategoryById(categoryId int) (Category, error)
	DeleteCategory(categoryId int) error
	GetCategorys() ([]Category, error)
}

type service struct {
	repository IRepository
}

func NewCategoryService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) CreateCategory(input CategoryCreateInput, role string) (CategoryResponse, error) {
	var category Category
	var categoryResponse CategoryResponse

	if role != "admin" {
		return categoryResponse, errors.New("Unauthorized")
	}

	category.Type = input.Type
	category, err := s.repository.CreateCategory(category)
	if err != nil {
		return categoryResponse, err
	}

	categoryResponse = CategoryResponse{
		ID:        categoryResponse.ID,
		Type:      category.Type,
		CreatedAt: &category.CreatedAt,
	}

	return categoryResponse, nil
}

func (s *service) GetCategoryById(categoryId int) (Category, error) {
	category, err := s.repository.GetCategoryById(categoryId)
	if err != nil {
		return category, err
	}

	// if category.ID != userId {
	// 	return category, errors.New("Unauthorized")
	// }

	return category, nil
}

func (s *service) UpdateCategory(input CategoryUpdateInput, categoryId, userId int) (CategoryResponse, error) {
	var categoryResponse CategoryResponse

	category, err := s.GetCategoryById(categoryId)
	if err != nil {
		return categoryResponse, err
	}

	category.Type = input.Type

	categoryResponse = CategoryResponse{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.ID * 5,
		UpdatedAt:         &category.UpdatedAt,
	}

	return categoryResponse, nil
}

func (s *service) DeleteCategory(categoryId int) error {
	if _, err := s.GetCategoryById(categoryId); err != nil {
		return err
	}

	if err := s.repository.DeleteCategory(categoryId); err != nil {
		return err
	}

	return nil
}

func (s *service) GetCategorys() ([]Category, error) {
	categorys, err := s.repository.GetCategorys()
	if err != nil {
		return categorys, err
	}

	return categorys, nil
}
