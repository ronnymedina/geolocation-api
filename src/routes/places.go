package routes

import (
	"log"
	"net/http"
	"ronnymedina/geolocation-api/src/validations"

	"github.com/gin-gonic/gin"
)

func CreatePlace(c *gin.Context) {
	log.SetPrefix("[CreatePlace] ")

	var place validations.CreatePlace

	if err := c.ShouldBindJSON(&place); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Current data to create in place")
	log.Println(place)

	c.JSON(200, gin.H{
		"message": "POST",
	})
}

func UpdatePlace(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UPDATE",
	})
}

func FindPlace(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET",
	})
}

func DeletePlace(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DELETE",
	})
}
