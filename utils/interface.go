package utils

import "event/delivery/view/transaction"

type ConfigMidtrans interface {
	CreateTransaction(GrossAmt int64) map[string]interface{}
	FinishPayment(order string) transaction.ResponsePayment
}
