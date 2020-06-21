package router

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
	rImage "github.com/q8s-io/heimdall/pkg/router/image"
	rSystem "github.com/q8s-io/heimdall/pkg/router/system"
	rTool "github.com/q8s-io/heimdall/pkg/router/tool"
)

var requestInput io.Writer = os.Stdout

func Run(serverTpye string) {
	router := gin.New()

	router.Use(ginext.GinLogger())
	router.Use(ginext.Cors())
	router.Use(ginext.GinPanic(requestInput))

	rSystem.Routes(router)
	customRoutes(serverTpye, router)

	_ = router.Run(":12001")
}

func customRoutes(serverTpye string, router *gin.Engine) {
	switch serverTpye {
	case "tool":
		rTool.Routes(router)
	case "image":
		rImage.Routes(router)
	default:
		log.Println(serverTpye)
	}
}
