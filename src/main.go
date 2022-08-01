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
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	config.StartConnection()
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

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
		v1.GET(routeName+"/nearby", routes.GetNearbyPlaces)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
