package event

import (
	middlewares "event/delivery/middleware"
	"event/delivery/view"
	evV "event/delivery/view/event"
	"event/entities"
	"event/repository/event"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ControlEvent struct {
	Repo  event.EventRepository
	Valid *validator.Validate
}

func NewControlEvent(NewCom event.EventRepository, validate *validator.Validate) *ControlEvent {
	return &ControlEvent{
		Repo:  NewCom,
		Valid: validate,
	}
}

func (e *ControlEvent) CreateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Insert evV.InsertEventRequest
		if err := c.Bind(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := e.Valid.Struct(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		NewAdd := entities.Event{
			UserID:      uint(UserID),
			CategoryID:  Insert.CategoryID,
			Name:        Insert.Name,
			Promotor:    Insert.Promotor,
			Price:       Insert.Price,
			Description: Insert.Description,
			Quota:       Insert.Quota,
			DateStart:   Insert.DateStart,
			DateEnd:     Insert.DateEnd,
			TimeStart:   Insert.TimeStart,
			TimeEnd:     Insert.TimeEnd,
		}
		result, errCreate := e.Repo.CreateEvent(NewAdd)
		if errCreate != nil {
			log.Warn(errCreate)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusCreated, evV.StatusCreate(result))
	}
}

func (e *ControlEvent) GetAllEvent() echo.HandlerFunc {
	return func(c echo.Context) error {

		result, err := e.Repo.GetAllEvent()
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, evV.StatusGetAllOk(result))
	}
}

func (e *ControlEvent) GetEventID() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		idcat, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		result, err := e.Repo.GetEventID(uint(idcat))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, evV.StatusGetIdOk(result))
	}
}

func (e *ControlEvent) UpdateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idcat, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		var update evV.UpdateEventRequest
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		UpdateEvent := entities.Event{

			Name:        update.Name,
			Promotor:    update.Promotor,
			Price:       update.Price,
			Description: update.Description,
			UrlEvent:    update.UrlEvent,
			Quota:       update.Quota,
			DateStart:   update.DateStart,
			DateEnd:     update.DateEnd,
			TimeStart:   update.TimeStart,
			TimeEnd:     update.TimeEnd,
		}

		result, errNotFound := e.Repo.UpdateEvent(uint(idcat), UpdateEvent, uint(UserID))
		if errNotFound != nil {
			log.Warn(errNotFound)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, evV.StatusUpdate(result))
	}
}
func (e *ControlEvent) DeleteEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		catid, err := strconv.Atoi(id)

		if err != nil {

			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)

		errDelete := e.Repo.DeleteEvent(uint(catid), uint(UserID))
		if errDelete != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
