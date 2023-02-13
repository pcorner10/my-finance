package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id       uint64 `gorm:"primaryKey"`
	Name     string
	LastName string
}
