package transaction

import (
	"encoding/json"
	"errors"
	middlewares "event/delivery/middleware"
	"event/delivery/view/transaction"
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
		token, _ = middlewares.CreateToken(1, "babe", "babe@gmail.com")
	})
}

// TEST METHOD CREATE TRANSACTION
func TestCreateTransaction(t *testing.T) {
	t.Run("Create Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":      "user",
			"email":     "user@gmail.com",
			"phone":     "123456",
			"event_id":  1,
			"totalBill": 30000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction")
		TransactionC := NewRepoTrans(&mockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(TransactionC.CreateTransaction())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 201, result.Code)
		assert.Equal(t, "Success Create Transaction", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":      "user",
			"email":     "user@gmail.com",
			"phone":     "123456",
			"event_id":  1,
			"totalBill": 30000,
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction")
		TransactionC := NewRepoTrans(&errMockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(TransactionC.CreateTransaction())(context)

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
		requestBody := "Error Access"
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction")
		TransactionC := NewRepoTrans(&errMockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(TransactionC.CreateTransaction())(context)

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
			"name":     "user",
			"email":    "user@gmail.com",
			"phone":    "123456",
			"event_id": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction")
		TransactionC := NewRepoTrans(&errMockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(TransactionC.CreateTransaction())(context)

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
	t.Run("Error Get URL Snap", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":      "user",
			"email":     "user@gmail.com",
			"phone":     "123456",
			"event_id":  1,
			"totalBill": 30000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction")
		TransactionC := NewRepoTrans(&mockTransaction{}, validator.New(), &errMockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(TransactionC.CreateTransaction())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 204, result.Code)
		assert.Equal(t, "Error Get Redirect Url Payment", result.Message)
		assert.False(t, result.Status)
	})
}

// TEST METHOD GET ALL TRANSACTION
func TestGetAllTransaction(t *testing.T) {
	t.Run("Success Get All Transaction", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction")
		GetTransaction := NewRepoTrans(&mockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.GetAllTransaction())(context)

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
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction")
		GetTransaction := NewRepoTrans(&errMockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.GetAllTransaction())(context)

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
}

// TEST METHOD GET TRANSACTION DETAIL
func TestGetTransactionDetail(t *testing.T) {
	t.Run("Success Get Transaction Detail", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction/:order_id/")
		context.SetParamNames("order_id")
		context.SetParamValues("order123")
		GetTransaction := NewRepoTrans(&mockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.GetTransactionDetail())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get Transaction Detail", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction/:order_id")
		context.SetParamNames("order_id")
		context.SetParamValues("order123")
		GetTransaction := NewRepoTrans(&errMockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.GetTransactionDetail())(context)

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
}

// TEST METHOD PAY TRANSACTION
func TestPayTransaction(t *testing.T) {
	t.Run("Pay Transaction Success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction/:order_id/pay")
		context.SetParamNames("order_id")
		context.SetParamValues("order123")
		GetTransaction := NewRepoTrans(&mockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.PayTransaction())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Pay Transaction", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction/:order_id/pay")
		context.SetParamNames("order_id")
		context.SetParamValues("order123")
		GetTransaction := NewRepoTrans(&errMockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.PayTransaction())(context)

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
}

// METHOD TEST CANCEL TRANSCATION
func TestCancelTransaction(t *testing.T) {
	t.Run("Success Cancel Transaction", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction/:order_id/cancel")
		context.SetParamNames("order_id")
		context.SetParamValues("order123")
		GetTransaction := NewRepoTrans(&mockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.CancelTransaction())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Cancel Transaction", result.Message)
		assert.True(t, result.Status)
	})
	t.Run("Error Cancel Transaction", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction/:order_id/cancel")
		context.SetParamNames("order_id")
		context.SetParamValues("order123")
		GetTransaction := NewRepoTrans(&errMockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.CancelTransaction())(context)

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
}

// TEST FINISH PAYMENT
func TestFinishPayment(t *testing.T) {
	t.Run("Success Update Transaction", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction?order_id=order-123")

		GetTransaction := NewRepoTrans(&mockTransaction{}, validator.New(), &MockSnap{})
		GetTransaction.FinishPayment()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Update Transaction Status", result.Message)
		assert.True(t, result.Status)
	})
	t.Run("Error Cancel Transaction", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/transaction?order_id=order-123")

		GetTransaction := NewRepoTrans(&errMockTransaction{}, validator.New(), &MockSnap{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.CancelTransaction())(context)

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
}

type MockSnap struct {
}

// MOCK SUCCESS
type mockTransaction struct {
}

//METHOD MOCK SUCCESS
func (m *mockTransaction) CreateTransaction(NewTransaction entities.Transaction) (entities.Transaction, error) {
	return entities.Transaction{Name: "Galih", OrderID: "Order-1"}, nil
}
func (m *mockTransaction) GetAllTransaction(UserID uint) ([]transaction.AllTrans, error) {
	return []transaction.AllTrans{{TransDetail: transaction.RespondTransaction{OrderID: "Order-1"}, Event: transaction.EventTransaction{Name: "Training"}}}, nil
}
func (m *mockTransaction) GetTransactionDetail(UserID uint, OrderID string) (transaction.AllTrans, error) {
	return transaction.AllTrans{TransDetail: transaction.RespondTransaction{OrderID: "Order-1"}, Event: transaction.EventTransaction{Name: "Baju"}}, nil
}
func (m *mockTransaction) PayTransaction(UserID uint, OrderID string) (entities.Transaction, error) {
	return entities.Transaction{OrderID: "Order-1"}, nil
}

func (m *mockTransaction) CancelTransaction(UserID uint, OrderID string) error {
	return nil
}

func (m *mockTransaction) FinishPayment(OrderID string, updateStatus entities.Transaction) (entities.Transaction, error) {
	return entities.Transaction{OrderID: "Order-1"}, nil
}

func (m *MockSnap) CreateTransaction(OrderID string, GrossAmt int64) map[string]interface{} {
	return map[string]interface{}{"Token": "HanyaMock"}
}

func (m *MockSnap) FinishPayment(OrderID string) transaction.ResponsePayment {
	return transaction.ResponsePayment{TransactionStatus: "Settlement"}
}

// MOCK ERROR
type errMockTransaction struct {
}

// METHOD MOCK ERROR
func (e *errMockTransaction) CreateTransaction(newAdd entities.Transaction) (entities.Transaction, error) {
	return entities.Transaction{}, errors.New("Access Database Error")
}

func (e *errMockTransaction) GetAllTransaction(UserID uint) ([]transaction.AllTrans, error) {
	return nil, errors.New("Access Database Error")
}

func (e *errMockTransaction) GetTransactionDetail(UserID uint, OrderID string) (transaction.AllTrans, error) {
	return transaction.AllTrans{}, errors.New("Access Database Error")
}

func (e *errMockTransaction) PayTransaction(UserID uint, OrderID string) (entities.Transaction, error) {
	return entities.Transaction{}, errors.New("Access Database Error")
}

func (e *errMockTransaction) CancelTransaction(UserID uint, OrderID string) error {
	return errors.New("Access Database Error")
}

func (e *errMockTransaction) FinishPayment(OrderID string, updateStatus entities.Transaction) (entities.Transaction, error) {
	return entities.Transaction{}, errors.New("Access Database Error")
}

type errMockSnap struct {
}

func (e *errMockSnap) CreateTransaction(OrderID string, GrossAmt int64) map[string]interface{} {
	return nil
}

func (e *errMockSnap) FinishPayment(OrderID string) transaction.ResponsePayment {
	return transaction.ResponsePayment{}
}
