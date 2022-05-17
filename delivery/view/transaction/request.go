package transaction

type InsertTransaction struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	EventID   uint   `json:"event_id"`
	Qty       int    `json:"qty"`
	TotalBill int    `json:"totalBill"`
}

type InsertStatusTransaction struct {
	Status string `json:"status" validate:"required"`
}
