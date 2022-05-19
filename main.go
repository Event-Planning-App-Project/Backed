package main

import (
	"event/utils/midtrans"
	"event/utils/rds"

	controllercat "event/delivery/controller/category"
	catRepo "event/repository/category"

	controllercomm "event/delivery/controller/comment"
	comRepo "event/repository/comment"

	controllerevent "event/delivery/controller/event"
	"event/delivery/routes"
	"event/repository/event"

	cTrans "event/delivery/controller/transaction"

	"event/repository/transaction"

	controllerus "event/delivery/controller/user"
	userRepo "event/repository/user"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	// Get Access Database
	database := rds.InitDB()
	rds.Migrate()
	userRepo := userRepo.New(database)
	userControl := controllerus.New(userRepo, validator.New())

	catRepo := catRepo.NewDB(database)
	categoryControl := controllercat.NewControlCategory(catRepo, validator.New())

	comRepo := comRepo.NewDB(database)
	commControl := controllercomm.NewControlComment(comRepo, validator.New())

	eveRepo := event.NewDB(database)
	eventtControl := controllerevent.NewControlEvent(eveRepo, validator.New())

	// Send Access DB to Transaction
	snap := midtrans.InitMidtrans()

	transRepo := transaction.NewTransDB(database)
	transControl := cTrans.NewRepoTrans(transRepo, validator.New(), snap)
	// Initiate Echo
	e := echo.New()
	// Akses Path Addressss
	routes.Path(e, userControl, transControl, categoryControl, commControl, eventtControl)
	e.Logger.Fatal(e.Start(":8000"))
}
