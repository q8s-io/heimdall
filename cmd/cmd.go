package cmd

import (
	"flag"

	"github.com/q8s-io/heimdall/pkg/domain/process"
	"github.com/q8s-io/heimdall/pkg/router"
)

var confPath = flag.String("conf", "./configs/pro.toml", "The conf path.")

func Run() {
	flag.Parse()
	// init
	process.Init(*confPath)
	// gin
	router.Run()
}
