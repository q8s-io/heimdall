package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/q8s-io/heimdall/docs"
	"github.com/q8s-io/heimdall/pkg/controller"
	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RouteCustom() *gin.Engine {
	router := gin.New()
	router.Use(ginext.GinLogger())
	router.Use(ginext.Cors())
	router.Use(ginext.BodyIntercept())
	router.Use(ginext.GinPanic())

	system := router.Group("/api/system")
	{
		// status
		system.GET("/status", controller.Status)
	}
	// swagger api
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
