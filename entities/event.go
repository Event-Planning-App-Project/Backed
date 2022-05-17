package entities

import (
	"time"

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
	Ticket      int       `json:"ticket"`
	DateStart   time.Time `json:"dateStart"`
	DateEnd     time.Time `json:"dateEnd"`
	TimeStart   time.Time `json:"timeStart"`
	TimeEnd     time.Time `json:"timeEnd"`
	Comment     []Comment `gorm:"foreignKey:EventID;references:id"`
}
