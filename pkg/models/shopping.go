package models

import (
	"time"

	"gorm.io/gorm"
)

type CreditCards struct {
	gorm.Model
	Id              uint64 `gorm:"primaryKey"`
	LastNumbers     uint64
	Bank            string
	IsCredit        bool
	DatePayment     time.Time
	DateCutoff      time.Time
	CreditLimit     float64
	CreditAvailable float64
}

// Hace referencia al tipo de operación que es; despensa, osio, 
type KindOperations struct {
	gorm.Model
	Id   uint64 `gorm:"primaryKey"`
	Name int64
}

// Se refiere a si compras unos tenis, verduras, gymnasio, audifonos
type KindProducts struct {
	gorm.Model
	Id   uint64 `gorm:"primaryKey"`
	Name string
}

type Stores struct {
	gorm.Model
	Id       uint64 `gorm:"primaryKey"`
	Name     string
	IsOnline bool
}

type MonthlyPayments struct {
	gorm.Model
	Id           uint64 `gorm:"primaryKey"`
	CreditCardId int64
	MonthlyPaid  float64
	Pending      float64
}

type Operations struct {
	gorm.Model
	Id              uint64 `gorm:"primaryKey"`
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
