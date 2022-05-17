package category

import (
	"event/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type CategoryDB struct {
	Db *gorm.DB
}

// Get Access DB
func NewDB(db *gorm.DB) *CategoryDB {
	return &CategoryDB{
		Db: db,
	}
}

func (c *CategoryDB) CreateCategory(newAdd entities.Category) (entities.Category, error) {
	if err := c.Db.Create(&newAdd).Error; err != nil {
		log.Warn(err)
		return newAdd, err
	}
	return newAdd, nil
}
func (c *CategoryDB) GetAllCategory() ([]entities.Category, error) {
	var category []entities.Category
	if err := c.Db.Find(&category).Error; err != nil {
		log.Warn("Error Get Data", err)
		return category, err
	}
	return category, nil
}

func (c *CategoryDB) GetCategoryID(id uint) (entities.Category, error) {
	var category entities.Category
	if err := c.Db.Where("id= ?", id).Find(&category).Error; err != nil {
		log.Warn("Error Get By ID", err)
		return category, err
	}
	return category, nil
}

func (c *CategoryDB) UpdateCat(id uint, UpdateCat entities.Category, UserID uint) (entities.Category, error) {
	var category entities.Category

	if err := c.Db.Where("id =? AND user_id =?", id, UserID).First(&category).Updates(&UpdateCat).Find(&category).Error; err != nil {
		log.Warn("Update Error", err)
		return category, err
	}

	return category, nil
}

func (c *CategoryDB) DeleteCat(id uint, UserID uint) error {

	var delete entities.Category
	if err := c.Db.Where("id = ? AND user_id = ?", id, UserID).First(&delete).Delete(&delete).Error; err != nil {
		return err
	}
	return nil
}
