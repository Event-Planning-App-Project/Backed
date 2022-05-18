package event

import "github.com/labstack/echo/v4"

type EventControll interface {
	CreateEvent() echo.HandlerFunc
	GetAllEvent() echo.HandlerFunc
	UpdateEvent() echo.HandlerFunc
	GetEventID() echo.HandlerFunc
	DeleteEvent() echo.HandlerFunc
}
