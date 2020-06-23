package router

import (
	"github.com/q8s-io/heimdall/pkg/controller"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
)

func RunAPI() {
	go mysql.Init()
	
	router := CustomRoutes()

	images := router.Group("/api/images")
	{
		//get vuln result
		images.POST("/vuln/", controller.GetImageVulnData)
	}

	scan := router.Group("/api/scan")
	{
		//create scan task
		scan.POST("/task/", controller.CreateScanTask)
		//get scan task data
		scan.GET("/task/:taskid", controller.GetScanTaskData)
	}

	_ = router.Run(":12001")
}
