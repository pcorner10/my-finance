package models

import (
	"time"

	"gorm.io/gorm"
)

type Guests struct {
	gorm.Model
	Id          uint64 `gorm:"primaryKey"`
	Name        string
	GroupId     string
	Table       int64
	Seat        int64
	IsConfirmed bool
	Description string
}

type AgeRanges struct {
	gorm.Model
	Id    uint64 `gorm:"primaryKey"`
	Range string
}

// Puede ser por familia, por escuela,
type Groups struct {
	gorm.Model
	Id   uint64 `gorm:"primaryKey"`
	Name string
}

type Paymets struct {
	gorm.Model
	Id          uint64 `gorm:"primaryKey"`
	FundingId   int64
	Amount      float64
	Pending     float64
	Description float64
}

// Por ejemplo: WeddingPlaner, Cuando se contrató,
// cuanto va a cobrar, organización a que pertenece,
// fecha limite de pago, datos de contactó
type Concepts struct {
	gorm.Model
	Id            uint64 `gorm:"primaryKey"`
	ContactId     int64
	Organization  string
	ContractDate  time.Time
	DeathLine     time.Time
	WhoContracted int64 // Pedro o Karla
	Status        int64
	Description   string
}

// Si ya se liquidó, si está en busqueda, reunión pendeinte,
type KindStatus struct {
	gorm.Model
	Id   uint64 `gorm:"primaryKey"`
	Name string
}

type ContactInfo struct {
	gorm.Model
	Id        uint64 `gorm:"primaryKey"`
	Name      string
	Cellphone string
	Email     string
	Facebook  string
	Address   string
	Instagram string
}

// Si es pronaces o CONACYT o Karla o Pedro o Oliver
type Funding struct {
	gorm.Model
	Id   uint64 `gorm:"primaryKey"`
	Name string
}

// Pendiente de una tabla para ir abondando los recursos
type AccumulatedMoney struct {
	gorm.Model
	Id          uint64 `gorm:"primaryKey"`
	FundingId   int64
	Amount      float64
	Accumulated float64
}
