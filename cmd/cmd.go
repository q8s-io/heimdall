package cmd

import (
	"flag"
	"log"

	"github.com/q8s-io/heimdall/pkg/domain/process"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/router"
)

var confPath = flag.String("conf", "./configs/pro.toml", "The path of config.")
var serverTpye = flag.String("type", "api", "The type of server.")

func Run() {
	flag.Parse()
	//init
	process.Init(*confPath)
	//app
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
	router.RunAPI()
}

func RunAnalyzer() {
	//kafka consumer

	//run pull、inspect、delete
	//timeout, write to redis, job id status is failed

	//write to mysql, image name、digest、layer

	//kafka producer, task id、job id、image name
}

func RunScannerAnchore() {
	//kafka consumer

	//run anchore
	//timeout, write to redis, job id status is failed

	//write to mysql, job data
	//write to redis, job status
}

func RunTool() {
	router.RunTool()
}
