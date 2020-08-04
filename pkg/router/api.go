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
		// update image analyzer data
		images.PUT("/analyzer/", controller.UpdateImageAnalyzerData)
		// update anchore data
		images.PUT("/anchore/", controller.UpdateAnchoreData)
		// update trivy data
		images.PUT("/trivy/", controller.UpdateTrivyData)
		// update clair data
		images.PUT("/clair/", controller.UpdateClairData)
	}
	_ = router.Run(":12001")
}
