package main

import (
	"event/config"

	controllercat "event/delivery/controller/category"
	catRepo "event/repository/category"

	"event/delivery/routes"

	controllerprod "event/delivery/controller/product"
	cTrans "event/delivery/controller/transaction"
	rprod "event/repository/product"

	"event/repository/transaction"
	"event/utils"

	controllerus "event/delivery/controller/user"
	userRepo "event/repository/user"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	// Get Access Database
	database := config.InitDB()

	userRepo := userRepo.New(database)
	userControl := controllerus.New(userRepo, validator.New())

	prodrep := rprod.New(database)
	productControl := controllerprod.New(prodrep, validator.New())

	catRepo := catRepo.NewDB(database)
	categoryControl := controllercat.NewControlCategory(catRepo, validator.New())
	// Initiate Echo

	// Send Access DB to Transaction
	snap := utils.InitMidtrans()
	transRepo := transaction.NewTransDB(database)
	transControl := cTrans.NewRepoTrans(transRepo, validator.New(), snap)

	// Initiate Echo
	e := echo.New()
	// Akses Path Addressss
	routes.Path(e, userControl, transControl, categoryControl, productControl)
	e.Logger.Fatal(e.Start(":8000"))
}
