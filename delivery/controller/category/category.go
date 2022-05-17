package category

import (
	middlewares "event/delivery/middleware"
	"event/delivery/view"
	catV "event/delivery/view/category"
	"event/entities"
	"event/repository/category"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ControlCategory struct {
	Repo  category.CategoryRepository
	Valid *validator.Validate
}

func NewControlCategory(NewCat category.CategoryRepository, validate *validator.Validate) *ControlCategory {
	return &ControlCategory{
		Repo:  NewCat,
		Valid: validate,
	}
}

// ADD NEW CART
func (r *ControlCategory) CreateCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Insert catV.InsertCat
		if err := c.Bind(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := r.Valid.Struct(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		NewAdd := entities.Category{
			UserID: int(UserID),
			Name:   Insert.Name,
		}
		result, errCreate := r.Repo.CreateCategory(NewAdd)
		if errCreate != nil {
			log.Warn(errCreate)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusCreated, catV.StatusCreate(result))
	}
}

func (r *ControlCategory) GetAllCategory() echo.HandlerFunc {
	return func(c echo.Context) error {

		result, err := r.Repo.GetAllCategory()
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, catV.StatusGetAllOk(result))
	}
}

func (r *ControlCategory) GetCategoryID() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		idcat, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		result, err := r.Repo.GetCategoryID(uint(idcat))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, catV.StatusGetIdOk(result))
	}
}

func (r *ControlCategory) UpdateCat() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idcat, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		var update catV.UpdateCat
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		UpdateCat := entities.Category{
			Name: update.Name,
		}

		result, errNotFound := r.Repo.UpdateCat(uint(idcat), UpdateCat, uint(UserID))
		if errNotFound != nil {
			log.Warn(errNotFound)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, catV.StatusUpdate(result))
	}
}
func (r *ControlCategory) DeleteCat() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		catid, err := strconv.Atoi(id)

		if err != nil {

			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)

		errDelete := r.Repo.DeleteCat(uint(catid), uint(UserID))
		if errDelete != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
