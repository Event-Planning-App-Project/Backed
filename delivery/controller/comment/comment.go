package comment

import (
	middlewares "event/delivery/middleware"
	"event/delivery/view"
	comV "event/delivery/view/comment"
	"event/entities"
	"event/repository/comment"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ControlComment struct {
	Repo  comment.CommentRepository
	Valid *validator.Validate
}

func NewControlComment(NewCom comment.CommentRepository, validate *validator.Validate) *ControlComment {
	return &ControlComment{
		Repo:  NewCom,
		Valid: validate,
	}
}

func (co *ControlComment) CreateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Insert comV.InsertComment
		if err := c.Bind(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := co.Valid.Struct(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		NewAdd := entities.Comment{
			UserID:  uint(UserID),
			EventID: Insert.EventID,
			Comment: Insert.Comment,
		}
		result, errCreate := co.Repo.CreateCom(NewAdd)
		if errCreate != nil {
			log.Warn(errCreate)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusCreated, comV.StatusCreate(result))
	}
}

func (co *ControlComment) GetAllComment() echo.HandlerFunc {
	return func(c echo.Context) error {

		result, err := co.Repo.GetAllCom()
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, comV.StatusGetAllOk(result))
	}
}

func (co *ControlComment) GetCommentID() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		idcat, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		result, err := co.Repo.GetCommentID(uint(idcat))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, comV.StatusGetIdOk(result))
	}
}

func (co *ControlComment) UpdateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idcat, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		var update comV.UpdateComment
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		UpdateCom := entities.Comment{
			Comment: update.Comment,
		}

		result, errNotFound := co.Repo.UpdateComment(uint(idcat), UpdateCom, uint(UserID))
		if errNotFound != nil {
			log.Warn(errNotFound)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, comV.StatusUpdate(result))
	}
}
func (co *ControlComment) DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		catid, err := strconv.Atoi(id)

		if err != nil {

			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)

		errDelete := co.Repo.DeleteComment(uint(catid), uint(UserID))
		if errDelete != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
