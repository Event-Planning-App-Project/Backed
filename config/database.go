package config

import (
	"event/entities"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := InitConfig()

	conString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.Username,
		config.Password,
		config.Address,
		config.DB_Port,
		config.Name,
	)
	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
func Migrate() {
	db := InitDB()
	db.AutoMigrate(&entities.User{}, &entities.Category{}, entities.Transaction{})
}
