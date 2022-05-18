package user

import (
	middlewares "event/delivery/middleware"
	"event/delivery/view"
	userview "event/delivery/view/user"
	"event/entities"
	ruser "event/repository/user"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserController struct {
	Repo  ruser.User
	Valid *validator.Validate
}

func New(repo ruser.User, valid *validator.Validate) *UserController {
	return &UserController{
		Repo:  repo,
		Valid: valid,
	}
}

func (uc *UserController) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpUser userview.InsertUserRequest

		if err := c.Bind(&tmpUser); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := uc.Valid.Struct(&tmpUser); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		newUser := entities.User{Email: tmpUser.Email, Username: tmpUser.Username, Password: tmpUser.Password}
		res, err := uc.Repo.InsertUser(newUser)

		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		response := userview.RespondUser{Username: res.Username, Email: res.Email, UserID: res.ID}

		log.Info("berhasil insert")
		return c.JSON(http.StatusCreated, userview.SuccessInsert(response))
	}
}

func (uc *UserController) GetUserbyID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		if UserID != float64(convID) {
			return c.JSON(http.StatusNotFound, view.NotFound())
		}

		res, err := uc.Repo.GetUserID(convID)

		if err != nil {
			log.Warn()
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		response := userview.RespondUser{Username: res.Username, Email: res.Email, UserID: res.ID}

		return c.JSON(http.StatusOK, userview.StatusGetIdOk(response))
	}

}

func (uc *UserController) UpdateUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var update userview.UpdateUserRequest

		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)

		if UserID != float64(id) {
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		UpdateUser := entities.User{Email: update.Email, Password: update.Password}

		res, err := uc.Repo.UpdateUser(id, UpdateUser)

		if err != nil {
			log.Warn(err)
			notFound := "data tidak ditemukan"
			if err.Error() == notFound {
				return c.JSON(http.StatusNotFound, view.NotFound())
			}
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())

		}
		response := userview.RespondUser{Username: res.Username, Email: res.Email, UserID: res.ID}

		return c.JSON(http.StatusOK, userview.StatusUpdate(response))
	}

}
func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := userview.LoginRequest{}

		if err := c.Bind(&param); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := uc.Valid.Struct(&param); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		hasil, err := uc.Repo.Login(param.Email, param.Password)

		if err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotFound, view.NotFound())
		}

		res := userview.LoginResponse{}

		if res.Token == "" {
			token, _ := middlewares.CreateToken(float64(hasil.ID), (hasil.Email), (hasil.Username))
			res.Token = token
		}

		return c.JSON(http.StatusOK, userview.LoginOK(res))
	}
}

func (uc *UserController) DeleteUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		if UserID != float64(convID) {
			return c.JSON(http.StatusNotFound, view.NotFound())
		}

		_, erro := uc.Repo.DeleteUser(convID)

		if erro != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
