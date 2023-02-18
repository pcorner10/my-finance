package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pcorner10/my-finance/api/controller"
	"github.com/pcorner10/my-finance/api/models"
	"github.com/pcorner10/my-finance/database"
	"github.com/pcorner10/my-finance/middleware"
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

	// Add CORS middleware

	// unable credentials

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	router.Use(cors.New(config))

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	privateRoutes := router.Group("/api")
	privateRoutes.Use(middleware.JWTAuthMiddleware())

	{
		shopping := privateRoutes.Group("/shopping")

		{
			creditCard := shopping.Group("/credit_card")
			creditCard.POST("/create", controller.CreateCreditCard)
			creditCard.GET("/getbyid/:id", controller.GetCreditCardsByUserId)

			kindOperation := shopping.Group("/kind_operation")
			kindOperation.POST("/create", controller.CreateKindOperation)
			kindOperation.GET("/get", controller.GetKindOperations)

			products := shopping.Group("/products")
			products.POST("/create", controller.CreateKindProduct)
			products.GET("/get", controller.GetKindProducts)

			store := shopping.Group("/store")
			store.POST("/create", controller.CreateStore)
			store.GET("/get", controller.GetStores)

			monthlyPayment := shopping.Group("/monthly_payments")
			monthlyPayment.POST("/create", controller.CreateMonthlyPayment)
			monthlyPayment.GET("/getbycreditcard", controller.GetMonthlyPaymentByCreditCard)

			operation := shopping.Group("/operation")
			operation.POST("/create", controller.CreateOperation)
			operation.GET("/lastoperations/:id", controller.GetLast5Operations)
		}
	}

	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
