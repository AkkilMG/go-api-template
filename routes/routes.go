package routes

import (
	"example/api-template/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", handlers.RedirectDetails)
	router.GET("/basics", handlers.BasicMiddleware, handlers.GetBasics)
	router.POST("/create", handlers.CreateBasic)
	router.GET("/:id", handlers.BasicInfo)
	router.GET("/search", handlers.SearchDetail)
	router.Static("/static", "./public")
}
