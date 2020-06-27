package cmd

import (
	"flag"
	"log"

	"github.com/q8s-io/heimdall/pkg/domain/analyzer"
	"github.com/q8s-io/heimdall/pkg/domain/process"
	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"github.com/q8s-io/heimdall/pkg/router"
)

var confPath = flag.String("conf", "./configs/pro.toml", "The path of config.")
var serverTpye = flag.String("type", "api", "The type of server.")

func Run() {
	flag.Parse()
	// init
	process.Init(*confPath)
	// app
	RunApp(*serverTpye)
}

func RunApp(serverTpye string) {
	switch serverTpye {
	case "scancenter":
		RunScanCenter()
	case "analyzer":
		RunAnalyzer()
	case "scanner-anchore":
		RunScannerAnchore()
	case "tool":
		RunTool()
	default:
		log.Println(serverTpye)
	}
}

func RunScanCenter() {
	mysql.Init()
	redis.Init()
	kafka.InitSyncProducer()
	router.RunAPI()
}

func RunAnalyzer() {
	kafka.InitConsumer("analyzer")
	analyzer.JobAnalyzer()
}

func RunScannerAnchore() {

}

func RunTool() {
	router.RunTool()
}
