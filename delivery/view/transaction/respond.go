package transaction

import (
	"net/http"
	"time"
)

type RespondTransaction struct {
	OrderID       string    `json:"order_id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	EventID       uint      `json:"event_id"`
	TotalBill     int       `json:"totalBill"`
	PaymentMethod string    `json:"paymentMethod"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
}

type EventTransaction struct {
	EventID  uint   `json:"event_id"`
	Name     string `json:"Name"`
	Promotor string `json:"Promotor"`
	Price    int    `json:"price"`
	UrlEvent string `json:"urlEvent"`
}

type AllTrans struct {
	TransDetail RespondTransaction
	Event       EventTransaction
}

type ResponsePayment struct {
	StatusCode        string `json:"status_code"`
	PaymentType       string `json:"payment_type"`
	TransactionStatus string `json:"transaction_status"`
}

func StatusGetAllOk(data []AllTrans) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All data",
		"status":  true,
		"data":    data,
	}
}

func StatusTransactionDetail(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Transaction Detail",
		"status":  true,
		"data":    data,
	}
}

func StatusCreate(OrderID string, snap map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Transaction",
		"status":  true,
		"data":    map[string]interface{}{"order-id": OrderID, "RedirectUrl": snap},
	}
}

func StatusPayTrans(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Pay Transaction",
		"status":  true,
		"data":    data,
	}
}

func StatusCancelTrans() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Cancel Transaction",
		"status":  true,
	}
}

func StatusErrorSnap() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNoContent,
		"message": "Error Get Redirect Url Payment",
		"status":  false,
	}
}

func StatusUpdateTransaction(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Update Transaction Status",
		"status":  true,
		"data":    data,
	}
}
