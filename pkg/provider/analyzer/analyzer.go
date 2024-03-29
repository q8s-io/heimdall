package analyzer

import (
	"encoding/json"
	"log"

	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/docker"
	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
	"github.com/q8s-io/heimdall/pkg/infrastructure/net"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func JobAnalyzer(scanTime int) {
	dockerConfig := model.Config.Docker
	docker.Init(dockerConfig.Host, dockerConfig.Version, nil, nil)

	// consumer msg from mq
	repository.ConsumerMsgJobImageAnalyzer()
	jobScannerMsg := new(model.JobScannerMsg)

	for msg := range kafka.Queue {
		log.Printf("consumer msg from kafka %s", msg)
		_ = json.Unmarshal(msg, &jobScannerMsg)
		// image analyzer
		imageName := jobScannerMsg.ImageName
		digest, layers, analyzerErr := docker.ImageAnalyzer(imageName, scanTime)
		jobImageAnalyzerInfo := convert.JobImageAnalyzerInfoByMsg(jobScannerMsg, digest, layers)

		if digest == nil || layers == nil {
			jobImageAnalyzerInfo.JobStatus = model.StatusFailed
			// 处理下载镜像超时
			if analyzerErr != nil && analyzerErr.Error() == "context deadline exceeded" {
				jobImageAnalyzerInfo.JobStatus = model.StatusTimeout
			}
		}

		// send data to scancenter
		requestJSON, _ := json.Marshal(jobImageAnalyzerInfo)
		log.Printf("analyzer process %s \t %s", imageName, jobImageAnalyzerInfo.JobStatus)
		_ = net.HTTPPUT(model.Config.ScanCenter.AnalyzerURL, string(requestJSON))
	}
}
