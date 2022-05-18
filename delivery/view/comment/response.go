package comment

import (
	"event/entities"
	"net/http"
)

type RespondComment struct {
	EventID uint   `json:"event_id"`
	UserID  uint   `json:"user_id"`
	Comment string `json:"comment"`
}

func StatusGetAllOk(data []entities.Comment) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All data",
		"status":  true,
		"data":    data,
	}
}

func StatusGetIdOk(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Data ID",
		"status":  true,
		"data":    data,
	}
}

func StatusCreate(data entities.Comment) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Comment",
		"status":  true,
		"data":    data,
	}
}

func StatusUpdate(data entities.Comment) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Updated",
		"status":  true,
		"data":    data,
	}
}
