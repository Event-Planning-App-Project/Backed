package transaction

import (
	"event/delivery/view/transaction"
	"event/entities"
)

type RepoTrans interface {
	CreateTransaction(NewTransaction entities.Transaction) (entities.Transaction, error)
	GetAllTransaction(UserID uint) ([]transaction.AllTrans, error)
	GetTransactionDetail(UserID uint, OrderID string) (transaction.AllTrans, error)
	PayTransaction(UserID uint, OrderID string) (entities.Transaction, error)
	CancelTransaction(UserID uint, OrderID string) error
	FinishPayment(OrderID string, updateStatus entities.Transaction) (entities.Transaction, error)
}
