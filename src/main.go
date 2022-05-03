package main

import (
	"database/sql"
	"fmt"
	"os"
	"ronnymedina/geolocation-api/src/helpers"
	"ronnymedina/geolocation-api/src/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	PG_HOST := os.Getenv("POSTGRE_HOST")
	PG_USER := os.Getenv("POSTGRE_USER")
	PG_PASS := os.Getenv("POSTGRE_PASS")
	PG_DB := os.Getenv("POSTGRE_DB")
	URL_CONNECT := "postgres://" + PG_USER + ":" + PG_PASS + "@" + PG_HOST + "/" + PG_DB + "?sslmode=disable"
	db, err := sql.Open("postgres", URL_CONNECT)

	helpers.CheckError(err)
	defer db.Close()

	err = db.Ping()
	helpers.CheckError(err)

	rows, err := db.Query("SELECT name FROM places")

	helpers.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var name string

		err = rows.Scan(&name)
		helpers.CheckError(err)
		fmt.Println(name)
	}

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
