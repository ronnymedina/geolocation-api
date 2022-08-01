package routes

import (
	"log"
	"net/http"
	"ronnymedina/geolocation-api/src/helpers"
	"ronnymedina/geolocation-api/src/models"
	"ronnymedina/geolocation-api/src/services"
	"ronnymedina/geolocation-api/src/validations"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// PingExample godoc
// @Schemes
// @Description Create a new place
// @Accept json
// @Produce json
// @Router /places [post]
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

	var place *models.Place
	var updatePlace validations.UpdatePlace

	if err := c.ShouldBindJSON(&updatePlace); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(updatePlace)
	place = c.MustGet("place").(*models.Place)
	services.UpdatePlace(place.Id, updatePlace)
	helpers.ResSuccess(c, 200, gin.H{"data": "ok"})
}

func FindPlace(c *gin.Context) {
	place := c.MustGet("place").(*models.Place)

	c.JSON(200, gin.H{"data": place})
}

func DeletePlace(c *gin.Context) {
	place := c.MustGet("place").(*models.Place)
	services.DeletePlace(place.Id)

	c.JSON(200, gin.H{"data": "ok"})
}

func GetNearbyPlaces(c *gin.Context) {
	var params validations.NearbyPlaces

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if params.Page == 1 {
		params.Page = 0
	}

	if params.Page > 0 {
		params.Page = params.Limit * (params.Page - 1)
	}

	rows := services.GetNearbyPlaces(params)

	c.JSON(200, gin.H{"data": rows})
}
