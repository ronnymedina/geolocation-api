package routes

import (
	"log"
	"net/http"
	"ronnymedina/geolocation-api/src/helpers"
	"ronnymedina/geolocation-api/src/services"
	"ronnymedina/geolocation-api/src/validations"

	"github.com/gin-gonic/gin"
)

func CreatePlace(c *gin.Context) {
	log.SetPrefix("[CreatePlace] ")
	defer helpers.ResInternalServerErr(c)

	var place validations.CreatePlace
	if err := c.ShouldBindJSON(&place); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Place to Create:")
	log.Println(place)

	id := services.CreatePlace(place)
	helpers.ResSuccess(c, 201, gin.H{"data": id})
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
