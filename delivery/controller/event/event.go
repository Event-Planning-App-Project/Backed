package event

import (
	"encoding/json"
	middlewares "event/delivery/middleware"
	"event/delivery/view"
	evV "event/delivery/view/event"
	"event/entities"
	"event/repository/event"
	"event/utils/s3"
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

		data := c.FormValue("data")
		json.Unmarshal([]byte(data), &Insert)

		if err := e.Valid.Struct(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		file, err := c.FormFile("myFile")
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		src, err := file.Open()
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.NotSupported())
		}
		defer src.Close()
		result, _ := s3.UploadToS3(c, file.Filename, src)

		UserID := middlewares.ExtractTokenUserId(c)
		NewAdd := entities.Event{
			UserID:      uint(UserID),
			CategoryID:  Insert.CategoryID,
			Name:        Insert.Name,
			Promotor:    Insert.Promotor,
			Price:       Insert.Price,
			Description: Insert.Description,
			Quota:       Insert.Quota,
			UrlEvent:    result,
			DateStart:   Insert.DateStart,
			DateEnd:     Insert.DateEnd,
			TimeStart:   Insert.TimeStart,
			TimeEnd:     Insert.TimeEnd,
		}
		NewEvent, errCreate := e.Repo.CreateEvent(NewAdd)
		if errCreate != nil {
			log.Warn(errCreate)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		respond := evV.RespondEvent{
			EventID:     NewEvent.ID,
			UserID:      NewEvent.UserID,
			CategoryID:  NewEvent.CategoryID,
			Name:        NewEvent.Name,
			Promotor:    NewEvent.Promotor,
			Price:       NewEvent.Price,
			Description: NewEvent.Description,
			Quota:       NewEvent.Quota,
			UrlEvent:    NewEvent.UrlEvent,
			DateStart:   NewEvent.DateStart,
			DateEnd:     NewEvent.DateEnd,
			TimeStart:   NewEvent.TimeStart,
			TimeEnd:     NewEvent.TimeEnd,
		}
		return c.JSON(http.StatusCreated, evV.StatusCreate(respond))
	}
}

func (e *ControlEvent) GetAllEvent() echo.HandlerFunc {
	return func(c echo.Context) error {

		result, err := e.Repo.GetAllEvent()
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		var respond []evV.RespondEvent
		for _, v := range result {
			res := evV.RespondEvent{
				EventID:     v.ID,
				UserID:      v.UserID,
				CategoryID:  v.CategoryID,
				Name:        v.Name,
				Promotor:    v.Promotor,
				Price:       v.Price,
				Description: v.Description,
				Quota:       v.Quota,
				UrlEvent:    v.UrlEvent,
				DateStart:   v.DateStart,
				DateEnd:     v.DateEnd,
				TimeStart:   v.TimeStart,
				TimeEnd:     v.TimeEnd,
			}
			respond = append(respond, res)
		}
		return c.JSON(http.StatusOK, evV.StatusGetAllOk(respond))
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
		respond := evV.RespondEvent{
			EventID:     result.ID,
			UserID:      result.UserID,
			CategoryID:  result.CategoryID,
			Name:        result.Name,
			Promotor:    result.Promotor,
			Price:       result.Price,
			Description: result.Description,
			Quota:       result.Quota,
			UrlEvent:    result.UrlEvent,
			DateStart:   result.DateStart,
			DateEnd:     result.DateEnd,
			TimeStart:   result.TimeStart,
			TimeEnd:     result.TimeEnd,
		}
		return c.JSON(http.StatusOK, evV.StatusGetIdOk(respond))
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
		respond := evV.RespondEvent{
			EventID:     result.ID,
			UserID:      result.UserID,
			CategoryID:  result.CategoryID,
			Name:        result.Name,
			Promotor:    result.Promotor,
			Price:       result.Price,
			Description: result.Description,
			Quota:       result.Quota,
			UrlEvent:    result.UrlEvent,
			DateStart:   result.DateStart,
			DateEnd:     result.DateEnd,
			TimeStart:   result.TimeStart,
			TimeEnd:     result.TimeEnd,
		}
		return c.JSON(http.StatusOK, evV.StatusUpdate(respond))
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
