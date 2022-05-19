package transaction

import (
	"errors"
	"event/delivery/view/transaction"
	"event/entities"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type TransDB struct {
	Db *gorm.DB
}

func NewTransDB(DB *gorm.DB) *TransDB {
	return &TransDB{
		Db: DB,
	}
}

// CREATE NEW TRANSACTION
func (t *TransDB) CreateTransaction(NewTransaction entities.Transaction) (entities.Transaction, error) {
	var transaction entities.Transaction
	if err := t.Db.Create(&NewTransaction).Find(&transaction).Error; err != nil {
		log.Warn(err)
		return NewTransaction, err
	}
	OrderID := fmt.Sprintf("Order-%d%d", NewTransaction.UserID, NewTransaction.ID)
	if err := t.Db.Where("id = ?", transaction.ID).Update("order_id", OrderID).Error; err != nil {
		log.Warn(err)
	}
	return NewTransaction, nil
}

// GET ALL Transaction IN DATABASE
func (t *TransDB) GetAllTransaction(UserID uint) ([]transaction.AllTrans, error) {
	var AllTransaction []entities.Transaction
	var resAllTrans []transaction.AllTrans

	if err := t.Db.Where("user_id = ?", UserID).Order("created_at DESC").Find(&AllTransaction).Error; err != nil {
		log.Warn("Error Get All Transaction", err)
		return resAllTrans, errors.New("Access Database Error")
	}

	for _, v := range AllTransaction {
		var resTrans transaction.AllTrans
		resTrans.TransDetail = transaction.RespondTransaction{
			OrderID:       v.OrderID,
			Name:          v.Name,
			Email:         v.Email,
			Phone:         v.Phone,
			EventID:       v.EventID,
			TotalBill:     v.TotalBill,
			PaymentMethod: v.PaymentMethod,
			Status:        v.Status,
			CreatedAt:     v.CreatedAt,
		}
		var events entities.Event
		if err := t.Db.Where("id=?", v.EventID).Find(&events).Error; err != nil {
			log.Warn("Error Get All Transaction", err)
			return resAllTrans, errors.New("Access Database Error")
		}

		Event := transaction.EventTransaction{
			EventID:  events.ID,
			Name:     events.Name,
			Promotor: events.Promotor,
			Price:    events.Price,
			UrlEvent: events.UrlEvent,
		}

		resTrans.Event = Event
		resAllTrans = append(resAllTrans, resTrans)
	}
	return resAllTrans, nil
}

// GET Transaction BY ID
func (t *TransDB) GetTransactionDetail(UserID uint, OrderID string) (transaction.AllTrans, error) {
	resAllTrans, err := t.GetAllTransaction(UserID)
	if err != nil {
		log.Warn(err)
		return transaction.AllTrans{}, err
	}
	for _, v := range resAllTrans {
		if v.TransDetail.OrderID == OrderID {
			return v, nil
		}
	}
	return transaction.AllTrans{}, errors.New("Get Transaction Error")
}

// UPDATE Transaction BY ID
func (t *TransDB) PayTransaction(UserID uint, OrderID string) (entities.Transaction, error) {
	var updated entities.Transaction

	if err := t.Db.Where("user_id =? AND order_id=?", UserID, OrderID).First(&updated).Update("status", "success").Error; err != nil {
		log.Warn("Pay Transaction Error", err)
		return updated, errors.New("Access Database Error")
	}
	return updated, nil
}

// DELETE Transaction BY ID
func (t *TransDB) CancelTransaction(UserID uint, OrderID string) error {
	var Cancel entities.Transaction
	if err := t.Db.Where("user_id = ? AND order_id = ?", UserID, OrderID).First(&Cancel).Update("status", "failured").Error; err != nil {
		log.Warn("Cancel Transaction Error")
		return err
	}
	return nil
}

// FINISH PAYMENT
func (t *TransDB) FinishPayment(OrderID string, updateStatus entities.Transaction) (entities.Transaction, error) {
	var result entities.Transaction
	if err := t.Db.Where("order_id = ?", OrderID).Updates(&updateStatus).Find(&result).Error; err != nil {
		log.Warn(err)
		return updateStatus, errors.New("Access Database Error")
	}
	return result, nil
}
