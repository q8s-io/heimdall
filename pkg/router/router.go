package router

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
	rDefault "github.com/q8s-io/heimdall/pkg/router/default"
	rTools "github.com/q8s-io/heimdall/pkg/router/tools"
)

var requestInput io.Writer = os.Stdout

func Run(serverTpye string) {
	router := gin.New()

	router.Use(ginext.GinLogger())
	router.Use(ginext.Cors())
	router.Use(ginext.GinPanic(requestInput))

	rDefault.Routes(router)
	customRoutes(serverTpye, router)

	_ = router.Run(":12001")
}

func customRoutes(serverTpye string, router *gin.Engine) {
	switch serverTpye {
	case "tools":
		rTools.Routes(router)
	default:
		log.Println(serverTpye)
	}
}
