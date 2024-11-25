package controller

import (
	"net/http"
	"teknikal-test/delivery/middleware"
	"teknikal-test/entity/request"
	"teknikal-test/usecase"
	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentUc usecase.PaymentUsecase
	rg        *gin.RouterGroup
	paymentMd middleware.AuthMiddleware
}

func (p *PaymentController) Payment(ctx *gin.Context) {
	var transactionRequest request.TransactionRequest
	if err := ctx.ShouldBindJSON(&transactionRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := p.paymentUc.Payment(transactionRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (p *PaymentController) Route() {
	p.rg.POST("/payment", p.paymentMd.RequireToken(), p.Payment)
}

func NewPaymentController(paymentUc usecase.PaymentUsecase, rg *gin.RouterGroup, paymentMd middleware.AuthMiddleware) *PaymentController {
	return &PaymentController{paymentUc, rg, paymentMd}
}