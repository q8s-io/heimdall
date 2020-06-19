package cmd

import (
	"flag"

	"github.com/q8s-io/heimdall/pkg/domain/process"
	"github.com/q8s-io/heimdall/pkg/router"
)

var confPath = flag.String("conf", "./configs/pro.toml", "The path of config.")
var serverTpye = flag.String("type", "api", "The type of server.")

func Run() {
	flag.Parse()
	// init
	process.Init(*confPath)
	// gin
	router.Run(*serverTpye)
}
