package scancenter

import (
	"time"
	
	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
	"github.com/q8s-io/heimdall/pkg/models"
)

func ConvertPreperTask(imageInfoRequest *models.ImageInfoRequest) models.ImageVulnInfo {
	ImageVulnInfo := new(models.ImageVulnInfo)
	ImageVulnInfo.TaskID = distribution.GetUUID()
	ImageVulnInfo.TaskStatus = models.TaskStatusRunning
	ImageVulnInfo.ImageName = imageInfoRequest.ImageName
	ImageVulnInfo.ImageDigest = imageInfoRequest.ImageDigest
	ImageVulnInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return *ImageVulnInfo
}

func ConvertPreperJobAnalyzer(imageVulnInfo *models.ImageVulnInfo) models.AnalyzerJobInfo {
	AnalyzerJobInfo := new(models.AnalyzerJobInfo)
	AnalyzerJobInfo.TaskID = imageVulnInfo.TaskID
	AnalyzerJobInfo.JobID = distribution.GetUUID()
	AnalyzerJobInfo.JobStatus = models.JobStatusRunning
	AnalyzerJobInfo.ImageName = imageVulnInfo.ImageName
	AnalyzerJobInfo.ImageDigest = imageVulnInfo.ImageDigest
	AnalyzerJobInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return *AnalyzerJobInfo
}
