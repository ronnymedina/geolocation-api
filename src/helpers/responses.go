package helpers

import (
	"github.com/gin-gonic/gin"
)

func ResSuccess(c *gin.Context, statusCode int, data gin.H) {
	c.JSON(statusCode, data)
}

func ResInternalServerErr(c *gin.Context) {
	if r := recover(); r != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
	}
}

func ResNotFound(c *gin.Context) {
	if r := recover(); r != nil {
		c.JSON(404, gin.H{"error": "Not found resource"})
	}
}
