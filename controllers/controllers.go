package controllers

import (
	"go-gin-api/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find((&users))

	c.JSON(http.StatusOK, gin.H{"data": users})
}
