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

// @Schemes
// @Description Create a new place
// @Param place body validations.CreatePlace true "Place"
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

// @Schemes
// @Description Update place
// @Param id path int true  "Place ID"
// @Param place body validations.UpdatePlace true "Place"
// @Accept json
// @Produce json
// @Router /places/{id} [patch]
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

// @Schemes
// @Description Find place
// @Param id path int true  "Place ID"
// @Produce json
// @Router /places/{id} [get]
func FindPlace(c *gin.Context) {
	place := c.MustGet("place").(*models.Place)

	c.JSON(200, gin.H{"data": place})
}

// @Schemes
// @Description delete a place
// @Produce json
// @Param id path int true  "Place ID"
// @Router /places/{id} [delete]
func DeletePlace(c *gin.Context) {
	place := c.MustGet("place").(*models.Place)
	services.DeletePlace(place.Id)

	c.JSON(200, gin.H{"data": "ok"})
}

// @Schemes
// @Description Get nearby place from current location
// @Param place body validations.NearbyPlaces true "NearbyPlaces"
// @Accept json
// @Produce json
// @Router /places/nearby [post]
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
