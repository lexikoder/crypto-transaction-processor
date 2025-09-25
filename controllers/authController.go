package controllers

import (
	"context"
	"crypto-transaction-processor/dto"
	service "crypto-transaction-processor/services"
	"crypto-transaction-processor/utils"
	"os"

	// "log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	Service *service.AuthService
	Validator *validator.Validate
}


func (authController *AuthController) ReqOtp() gin.HandlerFunc {
	return func(c *gin.Context) {
   var ctx, cancel = context.WithTimeout(context.Background(), 200*time.Second)
		defer cancel()
		reqotpdto := dto.BaseEmailDTO{}
		err := c.BindJSON(&reqotpdto)
		if err != nil {
			c.Error(utils.NewAppError(err.Error(), http.StatusBadRequest))
			return
		}

		ValidationError := authController.Validator.Struct(reqotpdto)

		if ValidationError != nil {
			c.Error(utils.ValidationAppError("validation error", http.StatusBadRequest,dto.FormatValidationErrors(ValidationError)))
			return
		}

		serviceErr := authController.Service.ReqOtpService(ctx,reqotpdto)
		if serviceErr != nil {
			c.Error(serviceErr) // handled by middleware
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "otp sent"})
		
	}
}

func (authController *AuthController) VerifyOtp() gin.HandlerFunc {
	return func(c *gin.Context) {
   var ctx, cancel = context.WithTimeout(context.Background(), 200*time.Second)
		defer cancel()
		verifyotpdto := dto.VerifyOtpDTO{}
		err := c.BindJSON(&verifyotpdto)
		if err != nil {
			c.Error(utils.NewAppError(err.Error(), http.StatusBadRequest))
			return
		}

		ValidationError := authController.Validator.Struct(verifyotpdto)

		if ValidationError != nil {
			c.Error(utils.ValidationAppError("validation error", http.StatusBadRequest,dto.FormatValidationErrors(ValidationError)))
			return
		}

		serviceErr := authController.Service.VerifyOtpService(ctx,verifyotpdto)
		if serviceErr != nil {
			c.Error(serviceErr) // handled by middleware
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "otp verified successfuly"})
		
	}
}


func (authController *AuthController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
	
		var ctx, cancel = context.WithTimeout(context.Background(), 200*time.Second)
		defer cancel()
		userdto := dto.SignupRequestDTO{}
		err := c.BindJSON(&userdto)
		if err != nil {
			c.Error(utils.NewAppError(err.Error(), http.StatusBadRequest))
			return
		}

		ValidationError := authController.Validator.Struct(userdto)

		if ValidationError != nil {
			c.Error(utils.ValidationAppError("validation error", http.StatusBadRequest,dto.FormatValidationErrors(ValidationError)))
			return
		}

		serviceErr := authController.Service.SignupService(ctx,userdto)
		if serviceErr != nil {
			c.Error(serviceErr) // handled by middleware
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "User created successfully"})
		// return

	}
}

func (authController *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		domain := os.Getenv("DOMAIN")
		env := os.Getenv("ENV")
		secure := true
		
		var ctx, cancel = context.WithTimeout(context.Background(), 200*time.Second)
		defer cancel()
		userdto := dto.LoginRequestDTO{}
		err := c.BindJSON(&userdto)
		if err != nil {
			c.Error(utils.NewAppError(err.Error(), http.StatusBadRequest))
			return
		}

		ValidationError := authController.Validator.Struct(userdto)

		if ValidationError != nil {
			c.Error(utils.ValidationAppError("validation error", http.StatusBadRequest,dto.FormatValidationErrors(ValidationError)))
			return
		}

		accessToken, refreshToken, serviceErr := authController.Service.LoginService(ctx,userdto)
		if serviceErr != nil {
			c.Error(serviceErr) // handled by middleware
			return
		} 

		if env == "production"{
           secure = false
		}

			// Secure cookie for access token
		c.SetSameSite(http.SameSiteStrictMode)
		c.SetCookie(
			"access_token",
			accessToken,
			service.AccessTokenExpiryTime * 60,              // 10 min
			"/",
			domain,    // production domain
			secure,             // secure = HTTPS only
			true,             // httpOnly
		)

		// Secure cookie for refresh token
		c.SetSameSite(http.SameSiteStrictMode)
		c.SetCookie(
			"refresh_token",
			refreshToken,
			service.RefreshTokenExpiryTime *60*60,       // 7 days
			"/",       // refresh-only endpoint
			domain,
			secure,
			true,    // this is httponly
		)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "logged in successfully",
		})
			
		// return

	}
	
}

func (authController *AuthController) RefreshAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func (authController *AuthController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func (authController *AuthController) Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "User created successfully"})
	}
}

func (authController *AuthController) Test2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "User created successfully"})
	}
}
