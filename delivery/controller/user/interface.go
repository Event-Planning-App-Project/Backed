package user

import (
	"github.com/labstack/echo/v4"
)

type ControllerUser interface {
	InsertUser() echo.HandlerFunc
	// GetAllUser(c echo.Context) error
	GetUserbyID() echo.HandlerFunc
	UpdateUserID() echo.HandlerFunc
	DeleteUserID() echo.HandlerFunc
	Login() echo.HandlerFunc
}
