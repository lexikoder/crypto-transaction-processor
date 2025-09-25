package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, exists := c.Get("user")
		if !exists {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, ok := userClaims.(*jwt.MapClaims)
		if !ok {
          c.JSON(401, gin.H{"error": "Invalid claims type"})
          c.Abort()
          return
         }
		userinfoMap := (*claims)["userinfo"].(map[string]any)
		
		if userinfoMap["Role"] != requiredRole {
			c.JSON(403, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
       
		c.Next()
	}
}
