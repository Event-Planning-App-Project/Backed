package event

import "time"

type InsertEventRequest struct {
	Name        string    `json:"name" validate:"required"`
	Promotor    string    `json:"promotor" validate:"required"`
	Price       int       `json:"price" validate:"required"`
	Description string    `json:"description" validate:"required"`
	UrlEvent    string    `json:"urlEvent" validate:"required"`
	Ticket      int       `json:"ticket" validate:"required"`
	DateStart   time.Time `json:"dateStart" validate:"required"`
	DateEnd     time.Time `json:"dateEnd" validate:"required"`
	TimeStart   time.Time `json:"timeStart" validate:"required"`
	TimeEnd     time.Time `json:"timeEnd" validate:"required"`
	CategoryID  uint      `json:"category_id" validate:"required"`
}

type UpdateEventRequest struct {
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
}
