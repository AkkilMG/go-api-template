package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BasicMiddleware(c *gin.Context) {
	ipAddress := c.ClientIP()
	fmt.Println("IP Address: ", ipAddress)
	c.Next()
}
