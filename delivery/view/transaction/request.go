package transaction

type InsertTransaction struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	EventID   uint   `json:"event_id" validate:"required"`
	TotalBill int    `json:"totalBill" validate:"required"`
}

type InsertStatusTransaction struct {
	Status string `json:"status" validate:"required"`
}
