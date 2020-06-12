package cmd

import (
	"flag"

	"github.com/70data/heimdall/pkg/domain/process"
	"github.com/70data/heimdall/pkg/router"
)

var confPath = flag.String("conf", "./configs/pro.toml", "The conf path.")

func Run() {
	flag.Parse()
	// init
	process.Init(*confPath)
	// gin
	router.Run()
}
