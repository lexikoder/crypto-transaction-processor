package middleware

import (
	"crypto-transaction-processor/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"github.com/jackc/pgx/v5/pgconn"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// Handle custom AppError
			// checking if error is an AppError
			appErr, isAppError := err.(*utils.AppError)
			if isAppError {
				c.JSON(appErr.StatusCode, gin.H{
					"success": false,
					"message": appErr.Message,
					
				})
				return
			}

			validationErr, isValidationError := err.(*utils.ValidationError)
			if isValidationError {
				c.JSON(validationErr.StatusCode, gin.H{
					"success": false,
					"message": validationErr.Message,
					"errors": validationErr.Errors,
				})
				return
			}

			 var pgErr *pgconn.PgError
    if errors.As(err, &pgErr) {
        if pgErr.Code == "23505" { // unique_violation
            c.JSON(http.StatusBadRequest, gin.H{
                "success": false,
                "message": "duplicate key value violates unique constraint",
            })
            return
        }
    }

			// Handle GORM not found error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"success": false,
					"message": "record not found", // 🟡 generic
				})
				return
			}

			// if errors.Is(err, jwt.ErrTokenExpired) {
			// 	c.JSON(http.StatusUnauthorized, gin.H{
			// 		"success": false,
			// 		"message": "Token expired",
			// 	})
			// 	return
			// }

			// if errors.Is(err, jwt.ErrTokenInvalidClaims) {
			// 	c.JSON(http.StatusUnauthorized, gin.H{
			// 		"success": false,
			// 		"message": "invalid token",
			// 	})
			// 	return
			// }

			// if errors.Is(err, jwt.ErrSignatureInvalid) {
			// 	c.JSON(http.StatusUnauthorized, gin.H{
			// 		"success": false,
			// 		"message": "Invalid token signature",
			// 	})
			// 	return
			// }

			// Fallback
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "something went wrong",
			})
		}
	}
}
