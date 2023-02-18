package controller

import (
	"fmt"
	"net/http"
	"strconv"

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
	userID := ctx.Param("id")
	var input models.CreditCard

	// userID is a string, so we need to convert it to an int64
	userIDInt, _ := strconv.ParseInt(userID, 10, 64)

	input.UserId = userIDInt
	listCreditCard, err := input.GetCreditCardsByUserId()
	if err != nil {
		fmt.Println(err, "error2")
		ctx.JSON(http.StatusBadRequest, gin.H{"CreditCards": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"listCreditCards": listCreditCard})
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

func GetKindOperations(ctx *gin.Context) {

	var input models.KindOperation

	listKindOperation, err := input.GetKindOperations()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"KindOperations": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"KindOperations": listKindOperation})
}

func GetLast5Operations(ctx *gin.Context) {
	userID := ctx.Param("id")
	var input models.Operation
	// userID is a string, so we need to convert it to an int64
	userIDInt, _ := strconv.ParseInt(userID, 10, 64)

	input.UserId = userIDInt
	
	listOperation, err := input.GetLast5Operations()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Last5Operations": listOperation})
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

func GetKindProducts(ctx *gin.Context) {

	var input models.KindProduct
	listKindProduct, err := input.GetKindProducts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"KindProducts": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"KindProducts": listKindProduct})
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

func GetStores(ctx *gin.Context) {

	var input models.Store

	listStore, err := input.GetStores()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Stores": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Stores": listStore})
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

func GetMonthlyPaymentByCreditCard(ctx *gin.Context) {

	var input models.MonthlyPayment
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	listMonthlyPayment, err := input.GetMonthlyPaymentByCreditCard()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"MonthlyPayments": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"MonthlyPayments": listMonthlyPayment})
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
