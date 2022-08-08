package routes

import (
	"urlShorter/controllers"

	"github.com/gin-gonic/gin"
)

type routerEngine struct {
	*gin.Engine
}

func Init() *routerEngine {
	return &routerEngine{gin.Default()}
}

func (router *routerEngine) UrlRoute() {
	router.GET("/:shortUrl", controllers.GetRedirect())
	router.POST("/getShortUrl", controllers.GetShortUrl())
}
