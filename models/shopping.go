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

func (cc *CreditCard) GetCreditCardByBank() (*CreditCard, error) {
	err := database.Database.Where("bank = ? AND is_credit = true").First(cc).Error
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (cc *CreditCard) UpdateCreditCard() (*CreditCard, error) {
	err := database.Database.Model(cc).Updates(cc).Error
	if err != nil {
		return nil, err
	}
	return cc, nil

}

func (cc *CreditCard) DeleteCreditCard() (*CreditCard, error) {
	if err != nil {
		return nil, err
	}
	return cc, nil

}

func (cc *CreditCard) GetCreditCards() (*[]CreditCard, error) {
	if err != nil {
		return nil, err
	}
	return cc, nil
}

// Hace referencia al tipo de operación que es; despensa, osio,
type KindOperation struct {
	gorm.Model
	Name int64
}

func (ko *KindOperation) CreateKindOperation() (*KindOperation, error) {
	err := database.Database.Create(ko).Error
	if err != nil {
		return nil, err
	}
	return ko, nil
}

func (ko *KindOperation) GetKindOperation() (*KindOperation, error) {
	if err != nil {
		return nil, err
	}
	return ko, nil

}

func (ko *KindOperation) GetKindOperations() (*[]KindOperation, error) {
	if err != nil {
		return nil, err
	}
	return ko, nil

}

func (ko *KindOperation) UpdateKindOperation() (*KindOperation, error) {
	if err != nil {
		return nil, err
	}
	return ko, nil

}

func (ko *KindOperation) DeleteKindOperation() (*KindOperation, error) {
	if err != nil {
		return nil, err
	}
	return ko, nil

}

// Se refiere a si compras unos tenis, verduras, gymnasio, audifonos
type KindProduct struct {
	gorm.Model
	Name string
}

func (kp *KindProduct) CreateKindProduct() (*KindProduct, error) {
	err := database.Database.Create(kp).Error
	if err != nil {
		return nil, err
	}
	return kp, nil
}

func (kp *KindProduct) GetKindProduct() (*KindProduct, error) {
	if err != nil {
		return nil, err
	}
	return kp, nil

}

func (kp *KindProduct) UpdateKindProduct() (*KindProduct, error) {
	if err != nil {
		return nil, err
	}
	return kp, nil

}

func (kp *KindProduct) DeleteKindProduct() (*KindProduct, error) {
	if err != nil {
		return nil, err
	}
	return kp, nil

}

func (kp *KindProduct) GetKindProducts() (*[]KindProduct, error) {
	if err != nil {
		return nil, err
	}
	return kp, nil

}

type Store struct {
	gorm.Model
	Name     string
	IsOnline bool
}

func (s *Store) CreateStore() (*Store, error) {
	err := database.Database.Create(s).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Store) GetStore() (*Store, error) {
	if err != nil {
		return nil, err
	}
	return s, nil

}

func (s *Store) UpdateStore() (*Store, error) {
	if err != nil {
		return nil, err
	}
	return s, nil

}

func (s *Store) DeleteStore() (*Store, error) {
	if err != nil {
		return nil, err
	}
	return s, nil

}

func (s *Store) GetStores() (*[]Store, error) {

}

type MonthlyPayment struct {
	gorm.Model
	CreditCardId int64
	MonthlyPaid  float64
	Pending      float64
}

func (mp *MonthlyPayment) CreateMonthlyPayment() (*MonthlyPayment, error) {
	err := database.Database.Create(mp).Error
	if err != nil {
		return nil, err
	}
	return mp, nil
}

func (mp *MonthlyPayment) GetMonthlyPayment() (*MonthlyPayment, error) {
	if err != nil {
		return nil, err
	}
	return mp, nil

}

func (mp *MonthlyPayment) UpdateMonthlyPayment() (*MonthlyPayment, error) {
	if err != nil {
		return nil, err
	}
	return mp, nil

}

func (mp *MonthlyPayment) DeleteMonthlyPayment() (*MonthlyPayment, error) {
	if err != nil {
		return nil, err
	}
	return mp, nil

}

func (mp *MonthlyPayment) GetMonthlyPayments() (*[]MonthlyPayment, error) {
	if err != nil {
		return nil, err
	}
	return mp, nil

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

func (o *Operation) CreateOperation() (*Operation, error) {
	err := database.Database.Create(o).Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *Operation) GetOperation() (*Operation, error) {
	if err != nil {
		return nil, err
	}
	return o, nil

}

func (o *Operation) UpdateOperation() (*Operation, error) {
	if err != nil {
		return nil, err
	}
	return o, nil

}

func (o *Operation) DeleteOperation() (*Operation, error) {
	if err != nil {
		return nil, err
	}
	return o, nil

}

func (o *Operation) GetOperations() (*[]Operation, error) {
	if err != nil {
		return nil, err
	}
	return o, nil

}
