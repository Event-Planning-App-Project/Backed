package event

type InsertEventRequest struct {
	CategoryID  uint   `json:"category_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Promotor    string `json:"promotor" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
	UrlEvent    string `json:"urlEvent" validate:"required"`
	Quota       int    `json:"quota" validate:"required"`
	DateStart   string `json:"dateStart" validate:"required"`
	DateEnd     string `json:"dateEnd" validate:"required"`
	TimeStart   string `json:"timeStart" validate:"required"`
	TimeEnd     string `json:"timeEnd" validate:"required"`
}

type UpdateEventRequest struct {
	Name        string `json:"name"`
	Promotor    string `json:"promotor"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	UrlEvent    string `json:"urlEvent"`
	Quota       int    `json:"quota"`
	DateStart   string `json:"dateStart"`
	DateEnd     string `json:"dateEnd"`
	TimeStart   string `json:"timeStart"`
	TimeEnd     string `json:"timeEnd"`
}
