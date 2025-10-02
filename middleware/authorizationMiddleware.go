package middleware

import (
	"crypto-transaction-processor/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, exists := c.Get("user")
		if !exists {
			c.Error(utils.NewAppError("Unauthorized", 401))
			// c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, ok := userClaims.(*jwt.MapClaims)
		if !ok {
		  c.Error(utils.NewAppError("Invalid claims type", 401))
        //   c.JSON(401, gin.H{"error": "Invalid claims type"})
          c.Abort()
          return
         }
		userinfoMap := (*claims)["userinfo"].(map[string]any)
		
		if userinfoMap["Role"] != requiredRole {
			c.Error(utils.NewAppError("Forbidden", 403))
			// c.JSON(403, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
       
		c.Next()
	}
}
