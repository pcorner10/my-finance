package models

import (
	"html"
	"strings"

	"github.com/pcorner10/my-finance/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	LastName string
	UserName string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
}

func (u *User) GetUserByUsername() (*User, error) {

	err := database.Database.First(u, "username = ?", u.UserName).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.UserName = html.EscapeString(strings.TrimSpace(user.UserName))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("Entries").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
