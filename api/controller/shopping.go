package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pcorner10/my-finance/api/models"
)

func CreateCreditCard(ctx *gin.Context) {
	var cc models.CreditCard

	if err := ctx.ShouldBindJSON(&cc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.CreateCreditCard()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"CreditCard": cc})
}

func GetCreditCardsByUserId(ctx *gin.Context) {

	var input models.CreditCard
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	listCreditCard, err := input.GetCreditCardsByUserId()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"CreditCards": err.Error()})
		return
	}

	ctx.JSON(http.StatusFound, gin.H{"user": listCreditCard})
}

func CreateKindOperation(ctx *gin.Context) {

	var input models.KindOperation
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := input.CreateKindOperation()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"KindOperation": input})
}

func CreateKindProduct(ctx *gin.Context) {

	var input models.KindProduct
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := input.CreateKindProduct()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"KindProduct": input})
}

func CreateStore(ctx *gin.Context) {

	var input models.Store
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := input.CreateStore()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Store": input})
}

func CreateMonthlyPayment(ctx *gin.Context) {

	var input models.MonthlyPayment
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := input.CreateMonthlyPayment()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"MonthlyPayment": input})
}

func CreateOperation(ctx *gin.Context) {

	var input models.Operation
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := input.CreateOperation()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Operation": input})
}
