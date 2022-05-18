package comment

import (
	"event/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type CommentDB struct {
	Db *gorm.DB
}

// Get Access DB
func NewDB(db *gorm.DB) *CommentDB {
	return &CommentDB{
		Db: db,
	}
}

func (cm *CommentDB) CreateCom(newAdd entities.Comment) (entities.Comment, error) {
	if err := cm.Db.Create(&newAdd).Error; err != nil {
		log.Warn(err)
		return newAdd, err
	}
	return newAdd, nil
}
func (cm *CommentDB) GetAllCom() ([]entities.Comment, error) {
	var comment []entities.Comment
	if err := cm.Db.Find(&comment).Error; err != nil {
		log.Warn("Error Get Data", err)
		return comment, err
	}
	return comment, nil
}

func (cm *CommentDB) GetCommentID(id uint) (entities.Comment, error) {
	var comment entities.Comment
	if err := cm.Db.Where("id= ?", id).Find(&comment).Error; err != nil {
		log.Warn("Error Get By ID", err)
		return comment, err
	}
	return comment, nil
}

func (cm *CommentDB) UpdateComment(id uint, UpdateComment entities.Comment, UserID uint) (entities.Comment, error) {
	var comment entities.Comment

	if err := cm.Db.Where("id =? AND user_id =?", id, UserID).First(&comment).Updates(&UpdateComment).Find(&comment).Error; err != nil {
		log.Warn("Update Error", err)
		return comment, err
	}

	return comment, nil
}

func (cm *CommentDB) DeleteComment(id uint, UserID uint) error {

	var delete entities.Category
	if err := cm.Db.Where("id = ? AND user_id = ?", id, UserID).First(&delete).Delete(&delete).Error; err != nil {
		return err
	}
	return nil
}
