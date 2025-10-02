package routes

import (
	"crypto-transaction-processor/controllers"
	"crypto-transaction-processor/middleware"
	"github.com/gin-gonic/gin"
)



func AuthRoutes(app *gin.Engine,authcontroller *controllers.AuthController){
	public := app.Group("/api")
	{
	public.POST("/auth/signup",authcontroller.SignUp())
	public.POST("/auth/login",authcontroller.Login())
	public.POST("/auth/reqotp",authcontroller.ReqOtp())
	public.POST("/auth/verifyotp",authcontroller.VerifyOtp())
	public.POST("/auth/refreshtoken",authcontroller.RefreshAccessToken())
	}

	protected := app.Group("/api")
	protected.Use(middleware.AuthMiddleware()) 
	{
	protected.GET("/wallet/test",authcontroller.Test())
	protected.POST("/auth/logout",authcontroller.Logout())


	  // Role-based auth
    protectedAdmin := protected.Group("")
    protectedAdmin.Use(middleware.AuthorizeRole("USER"))
    {
        protectedAdmin.GET("/wallet/test2", authcontroller.Test2())
    }
	}
    
}