package seed

import (
	"log"

	"github.com/pcorner10/my-finance/pkg/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Users{},

		&models.CreditCards{}, &models.KindOperations{}, &models.KindProducts{}, &models.Stores{},
		&models.MonthlyPayments{}, &models.Operations{},

		&models.Guests{}, &models.AgeRanges{}, &models.Groups{}, &models.Paymets{}, &models.Concepts{},
		&models.KindStatus{}, &models.ContactInfo{}, &models.Funding{}, &models.AccumulatedMoney{},
	)

	if err != nil {
		log.Fatalln(err)
		return
	}
}
