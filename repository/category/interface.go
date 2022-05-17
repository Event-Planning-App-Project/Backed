package category

import "event/entities"

type CategoryRepository interface {
	CreateCategory(newAdd entities.Category) (entities.Category, error)
	GetAllCategory() ([]entities.Category, error)
	GetCategoryID(id uint) (entities.Category, error)
	UpdateCat(id uint, UpdateCat entities.Category, UserID uint) (entities.Category, error)
	DeleteCat(id uint, UserID uint) error
}
