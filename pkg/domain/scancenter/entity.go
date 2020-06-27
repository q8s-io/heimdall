package scancenter

import (
	"time"

	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
	"github.com/q8s-io/heimdall/pkg/models"
)

func ConvertTaskImageScan(imageRequestInfo *models.ImageRequestInfo) *models.ImageVulnInfo {
	imageVulnInfo := new(models.ImageVulnInfo)
	imageVulnInfo.TaskID = distribution.GetUUID()
	imageVulnInfo.TaskStatus = models.StatusRunning
	imageVulnInfo.ImageName = imageRequestInfo.ImageName
	imageVulnInfo.ImageDigest = imageRequestInfo.ImageDigest
	imageVulnInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return imageVulnInfo
}

func ConvertJobImageAnalyzer(imageVulnInfo *models.ImageVulnInfo) *models.JobAnalyzerInfo {
	jobAnalyzerInfo := new(models.JobAnalyzerInfo)
	jobAnalyzerInfo.TaskID = imageVulnInfo.TaskID
	jobAnalyzerInfo.JobID = distribution.GetUUID()
	jobAnalyzerInfo.JobStatus = models.StatusRunning
	jobAnalyzerInfo.ImageName = imageVulnInfo.ImageName
	jobAnalyzerInfo.ImageDigest = imageVulnInfo.ImageDigest
	jobAnalyzerInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return jobAnalyzerInfo
}

func ConvertImageAnalyzerMsg(jobAnalyzerInfo *models.JobAnalyzerInfo) *models.JobAnalyzerMsg {
	jobAnalyzerMsg := new(models.JobAnalyzerMsg)
	jobAnalyzerMsg.TaskID = jobAnalyzerInfo.TaskID
	jobAnalyzerMsg.JobID = jobAnalyzerInfo.JobID
	jobAnalyzerMsg.ImageName = jobAnalyzerInfo.ImageName
	return jobAnalyzerMsg
}
