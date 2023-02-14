package server

import (
	"github.com/pcorner10/my-finance/pkg/api/users"
	"gorm.io/gorm"
)

func bindUser(dbHandler *gorm.DB) users.Handler {
	repository := users.NewRepository(dbHandler)
	service := users.NewService(repository)

	return users.NewHandler(service)
}
