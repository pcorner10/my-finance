package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pcorner10/my-finance/controller"
	"github.com/pcorner10/my-finance/database"
	"github.com/pcorner10/my-finance/middleware"
	"github.com/pcorner10/my-finance/models"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(
		&models.User{},

		&models.Guest{}, &models.AgeRange{}, &models.Group{}, &models.Paymet{}, &models.Concept{}, &models.KindStatu{},
		&models.ContactInfo{}, &models.Funding{}, &models.ContactInfo{}, &models.AccumulatedMoney{},

		&models.CreditCard{}, &models.KindOperation{}, &models.KindProduct{}, &models.Store{}, &models.MonthlyPayment{},
		&models.Operation{},
	)
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
