package comment

import "github.com/labstack/echo/v4"

type CommentControll interface {
	CreateComment() echo.HandlerFunc
	GetAllComment() echo.HandlerFunc
	UpdateComment() echo.HandlerFunc
	GetCommentID() echo.HandlerFunc
	DeleteComment() echo.HandlerFunc
}
