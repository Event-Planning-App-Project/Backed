package category

import (
	"encoding/json"
	"errors"
	middlewares "event/delivery/middleware"
	"event/entities"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var token string

// INITIATE TOKEN
func TestCreateToken(t *testing.T) {
	t.Run("Create Token", func(t *testing.T) {
		token, _ = middlewares.CreateToken(1, "Galih", "Galih@gmail.com")
	})
}

func TestCreateCategory(t *testing.T) {
	t.Run("Create Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"user_id": 1,
			"name":    "kecantikan",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category")
		categoryC := NewControlCategory(&mockCategory{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(categoryC.CreateCategory())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 201, result.Code)
		assert.Equal(t, "Success Create Category", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"user_id": 1,
			"name":    "kecantikan",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category")
		categoryC := NewControlCategory(&errMockCategory{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(categoryC.CreateCategory())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()

		requestBody := "kecantikan"

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category")
		categoryC := NewControlCategory(&errMockCategory{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(categoryC.CreateCategory())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 415, result.Code)
		assert.Equal(t, "Cannot Bind Data", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Validate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama": "kecantikan",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category")
		categoryC := NewControlCategory(&errMockCategory{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(categoryC.CreateCategory())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Validate Error", result.Message)
		assert.False(t, result.Status)
	})
}

func TestGetAllCategory(t *testing.T) {
	t.Run("Success Get All Category", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category")
		Getcategory := NewControlCategory(&mockCategory{}, validator.New())

		Getcategory.GetAllCategory()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get All data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category")
		Getcategory := NewControlCategory(&errMockCategory{}, validator.New())

		Getcategory.GetAllCategory()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
}

func TestGetCategoryID(t *testing.T) {
	t.Run("Success Get Category By ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		GetAddress := NewControlCategory(&mockCategory{}, validator.New())

		GetAddress.GetCategoryID()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response

		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get Data ID", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		GetCat := NewControlCategory(&errMockCategory{}, validator.New())

		GetCat.GetCategoryID()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Convert ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		GetCat := NewControlCategory(&errMockCategory{}, validator.New())

		GetCat.GetCategoryID()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
}

func TestUpdateAddress(t *testing.T) {
	t.Run("Update Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name": "kecantikan",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		GetCat := NewControlCategory(&mockCategory{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCat.UpdateCat())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Updated", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		GetCat := NewControlCategory(&errMockCategory{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCat.UpdateCat())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody := "kecantikan"
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		catego := NewControlCategory(&errMockCategory{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(catego.UpdateCat())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 415, result.Code)
		assert.Equal(t, "Cannot Bind Data", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Convert ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		Getcategory := NewControlCategory(&errMockCategory{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Getcategory.UpdateCat())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("Success Delete Category", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		GetCat := NewControlCategory(&mockCategory{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCat.DeleteCat())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Deleted", result.Message)
		assert.True(t, result.Status)
	})
	t.Run("Error Delete Category", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		GetCategory := NewControlCategory(&errMockCategory{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCategory.DeleteCat())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Convert ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		Getcat := NewControlCategory(&errMockCategory{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Getcat.DeleteCat())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
}

type mockCategory struct{}

func (m *mockCategory) CreateCategory(newAdd entities.Category) (entities.Category, error) {
	return entities.Category{Name: "kecantikan"}, nil
}
func (m *mockCategory) GetAllCategory() ([]entities.Category, error) {
	return []entities.Category{{Name: "kecantikan"}, {Name: "kesehatan"}}, nil
}
func (m *mockCategory) GetCategoryID(id uint) (entities.Category, error) {
	return entities.Category{Name: "kecantikan"}, nil
}
func (m *mockCategory) UpdateCat(id uint, UpdateCat entities.Category, UserID uint) (entities.Category, error) {
	return entities.Category{Name: "kecantikan"}, nil
}
func (m *mockCategory) DeleteCat(id uint, UserID uint) error {
	return nil
}

type errMockCategory struct {
}

// METHOD MOCK ERROR
func (e *errMockCategory) CreateCategory(newAdd entities.Category) (entities.Category, error) {
	return entities.Category{}, errors.New("Access Database Error")
}

func (e *errMockCategory) GetAllCategory() ([]entities.Category, error) {
	return nil, errors.New("Access Database Error")
}

func (e *errMockCategory) GetCategoryID(id uint) (entities.Category, error) {
	return entities.Category{}, errors.New("Access Database Error")
}

func (e *errMockCategory) UpdateCat(id uint, UpdateCat entities.Category, UserID uint) (entities.Category, error) {
	return entities.Category{}, errors.New("Access Database Error")
}

func (e *errMockCategory) DeleteCat(id uint, UserID uint) error {
	return errors.New("Access Database Error")
}
