package router

import (
	"github.com/q8s-io/heimdall/pkg/controller"
)

func RunAPI() {
	router := RouteCustom()

	images := router.Group("/api/images")
	{
		// get image vuln result
		images.POST("/vuln/", controller.GetImageVulnData)
	}

	_ = router.Run(":12001")
}
