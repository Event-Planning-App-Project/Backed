package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string        `json:"username"`
	Email       string        `json:"email" gorm:"unique"`
	Password    string        `json:"password" form:"password"`
	Event       []Event       `gorm:"foreignKey:UserID;references:id"`
	Transaction []Transaction `gorm:"foreignKey:UserID;references:id"`
	Comment     []Comment     `gorm:"foreignKey:UserID;references:id"`
}
