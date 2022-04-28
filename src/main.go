package main

import (
	"ronnymedina/geolocation-api/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routeName := "/places"

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// places

	r.POST(routeName, routes.CreatePlace)
	r.PUT(routeName+"/:id", routes.UpdatePlace)
	r.GET(routeName+"/:id", routes.FindPlace)
	r.DELETE(routeName+"/:id", routes.DeletePlace)

	r.Run() // listen and serve on 0.0.0.0:8080
}
