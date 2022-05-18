package entities

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	Name        string    `json:"name"`
	Promotor    string    `json:"promotor"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	UrlEvent    string    `json:"urlEvent"`
	Quota       int       `json:"quota"`
	DateStart   string    `json:"dateStart"`
	DateEnd     string    `json:"dateEnd"`
	TimeStart   string    `json:"timeStart"`
	TimeEnd     string    `json:"timeEnd"`
	Comment     []Comment `gorm:"foreignKey:EventID;references:id"`
}
