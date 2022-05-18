package comment

import "event/entities"

type CommentRepository interface {
	CreateCom(newAdd entities.Comment) (entities.Comment, error)
	GetAllCom() ([]entities.Comment, error)
	GetCommentID(id uint) (entities.Comment, error)
	UpdateComment(id uint, UpdateComment entities.Comment, UserID uint) (entities.Comment, error)
	DeleteComment(id uint, UserID uint) error
}
