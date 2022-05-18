package event

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

func TestCreateEvent(t *testing.T) {
	t.Run("Create Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"category_id": 1,
			"name":        "Kahitna Live Music",
			"promotor":    "j Entertaiment",
			"price":       120000,
			"description": "live music",
			"urlEvent":    "yiyyiuiuiu",
			"quota":       100,
			"dateStart":   "2022-05-18",
			"dateEnd":     "2022-05-29",
			"timeStart":   "17.00",
			"timeEnd":     "21.00",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/event")
		eventC := NewControlEvent(&mockEvent{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(eventC.CreateEvent())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 201, result.Code)
		assert.Equal(t, "Success Create Comment", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"category_id": 1,
			"name":        "Kahitna Live Music",
			"promotor":    "j Entertaiment",
			"price":       120000,
			"description": "live music",
			"urlEvent":    "yiyyiuiuiu",
			"quota":       100,
			"dateStart":   "2022-05-18",
			"dateEnd":     "2022-05-29",
			"timeStart":   "17.00",
			"timeEnd":     "21.00",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/event")
		eventC := NewControlEvent(&errMockEvent{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(eventC.CreateEvent())(context)

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
		context.SetPath("/event")
		eventC := NewControlEvent(&errMockEvent{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(eventC.CreateEvent())(context)

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
			"name":        "Kahitna Live Music",
			"promotor":    "j Entertaiment",
			"price":       120000,
			"description": "live music",
			"url_event":   "yiyyiuiuiu",
			"quota":       100,
			"dateStart":   "2022-05-18",
			"dateEnd":     "2022-05-29",
			"timeStart":   "17.00",
			"timeEnd":     "21.00",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/event")
		eventC := NewControlEvent(&errMockEvent{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(eventC.CreateEvent())(context)

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

func TestGetAllEvent(t *testing.T) {
	t.Run("Success Get All Event", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/event")
		Getevent := NewControlEvent(&mockEvent{}, validator.New())

		Getevent.GetAllEvent()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get All Data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/event")
		Getevent := NewControlEvent(&errMockEvent{}, validator.New())

		Getevent.GetAllEvent()(context)

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

func TestGetEventID(t *testing.T) {
	t.Run("Success Get Event By ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/event/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		GetEvent := NewControlEvent(&mockEvent{}, validator.New())

		GetEvent.GetEventID()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response

		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get Data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/event/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		GetEv := NewControlEvent(&errMockEvent{}, validator.New())

		GetEv.GetEventID()(context)

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
		context.SetPath("/event/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		GetEvent := NewControlEvent(&errMockEvent{}, validator.New())

		GetEvent.GetEventID()(context)

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

func TestUpdateEvent(t *testing.T) {
	t.Run("Update Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{

			"name":        "Kahitna Live Music",
			"promotor":    "j Entertaiment",
			"price":       120000,
			"description": "live music",
			"url_event":   "yiyyiuiuiu",
			"quota":       100,
			"dateStart":   "2022-05-18",
			"dateEnd":     "2022-05-29",
			"timeStart":   "17.00",
			"timeEnd":     "21.00",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/event/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		GetEvent := NewControlEvent(&mockEvent{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetEvent.UpdateEvent())(context)

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
		context.SetPath("/event/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		GetEv := NewControlEvent(&errMockEvent{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetEv.UpdateEvent())(context)

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
		context.SetPath("/event/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		event := NewControlEvent(&errMockEvent{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(event.UpdateEvent())(context)

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
		context.SetPath("/event/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		GetEvent := NewControlEvent(&errMockEvent{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetEvent.UpdateEvent())(context)

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

func TestDelete(t *testing.T) {
	t.Run("Success Delete Event", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/event/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		GetEv := NewControlEvent(&mockEvent{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetEv.DeleteEvent())(context)

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
	t.Run("Error Delete Event", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/event/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		GetEvent := NewControlEvent(&errMockEvent{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetEvent.DeleteEvent())(context)

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
		context.SetPath("/event/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		Getev := NewControlEvent(&errMockEvent{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Getev.DeleteEvent())(context)

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

type mockEvent struct{}

func (c *mockEvent) CreateEvent(newAdd entities.Event) (entities.Event, error) {
	return entities.Event{UserID: 1, CategoryID: 1, Name: "Kahitna Concert", Promotor: "J Entertaiment", Price: 200000, Description: "live music", UrlEvent: "6gghgh", Quota: 100, DateStart: "2022-05-18", DateEnd: "2022-05-29", TimeStart: "17.00", TimeEnd: "21.00"}, nil
}
func (c *mockEvent) GetAllEvent() ([]entities.Event, error) {
	return []entities.Event{{UserID: 1, CategoryID: 1, Name: "Kahitna Concert", Promotor: "J Entertaiment", Price: 200000, Description: "live music", UrlEvent: "6gghgh", Quota: 100, DateStart: "2022-05-18", DateEnd: "2022-05-29", TimeStart: "17.00", TimeEnd: "21.00"}}, nil
}
func (c *mockEvent) GetEventID(id uint) (entities.Event, error) {
	return entities.Event{UserID: 1, CategoryID: 1, Name: "Kahitna Concert", Promotor: "J Entertaiment", Price: 200000, Description: "live music", UrlEvent: "6gghgh", Quota: 100, DateStart: "2022-05-18", DateEnd: "2022-05-29", TimeStart: "17.00", TimeEnd: "21.00"}, nil
}
func (c *mockEvent) UpdateEvent(id uint, UpdateEvent entities.Event, UserID uint) (entities.Event, error) {
	return entities.Event{UserID: 1, CategoryID: 1, Name: "Kahitna Concert", Promotor: "J Entertaiment", Price: 200000, Description: "live music", UrlEvent: "6gghgh", Quota: 100, DateStart: "2022-05-18", DateEnd: "2022-05-29", TimeStart: "17.00", TimeEnd: "21.00"}, nil
}
func (c *mockEvent) DeleteEvent(id uint, UserID uint) error {
	return nil
}

type errMockEvent struct {
}

// METHOD MOCK ERROR
func (e *errMockEvent) CreateEvent(newAdd entities.Event) (entities.Event, error) {
	return entities.Event{}, errors.New("Access Database Error")
}

func (e *errMockEvent) GetAllEvent() ([]entities.Event, error) {
	return nil, errors.New("Access Database Error")
}

func (e *errMockEvent) GetEventID(id uint) (entities.Event, error) {
	return entities.Event{}, errors.New("Access Database Error")
}

func (e *errMockEvent) UpdateEvent(id uint, UpdateEvent entities.Event, UserID uint) (entities.Event, error) {
	return entities.Event{}, errors.New("Access Database Error")
}

func (e *errMockEvent) DeleteEvent(id uint, UserID uint) error {
	return errors.New("Access Database Error")
}
