package category

type InsertCat struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCat struct {
	Name string `json:"name"`
}
