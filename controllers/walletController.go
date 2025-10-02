package controllers

import (
	"context"
	"crypto-transaction-processor/dto"
	service "crypto-transaction-processor/services"
	"crypto-transaction-processor/utils"
	
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// "github.com/gin-gonic/gin"
)

type WalletController struct {
	Service   *service.WalletService
	Validator *validator.Validate
}

func (walletController *WalletController) CreateWallet() gin.HandlerFunc {
	return func(c *gin.Context) {
        var ctx, cancel = context.WithTimeout(context.Background(), 200*time.Second)
		defer cancel()
      walletdto := dto.CreateWalletDTO{}
		err := c.BindJSON(&walletdto)
		if err != nil {
			c.Error(utils.NewAppError(err.Error(), http.StatusBadRequest))
			return
		}

		ValidationError := walletController.Validator.Struct(walletdto)

		if ValidationError != nil {
			c.Error(utils.ValidationAppError("validation error", http.StatusBadRequest,utils.FormatValidationErrors(ValidationError)))
			return
		}
        
		// reqotpdto := dto.BaseEmailDTO{}
	    apiresponsedto := dto.CreateWalletApiResponseDTO{}
		serviceErr := walletController.Service.CreateWalletService(ctx,walletdto,apiresponsedto)
		if serviceErr != nil {
			c.Error(serviceErr) // handled by middleware
			return
		}

       c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "wallet created successfully",
		})
	}
	
}


func (walletController *WalletController) DepositToken() gin.HandlerFunc {
	return func(c *gin.Context) {
        var ctx, cancel = context.WithTimeout(context.Background(), 200*time.Second)
		defer cancel()
      depositdto := dto.DepositTokenDTO{}
		err := c.BindJSON(&depositdto)
		if err != nil {
			c.Error(utils.NewAppError(err.Error(), http.StatusBadRequest))
			return
		}

		ValidationError := walletController.Validator.Struct(depositdto)

		if ValidationError != nil {
			c.Error(utils.ValidationAppError("validation error", http.StatusBadRequest,utils.FormatValidationErrors(ValidationError)))
			return
		}
        
		
		serviceErr := walletController.Service.DepositTokenService(ctx,depositdto)
		if serviceErr != nil {
			c.Error(serviceErr) // handled by middleware
			return
		}

       c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "deposited  successfully",
		})

	}
}


func (walletController *WalletController) WithdrawToken() gin.HandlerFunc {
	return func(c *gin.Context) {
        var ctx, cancel = context.WithTimeout(context.Background(), 200*time.Second)
		defer cancel()
      withdrawdto := dto.WithdrawTokenDTO{}
		err := c.BindJSON(&withdrawdto)
		if err != nil {
			c.Error(utils.NewAppError(err.Error(), http.StatusBadRequest))
			return
		}

		ValidationError := walletController.Validator.Struct(withdrawdto)

		if ValidationError != nil {
			c.Error(utils.ValidationAppError("validation error", http.StatusBadRequest,utils.FormatValidationErrors(ValidationError)))
			return
		}
        
		apiresponsedto := dto.TransfertokenApiResponseDTO
		serviceErr := walletController.Service.WithdrawTokenService(ctx,withdrawdto,apiresponsedto )
		if serviceErr != nil {
			c.Error(serviceErr) // handled by middleware
			return
		}

       c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "withdraw successful",
		})

	}
}

// func (walletController *WalletController) WithdrawToken() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 	}
// }

func (walletController *WalletController) TransferToken() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (walletController *WalletController) TransactionHistory() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}