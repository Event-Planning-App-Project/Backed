package entities

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Comment string `json:"comment"`
	UserID  uint   `json:"user_id"`
	EventID uint   `json:"event_id"`
}
