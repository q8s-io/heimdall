package analyzer

import (
	"encoding/json"

	"github.com/Shopify/sarama"

	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/docker"
	"github.com/q8s-io/heimdall/pkg/infrastructure/net"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func JobAnalyzer() {
	var queue chan *sarama.ConsumerMessage
	queue = make(chan *sarama.ConsumerMessage, 1000)

	// consumer msg from mq
	repository.ConsumerMsgJobImageAnalyzer(queue)
	jobScannerMsg := new(model.JobScannerMsg)
	for msg := range queue {
		_ = json.Unmarshal(msg.Value, &jobScannerMsg)

		// image analyzer
		imageName := jobScannerMsg.ImageName
		digest, layers := docker.ImageAnalyzer(imageName)
		jobImageAnalyzerInfo := convert.JobImageAnalyzerInfoByMsg(jobScannerMsg, digest, layers)

		// send data to scancenter
		requestJSON, _ := json.Marshal(jobImageAnalyzerInfo)
		_ = net.HTTPPUT(model.Config.ScanCenter.AnalyzerURL, string(requestJSON))
	}

	close(queue)
}
