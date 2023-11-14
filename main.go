package main

import (
	"fmt"
	"go-gin-api/controllers"
	"go-gin-api/models"
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

	r.LoadHTMLGlob("./public/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/users", controllers.GetUsers)

	port := os.Getenv("PORT")
	fmt.Println("**** Starting server at port", os.Getenv("PORT"), "****")
	r.Run(":" + port)
}
