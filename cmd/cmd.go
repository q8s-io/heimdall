package cmd

import (
	"flag"
	"log"

	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"github.com/q8s-io/heimdall/pkg/provider/analyzer"
	"github.com/q8s-io/heimdall/pkg/provider/process"
	"github.com/q8s-io/heimdall/pkg/provider/scanner"
	"github.com/q8s-io/heimdall/pkg/router"
)

var confPath = flag.String("conf", "./configs/pro.toml", "The path of config.")

// 服务类型
var serverType = flag.String("type", "scancenter", "The type of server.")

// scanner扫描时间，单位分钟
var scanTime = flag.Int("scanTime", 5, "The spend time of scanner.")

func Run() {
	flag.Parse()
	// init
	process.Init(*confPath)
	// app
	RunApp(*serverType)
}

func RunApp(serverType string) {
	switch serverType {
	case "scancenter":
		RunScanCenter()
	case "analyzer":
		RunAnalyzer(*scanTime)
	case "scanner-anchore":
		RunScannerAnchore(*scanTime)
	case "scanner-trivy":
		RunScannerTrivy(*scanTime)
	case "scanner-clair":
		RunScannerClair(*scanTime)
	case "tool":
		RunTool()
	default:
		log.Println(serverType)
	}
}

func RunScanCenter() {
	mysql.Init()
	redis.Init()
	kafka.InitSyncProducer()
	router.RunAPI()
}

func RunAnalyzer(scanTime int) {
	kafka.InitConsumer()
	go analyzer.JobAnalyzer(scanTime)
	analyzer.Signal()
}

func RunScannerAnchore(scanTime int) {
	kafka.InitConsumer()
	scanner.JobAnchore(scanTime)
}

func RunScannerTrivy(scanTime int) {
	kafka.InitConsumer()
	scanner.JobTrivy(scanTime)
}

func RunScannerClair(scanTime int) {
	kafka.InitConsumer()
	scanner.JobClair(scanTime)
}

func RunTool() {
	router.RunTool()
}
