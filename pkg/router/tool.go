package router

import (
	"github.com/q8s-io/heimdall/pkg/controller"
)

func RunTool() {
	router := CustomRoutes()

	tools := router.Group("/api/tools")
	{
		//id
		tools.GET("/id", controller.GetID)
	}

	_ = router.Run(":12001")
}
