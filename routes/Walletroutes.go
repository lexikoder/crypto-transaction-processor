package routes

import (
	"crypto-transaction-processor/controllers"
	"crypto-transaction-processor/middleware"
	"github.com/gin-gonic/gin"
)



func WalletRoutes(app *gin.Engine,walletcontroller *controllers.WalletController){
	
	protected := app.Group("/api")
	protected.Use(middleware.AuthMiddleware()) 
	{
	protected.POST("/wallet/createwallet",walletcontroller.CreateWallet())
	protected.POST("/wallet/deposittoken",walletcontroller.DepositToken())
	protected.POST("/wallet/withdrawtoken",walletcontroller.WithdrawToken())
	// protected.POST("/wallet/transfertoken",walletcontroller.TransferToken())
	protected.POST("/wallet/transactionhistory",walletcontroller.TransactionHistory())
	
	}

	  // Role-based auth
    // protectedAdmin := protected.Group("")
    // protectedAdmin.Use(middleware.AuthorizeRole("USER"))
    // {
    //     protectedAdmin.GET("/wallet/test2", authcontroller.Test2())
    // }
	// }
    
}