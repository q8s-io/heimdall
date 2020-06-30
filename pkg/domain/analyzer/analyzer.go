package analyzer

import (
	"encoding/json"

	"github.com/Shopify/sarama"

	"github.com/q8s-io/heimdall/pkg/infrastructure/net"
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
)

func JobAnalyzer() {
	var queue chan *sarama.ConsumerMessage
	queue = make(chan *sarama.ConsumerMessage, 1000)

	// consumer msg from mq
	service.ConsumerJobImageAnalyzerMsg(queue)
	jobImageAnalyzerMsg := new(models.JobImageAnalyzerMsg)
	for msg := range queue {
		_ = json.Unmarshal(msg.Value, &jobImageAnalyzerMsg)

		// image analyzer
		imageName := jobImageAnalyzerMsg.ImageName
		digest, layers := ImageAnalyzer(imageName)
		jobImageAnalyzerInfo := ConvertJobImageAnalyzerInfoByMsg(jobImageAnalyzerMsg, digest, layers)

		// send data to scancenter
		requestJSON, _ := json.Marshal(jobImageAnalyzerInfo)
		_ = net.HTTPPUT(models.Config.ScanCenter.AnalyzerURL, string(requestJSON))
	}

	close(queue)
}
