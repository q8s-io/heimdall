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

func JobAnalyzer() {
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
		digest, layers := docker.ImageAnalyzer(imageName)
		jobImageAnalyzerInfo := convert.JobImageAnalyzerInfoByMsg(jobScannerMsg, digest, layers)
		if digest == nil || layers == nil {
			//// 镜像下载失败，直接标记失败。不触发其他引擎的任务。
			//// 删除analyzer表中任务。
			//xray.ErrMini(errors.New("download mirror failed"))
			//repository.DeleteJobImageAnalyzer(jobImageAnalyzerInfo.TaskID)
			//// 标记任务失败。
			//repository.UpdateTaskImageScanStatus(jobImageAnalyzerInfo.TaskID, model.StatusFailed)
			//// redis缓存清除
			//// 直接返回，不触发其他scanner
			jobImageAnalyzerInfo.JobStatus = model.StatusFailed
		}
		// send data to scancenter
		requestJSON, _ := json.Marshal(jobImageAnalyzerInfo)
		if jobImageAnalyzerInfo.JobStatus == model.StatusSucceed {
			log.Printf("analyzer process succeed %s", imageName)
		} else {
			log.Printf("analyzer process failed %s", imageName)
		}
		_ = net.HTTPPUT(model.Config.ScanCenter.AnalyzerURL, string(requestJSON))
	}
}
