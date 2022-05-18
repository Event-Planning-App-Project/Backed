package comment

type InsertComment struct {
	EventID uint   `json:"event_id" validate:"required"`
	Comment string `json:"comment"`
}

type UpdateComment struct {
	Comment string `json:"comment"`
}
