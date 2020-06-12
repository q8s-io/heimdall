package router

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"

	ge "github.com/q8s-io/heimdall/pkg/infrastructure/gin-extender"
)

var requestInput io.Writer = os.Stdout

func Run() {
	router := gin.New()

	router.Use(ge.GinLogger())
	router.Use(ge.Cors())
	router.Use(ge.GinPanic(requestInput))

	statusRoutes(router)

	_ = router.Run(":12001")
}
