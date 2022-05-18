package main

import (
	"event/config"

	controllercat "event/delivery/controller/category"
	catRepo "event/repository/category"

	controllercomm "event/delivery/controller/comment"
	comRepo "event/repository/comment"

	"event/delivery/routes"

	cTrans "event/delivery/controller/transaction"

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
	config.Migrate()

	userRepo := userRepo.New(database)
	userControl := controllerus.New(userRepo, validator.New())

	catRepo := catRepo.NewDB(database)
	categoryControl := controllercat.NewControlCategory(catRepo, validator.New())

	comRepo := comRepo.NewDB(database)
	commControl := controllercomm.NewControlComment(comRepo, validator.New())

	// Initiate Echo

	// Send Access DB to Transaction
	snap := utils.InitMidtrans()
	transRepo := transaction.NewTransDB(database)
	transControl := cTrans.NewRepoTrans(transRepo, validator.New(), snap)

	// Initiate Echo
	e := echo.New()
	// Akses Path Addressss
	routes.Path(e, userControl, transControl, categoryControl, commControl)
	e.Logger.Fatal(e.Start(":8000"))
}
