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
		log.Fatal("could not migrate db")
	}
    var validate = validator.New()
	repo := &database.Repository{DB:db}
	AuthService := &service.AuthService{Db:repo}
    authController := &controllers.AuthController{Service:AuthService,Validator:validate}
    // walletController := WalletController{Repo:repo}
    
    gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorHandler())
	routes.AuthRoutes(router ,authController)

	// router.Use(middleware.Authentication())
	// routes.AuthRoutes(router)

	
    log.Println("server started")
	log.Fatal(router.Run(":" + port ))


}