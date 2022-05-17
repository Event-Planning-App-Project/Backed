package user

type InsertUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required" gorm:"unique"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type UpdateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
