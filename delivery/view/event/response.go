package event

import (
	"net/http"
	"time"
)

type RespondEvent struct {
	EventID     uint      `json:"event_id"`
	Name        string    `json:"name"`
	Promotor    string    `json:"promotor"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	UrlEvent    string    `json:"urlEvent"`
	Ticket      int       `json:"ticket"`
	DateStart   time.Time `json:"dateStart"`
	DateEnd     time.Time `json:"dateEnd"`
	TimeStart   time.Time `json:"timeStart"`
	TimeEnd     time.Time `json:"timeEnd"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
}

func StatusGetAllOk(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All Data",
		"status":  true,
		"data":    data,
	}
}

func StatusGetIdOk(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Data",
		"status":  true,
		"data":    data,
	}
}

func StatusCreate(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Event",
		"status":  true,
		"data":    data,
	}
}

func StatusUpdate(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Updated",
		"status":  true,
		"data":    data,
	}
}
