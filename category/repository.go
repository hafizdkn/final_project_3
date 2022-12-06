package category

import "gorm.io/gorm"

type IRepository interface {
	CreateCategory(category Category) (Category, error)
	UpdateCategory(category Category) (Category, error)
	GetCategoryById(id int) (Category, error)
	GetCategorys() ([]Category, error)
	DeleteCategory(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCategory(category Category) (Category, error) {
	if err := r.db.Debug().Create(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) GetCategorys() ([]Category, error) {
	category := make([]Category, 0)

	if err := r.db.Preload("Task").Find(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) GetCategoryById(id int) (Category, error) {
	var category Category

	if err := r.db.Debug().Where("id = ?", id).First(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) UpdateCategory(category Category) (Category, error) {
	if err := r.db.Debug().Save(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) DeleteCategory(id int) error {
	if err := r.db.Debug().Delete(&Category{}, id).Error; err != nil {
		return err
	}

	return nil
}
