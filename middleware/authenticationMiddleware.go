package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")
		// authHeader := c.GetHeader("Authorization")
		// if authHeader == "" {
		// 	c.JSON(401, gin.H{"error": "Missing token"})
		// 	c.Abort()
		// 	return
		// }
		
		tokenString, err := c.Cookie("access_token")
		log.Println("Password hashing error:", err)
        if err != nil {
           c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token missing"})
           c.Abort()
           return
        }
		// tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// if tokenString == authHeader{
		// 	c.JSON(401, gin.H{"error": "bearer token is required"})
		// 	c.Abort()
		// 	return
		// }
		token, err := jwt.ParseWithClaims(tokenString,&jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(JWT_SECRET_KEY), nil
		})


		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims,ok := token.Claims.(*jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("user", claims) // ✅ store in context

		c.Next()
	}
}
