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

	v1 := r.Group("/v1")
	{
		v1.POST(routeName, routes.CreatePlace)
		v1.PUT(routeName+"/:id", routes.UpdatePlace)
		v1.GET(routeName+"/:id", routes.FindPlace)
		v1.DELETE(routeName+"/:id", routes.DeletePlace)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
