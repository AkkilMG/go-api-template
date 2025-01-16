package main

import (
	"example/api-template/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
