package routes

import "github.com/gin-gonic/gin"

func CreatePlace(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "POST",
	})
}

func UpdatePlace(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UPDATE",
	})
}

func FindPlace(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET",
	})
}

func DeletePlace(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DELETE",
	})
}
