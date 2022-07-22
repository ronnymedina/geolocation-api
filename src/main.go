package main

import (
	"ronnymedina/geolocation-api/src/config"
	"ronnymedina/geolocation-api/src/middlewares"
	"ronnymedina/geolocation-api/src/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	config.StartConnection()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// places
	v1 := r.Group("/v1")
	{
		routeName := "/places"
		v1.POST(routeName, routes.CreatePlace)
		v1.PATCH(routeName+"/:id", middlewares.FindPlaceOrReturnNotFound(), routes.UpdatePlace)
		v1.GET(routeName+"/:id", middlewares.FindPlaceOrReturnNotFound(), routes.FindPlace)
		v1.DELETE(routeName+"/:id", middlewares.FindPlaceOrReturnNotFound(), routes.DeletePlace)
		v1.GET(routeName+"/nearby", routes.GetNearbyPlaces)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
