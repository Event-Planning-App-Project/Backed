package category

import (
	"event/entities"
	"net/http"
)

type RespondCategory struct {
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
	UserID     uint   `json:"user_id"`
}

func StatusGetAllOk(data []entities.Category) map[string]interface{} {
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

func StatusCreate(data entities.Category) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Category",
		"status":  true,
		"data":    data,
	}
}

func StatusUpdate(data entities.Category) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Updated",
		"status":  true,
		"data":    data,
	}
}
