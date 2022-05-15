package routes

import (
	"log"
	"net/http"
	"ronnymedina/geolocation-api/src/helpers"
	"ronnymedina/geolocation-api/src/services"
	"ronnymedina/geolocation-api/src/validations"
	"strconv"

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
	log.SetPrefix("[UpdatePlace] ")
	defer helpers.ResInternalServerErr(c)

	var place validations.UpdatePlace
	var id int64
	if err := c.ShouldBindJSON(&place); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Place to Update with id:" + c.Param("id"))
	log.Println(place)

	id, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	services.UpdatePlace(id, place)

	helpers.ResSuccess(c, 200, gin.H{"data": "ok"})
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
