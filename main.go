package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-gin-api/models"

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
	r.LoadHTMLGlob("./public/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	port := os.Getenv("PORT")
	fmt.Println("**** Starting server at port", os.Getenv("PORT"), "****")
	r.Run(":" + port)
}
