package transaction

import "github.com/labstack/echo/v4"

type TransController interface {
	CreateTransaction() echo.HandlerFunc
	GetAllTransaction() echo.HandlerFunc
	GetTransactionDetail() echo.HandlerFunc
	PayTransaction() echo.HandlerFunc
	CancelTransaction() echo.HandlerFunc
	FinishPayment() echo.HandlerFunc
}
