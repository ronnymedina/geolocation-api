package main

import (
	"ronnymedina/geolocation-api/src/config"
	"ronnymedina/geolocation-api/src/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	config.StartConnection()
	// rows, err := db.Query("SELECT name FROM places")

	// defer rows.Close()

	// for rows.Next() {
	// 	var name string

	// 	err = rows.Scan(&name)
	// 	helpers.CheckError(err)
	// 	fmt.Println(name)
	// }

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// places
	v1 := r.Group("/v1")
	{
		routeName := "/places"
		v1.POST(routeName, routes.CreatePlace)
		v1.PUT(routeName+"/:id", routes.UpdatePlace)
		v1.GET(routeName+"/:id", routes.FindPlace)
		v1.DELETE(routeName+"/:id", routes.DeletePlace)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
