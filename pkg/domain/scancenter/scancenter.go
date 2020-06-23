package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
	"github.com/q8s-io/heimdall/pkg/models"
)

// RunScanCenter
func RunScanCenter() {

	//kafka consumer

	//write to redis, analyzer job status
	//generate scanner job id

	//kafka producer

	//scan redis
	ScanRedis()
}

func PreperScanenter(imageInfoRequest *models.ImageInfoRequest) models.ImageVulnInfo {
	//preper data
	ImageVulnInfo := new(models.ImageVulnInfo)
	ImageVulnInfo.ImageName = imageInfoRequest.ImageName
	ImageVulnInfo.ImageDigest = imageInfoRequest.ImageDigest
	ImageVulnInfo.TaskID = distribution.GetUUID()
	ImageVulnInfo.TaskStatus = models.TaskStatusRunning
	ImageVulnInfo.AnchoreTaskID = distribution.GetUUID()
	//write to mysql
	//write to redis
	//write to kafka, trigger analyzer
	return *ImageVulnInfo
}
