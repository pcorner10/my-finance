package models

import (
	"time"

	"github.com/pcorner10/my-finance/database"
	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	LastNumbers     uint64    `json:"last_numbers,string" binding:"required" gorm:"unique;not null"`
	UserId          int64     `json:"user_id,string" binding:"required"`
	Bank            string    `json:"bank" binding:"required"`
	IsCredit        bool      `json:"is_credit,string" binding:"required"`
	DatePayment     time.Time `json:"date_payment" binding:"required"`
	DateCutoff      time.Time `json:"date_cutoff" binding:"required"`
	CreditLimit     float64   `json:"credit_limit,string" binding:"required"`
	CreditAvailable float64   `json:"credit_available,string" binding:"required"`
}

func (cc *CreditCard) CreateCreditCard() error {
	err := database.Database.Create(cc).Error
	if err != nil {
		return err
	}
	return nil
}

func (cc *CreditCard) GetCreditCardByBank() (*CreditCard, error) {
	err := database.Database.Where("bank = ? AND is_credit = true").First(cc).Error
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (cc *CreditCard) UpdateCreditCard() (*CreditCard, error) {
	err := database.Database.Updates(cc).Error
	if err != nil {
		return nil, err
	}
	return cc, nil

}

func (cc *CreditCard) DeleteCreditCard() (*CreditCard, error) {
	err := database.Database.Delete(cc).Error
	if err != nil {
		return nil, err
	}
	return cc, nil

}

func (cc *CreditCard) GetCreditCardsByUserId() (*[]CreditCard, error) {
	var creditCards *[]CreditCard
	err := database.Database.Where("user_id = ?", cc.UserId).Find(creditCards).Error
	if err != nil {
		return nil, err
	}
	return creditCards, nil
}

// Hace referencia al tipo de operación que es; despensa, osio,
type KindOperation struct {
	gorm.Model
	Name string `json:"name_operation" binding:"required"`
}

func (ko *KindOperation) CreateKindOperation() error {
	err := database.Database.Create(ko).Error
	if err != nil {
		return err
	}
	return nil
}

func (ko *KindOperation) GetKindOperations() (*[]KindOperation, error) {
	var kindOperations *[]KindOperation
	err := database.Database.Find(kindOperations).Error
	if err != nil {
		return nil, err
	}
	return kindOperations, nil

}

func (ko *KindOperation) UpdateKindOperation() (*KindOperation, error) {
	err := database.Database.Updates(ko).Error
	if err != nil {
		return nil, err
	}
	return ko, nil

}

func (ko *KindOperation) DeleteKindOperation() (*KindOperation, error) {
	err := database.Database.Delete(ko).Error
	if err != nil {
		return nil, err
	}
	return ko, nil

}

// Se refiere a si compras unos tenis, verduras, gymnasio, audifonos
type KindProduct struct {
	gorm.Model
	Name string `json:"name_product" binding:"required"`
}

func (kp *KindProduct) CreateKindProduct() error {
	err := database.Database.Create(kp).Error
	if err != nil {
		return err
	}
	return nil
}

func (kp *KindProduct) UpdateKindProduct() (*KindProduct, error) {
	err := database.Database.Updates(kp).Error
	if err != nil {
		return nil, err
	}
	return kp, nil

}

func (kp *KindProduct) DeleteKindProduct() (*KindProduct, error) {
	err := database.Database.Delete(kp).Error
	if err != nil {
		return nil, err
	}

	return kp, nil
}

func (kp *KindProduct) GetKindProducts() (*[]KindProduct, error) {
	var kindProducts *[]KindProduct
	err := database.Database.Find(kindProducts).Error
	if err != nil {
		return nil, err
	}
	return kindProducts, nil

}

type Store struct {
	gorm.Model
	Name     string `json:"name_store" binding:"required"`
	IsOnline bool   `json:"is_online,string" binding:"required"`
}

func (s *Store) CreateStore() error {
	err := database.Database.Create(s).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateStore() (*Store, error) {
	err := database.Database.Updates(s).Error
	if err != nil {
		return nil, err
	}
	return s, nil

}

func (s *Store) DeleteStore() (*Store, error) {
	err := database.Database.Delete(s).Error
	if err != nil {
		return nil, err
	}
	return s, nil

}

func (s *Store) GetStores() (*[]Store, error) {
	var Stores *[]Store
	err := database.Database.Find(Stores).Error
	if err != nil {
		return nil, err
	}
	return Stores, nil
}

type MonthlyPayment struct {
	gorm.Model
	UserId       int64   `json:"user_id,string" binding:"required"`
	CreditCardId int64   `json:"credit_card_id,string" binding:"required"`
	MonthlyPaid  float64 `json:"monthly_paid,string" binding:"required"`
	Pending      float64 `json:"pending,string" binding:"required"`
}

func (mp *MonthlyPayment) CreateMonthlyPayment() error {
	err := database.Database.Create(mp).Error
	if err != nil {
		return err
	}
	return nil
}

func (mp *MonthlyPayment) GetMonthlyPaymentByCreditCard() (*MonthlyPayment, error) {
	err := database.Database.Where("credit_card_id = ?", mp.CreditCardId).First(mp).Error
	if err != nil {
		return nil, err
	}
	return mp, nil

}

func (mp *MonthlyPayment) UpdateMonthlyPayment() (*MonthlyPayment, error) {
	err := database.Database.Updates(mp).Error
	if err != nil {
		return nil, err
	}
	return mp, nil

}

func (mp *MonthlyPayment) DeleteMonthlyPayment() (*MonthlyPayment, error) {
	err := database.Database.Delete(mp).Error
	if err != nil {
		return nil, err
	}
	return mp, nil

}

func (mp *MonthlyPayment) GetMonthlyPayments() (*[]MonthlyPayment, error) {
	var MonthlyPayments *[]MonthlyPayment
	err := database.Database.Where("user_id = ?", mp.UserId).Find(MonthlyPayments).Error
	if err != nil {
		return nil, err
	}
	return MonthlyPayments, nil

}

type Operation struct {
	gorm.Model
	UserId          int64 `json:"user_id,string" binding:"required"`
	CreditCardId    int64 `json:"credit_card_id,string" binding:"required"`
	StoreId         int64 `json:"store_id,string" binding:"required"`
	KindOperationId int64 `json:"kind_operation_id,string" binding:"required"`
	KindProductId   int64 `json:"kind_product_id,string" binding:"required"`
	Amount          int64 `json:"amount,string" binding:"required"`
	Description     string `json:"description" binding:"required"`
	Periods         int64 `json:"periods,string" binding:"required"`
	PendingPeriods  int64 `json:"pending_periods,string" binding:"required"`
	User            User
	CreditCard      CreditCard
	Store           Store
	KindOperation   KindOperation
	KindProduct     KindProduct
}

func (o *Operation) CreateOperation() error {
	err := database.Database.Create(o).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *Operation) UpdateOperation() (*Operation, error) {
	err := database.Database.Updates(o).Error
	if err != nil {
		return nil, err
	}
	return o, nil

}

func (o *Operation) DeleteOperation() (*Operation, error) {
	err := database.Database.Delete(o).Error
	if err != nil {
		return nil, err
	}
	return o, nil

}

func (o *Operation) GetOperationsByDateRange(DateRange DateRange) (*[]Operation, error) {

	var operations *[]Operation
	err := database.Database.Where("user_id = ? AND created_at BETWEEN ? AND ?", o.UserId, DateRange.StartDate, DateRange.EndDate).Find(operations).Error

	if err != nil {
		return nil, err
	}
	return operations, nil

}

type DateRange struct {
	StartDate time.Time
	EndDate   time.Time
}
