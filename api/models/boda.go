package models

import (
	"time"

	"gorm.io/gorm"
)

type Guest struct {
	gorm.Model
	Name        string
	GroupId     string
	Table       int64
	Seat        int64
	IsConfirmed bool
	Description string
}

type AgeRange struct {
	gorm.Model
	Range string
}

// Puede ser por familia, por escuela,
type Group struct {
	gorm.Model
	Name string
}

type Paymet struct {
	gorm.Model
	FundingId   int64
	Amount      float64
	Pending     float64
	Description float64
}

// Por ejemplo: WeddingPlaner, Cuando se contrató,
// cuanto va a cobrar, organización a que pertenece,
// fecha limite de pago, datos de contactó
type Concept struct {
	gorm.Model
	ContactId     int64
	Organization  string
	ContractDate  time.Time
	DeathLine     time.Time
	WhoContracted int64 // Pedro o Karla
	Status        int64
	Description   string
}

// Si ya se liquidó, si está en busqueda, reunión pendeinte,
type KindStatu struct {
	gorm.Model
	Name string
}

type ContactInfo struct {
	gorm.Model
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
	Name string
}

// Pendiente de una tabla para ir abondando los recursos
type AccumulatedMoney struct {
	gorm.Model
	FundingId   int64
	Amount      float64
	Accumulated float64
}
