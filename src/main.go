package main

import (
	"ronnymedina/geolocation-api/src/config"
	"ronnymedina/geolocation-api/src/middlewares"
	"ronnymedina/geolocation-api/src/routes"

	docs "ronnymedina/geolocation-api/src/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// PingExample godoc
// @Schemes
// @Description do ping
// @Accept json
// @Produce json
// @Success 200
// @Router /ping [get]
func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @host      localhost:8080
// @BasePath  /v1
func main() {
	config.StartConnection()
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", ping)

	// places
	v1 := r.Group("/v1")
	{
		routeName := "/places"
		v1.POST(routeName, routes.CreatePlace)
		v1.PATCH(routeName+"/:id", middlewares.FindPlaceOrReturnNotFound(), routes.UpdatePlace)
		v1.GET(routeName+"/:id", middlewares.FindPlaceOrReturnNotFound(), routes.FindPlace)
		v1.DELETE(routeName+"/:id", middlewares.FindPlaceOrReturnNotFound(), routes.DeletePlace)
		v1.POST(routeName+"/nearby", routes.GetNearbyPlaces)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
