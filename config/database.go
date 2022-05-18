package config

import (
	"event/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// config := InitConfig()

	// conString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
	// 	config.Username,
	// 	config.Password,
	// 	config.Address,
	// 	config.DB_Port,
	// 	config.Name,
	// )
	conString := "root:Kuroko25nara@tcp(db-learn.cb8meadbge6r.ap-southeast-1.rds.amazonaws.com:3306)/Event?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
func Migrate() {
	db := InitDB()
	db.AutoMigrate(&entities.User{}, &entities.Category{}, entities.Transaction{}, entities.Comment{})
}
