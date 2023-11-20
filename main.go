package main

import (
	"go-gin-api/auth"
	"go-gin-api/models"
	"go-gin-api/router"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	models.ConnectDatabase()

	auth, err := auth.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}
	rtr := router.New(auth)

	port := os.Getenv("PORT")
	if err := http.ListenAndServe("0.0.0.0:"+port, rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}

	r.Run(":" + port)
}
