package router

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/controller"
	"github.com/q8s-io/heimdall/pkg/domain/analyzer"
	"github.com/q8s-io/heimdall/pkg/domain/scancenter"
	"github.com/q8s-io/heimdall/pkg/domain/scanner"
	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
)

func CustomRoutes() *gin.Engine {
	var requestInput io.Writer = os.Stdout
	router := gin.New()
	router.Use(ginext.GinLogger())
	router.Use(ginext.Cors())
	router.Use(ginext.GinPanic(requestInput))

	system := router.Group("/api/system")
	{
		//status
		system.GET("/status", controller.Status)
	}

	return router
}

func Run(serverTpye string) {
	switch serverTpye {
	case "api":
		RunAPI()
	case "scancenter":
		scancenter.RunScanCenter()
	case "analyzer":
		analyzer.RunAnalyzer()
	case "scanner":
		scanner.RunScanner()
	case "tool":
		RunTool()
	default:
		log.Println(serverTpye)
	}
}
