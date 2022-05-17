package category

import "github.com/labstack/echo/v4"

type CategoryControl interface {
	CreateCategory() echo.HandlerFunc
	UpdateCat() echo.HandlerFunc
	GetAllCategory() echo.HandlerFunc
	GetCategoryID() echo.HandlerFunc
	DeleteCat() echo.HandlerFunc
}
