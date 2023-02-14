package models

import (
	"time"

	"github.com/pcorner10/my-finance/database"
	"gorm.io/gorm"
)

type CreditCard struct {
	LastNumbers     uint64
	Bank            string
	IsCredit        bool
	DatePayment     time.Time
	DateCutoff      time.Time
	CreditLimit     float64
	CreditAvailable float64
	gorm.Model
}

func (cc *CreditCard) CreateCreditCard() (*CreditCard, error) {
	err := database.Database.Create(cc).Error
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (cc *CreditCard) UpdateCreditCard() (*CreditCard, error) {

}

func (cc *CreditCard) GetCreditCards() (*[]CreditCard, error) {

}

func (cc *CreditCard) GetCreditCard() (*CreditCard, error) {

}

func (cc *CreditCard) DeleteCreditCard() (*CreditCard, error) {

}

// Hace referencia al tipo de operación que es; despensa, osio,
type KindOperation struct {
	gorm.Model
	Name int64
}

func (ko *KindOperation) CreateKindOperation() (*KindOperation, error) {

}

func (ko *KindOperation) GetKindOperation() (*KindOperation, error) {

}

func (ko *KindOperation) GetKindOperations() (*KindOperation, error) {

}

func (ko *KindOperation) UpdateKindOperation() (*KindOperation, error) {

}

func (ko *KindOperation) DeleteKindOperation() (*KindOperation, error) {

}

// Se refiere a si compras unos tenis, verduras, gymnasio, audifonos
type KindProduct struct {
	gorm.Model
	Name string
}

func (kp *KindProduct) CreateKindProduct() (*KindProduct, error) {

}

func (kp *KindProduct) GetKindProduct() (*KindProduct, error) {

}

func (kp *KindProduct) UpdateKindProduct() (*KindProduct, error) {

}

func (kp *KindProduct) DeleteKindProduct() (*KindProduct, error) {

}

func (kp *KindProduct) GetKindProducts() (*[]KindProduct, error) {

}

type Store struct {
	gorm.Model
	Name     string
	IsOnline bool
}

func (s *Store) CreateStore() (*Store, error) {

}

func (s *Store) GetStore() (*Store, error) {

}

func (s *Store) UpdateStore() (*Store, error) {

}

func (s *Store) DeleteStore() (*Store, error) {

}

func (s *Store) GetStores() (*[]Store, error) {

}

type MonthlyPayment struct {
	gorm.Model
	CreditCardId int64
	MonthlyPaid  float64
	Pending      float64
}

func (s *MonthlyPayment) CreateMonthlyPayment() (*MonthlyPayment, error) {

}

func (s *MonthlyPayment) GetMonthlyPayment() (*MonthlyPayment, error) {

}

func (s *MonthlyPayment) UpdateMonthlyPayment() (*MonthlyPayment, error) {

}

func (s *MonthlyPayment) DeleteMonthlyPayment() (*MonthlyPayment, error) {

}

func (s *MonthlyPayment) GetMonthlyPayments() (*[]MonthlyPayment, error) {

}

type Operation struct {
	gorm.Model
	UserId          int64
	CreditCardId    int64
	StoreId         int64
	KindOperationId int64
	KindProductId   int64
	Amount          int64
	Description     string
	Periods         int64
	PendingPeriods  int64
}
