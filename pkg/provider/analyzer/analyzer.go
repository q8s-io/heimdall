package analyzer

import (
	"encoding/json"

	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/docker"
	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
	"github.com/q8s-io/heimdall/pkg/infrastructure/net"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func JobAnalyzer() {
	// consumer msg from mq
	repository.ConsumerMsgJobImageAnalyzer()
	jobScannerMsg := new(model.JobScannerMsg)
	for msg := range kafka.Queue {
		_ = json.Unmarshal(msg, &jobScannerMsg)

		// image analyzer
		imageName := jobScannerMsg.ImageName
		digest, layers := docker.ImageAnalyzer(imageName)
		jobImageAnalyzerInfo := convert.JobImageAnalyzerInfoByMsg(jobScannerMsg, digest, layers)

		// send data to scancenter
		requestJSON, _ := json.Marshal(jobImageAnalyzerInfo)
		_ = net.HTTPPUT(model.Config.ScanCenter.AnalyzerURL, string(requestJSON))
	}
}