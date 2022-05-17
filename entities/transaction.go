package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	OrderID       string `json:"orderID"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	EventID       uint   `json:"event_id"`
	Qty           int    `json:"qty"`
	PaymentMethod string `json:"paymentMethod"`
	TotalBill     int    `json:"totalBill"`
	Status        string `json:"status" gorm:"default:pending"`
	UserID        uint   `json:"user_id"`
}
