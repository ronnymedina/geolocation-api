package middlewares

import (
	"fmt"
	"net/http"
	"ronnymedina/geolocation-api/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindPlaceOrReturnNotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.AbortWithStatus(http.StatusNotFound)

		var id int64
		id, _ = strconv.ParseInt(c.Param("id"), 10, 64)

		fmt.Println(services.FindPlace(id), "resultado")
		// before request
		c.Next()
	}
}
