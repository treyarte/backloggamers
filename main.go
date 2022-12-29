package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	gamesRouter := router.Group("/games")

	gamesRouter.GET("", )
	gamesRouter.GET("/:id", )
	gamesRouter.POST("/:id", )
	gamesRouter.PATCH("/:id", )
	gamesRouter.DELETE("/:id", )
}