package event

import (
	"net/http"
)

type RespondEvent struct {
	EventID     uint   `json:"event_id"`
	Name        string `json:"name"`
	Promotor    string `json:"promotor"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Quota       int    `json:"quota"`
	UrlEvent    string `json:"urlEvent"`
	DateStart   string `json:"dateStart"`
	DateEnd     string `json:"dateEnd"`
	TimeStart   string `json:"timeStart"`
	TimeEnd     string `json:"timeEnd"`
	UserID      uint   `json:"user_id"`
	CategoryID  uint   `json:"category_id"`
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
