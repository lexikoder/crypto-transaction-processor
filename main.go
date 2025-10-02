package main

import (
	"crypto-transaction-processor/controllers"
	"crypto-transaction-processor/database"
	"crypto-transaction-processor/middleware"
	"crypto-transaction-processor/models"
	"crypto-transaction-processor/routes"
	service "crypto-transaction-processor/services"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)



func main(){

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
    
	db ,err := database.NewConnection()

	if err !=nil {
		log.Fatal("could not load the database ")
	}


	err = models.MigrateBooks(db)
	if err != nil{
		log.Fatal("could not migrate db:",err)
	}
    var validate = validator.New()
	repo := &database.Repository{DB:db}
	AuthService := &service.AuthService{Db:repo}
    authController := &controllers.AuthController{Service:AuthService,Validator:validate}
	WalletService := &service.WalletService{Db:repo}
    walletController := &controllers.WalletController{Service:WalletService,Validator:validate}
    // walletController := WalletController{Repo:repo}
    
    gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// 20 request in 10 muintes rate limiting  20/(60 *10)
	router.Use(middleware.RateLimitMiddleware(0.03, 20))
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://lll.com"}, // allow all origins
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: false,          // cannot be true when using "*"
        MaxAge:           12 * time.Hour,
    }))
	router.Use(func(c *gin.Context) {
    // remove any X-Powered-By
    c.Writer.Header().Del("X-Powered-By")
   
    c.Next()
})

	router.Use(middleware.ErrorHandler())
	routes.AuthRoutes(router ,authController)
	routes.WalletRoutes(router , walletController)

	// router.Use(middleware.Authentication())
	// routes.AuthRoutes(router)

	             
    log.Println("server started")
	log.Fatal(router.Run(":" + port ))


}