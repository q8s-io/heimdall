package scancenter

import (
	"time"

	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
	"github.com/q8s-io/heimdall/pkg/models"
)

func CreateTaskImageScanInfo(imageRequestInfo *models.ImageRequestInfo) *models.TaskImageScanInfo {
	taskImageScanInfo := new(models.TaskImageScanInfo)
	taskImageScanInfo.TaskID = distribution.GetUUID()
	taskImageScanInfo.TaskStatus = models.StatusRunning
	taskImageScanInfo.ImageName = imageRequestInfo.ImageName
	taskImageScanInfo.ImageDigest = imageRequestInfo.ImageDigest
	taskImageScanInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return taskImageScanInfo
}

func ConvertTaskImageScanData(taskImageScanInfo *models.TaskImageScanInfo, active int) *models.TaskImageScanData {
	taskImageScanData := new(models.TaskImageScanData)
	taskImageScanData.TaskID = taskImageScanInfo.TaskID
	taskImageScanData.TaskStatus = taskImageScanInfo.TaskStatus
	taskImageScanData.ImageName = taskImageScanInfo.ImageName
	taskImageScanData.ImageDigest = taskImageScanInfo.ImageDigest
	taskImageScanData.CreateTime = taskImageScanInfo.CreateTime
	taskImageScanData.Active = active
	return taskImageScanData
}

func ConvertTaskImageScanDataByAnalyzerInfo(jobImageAnalyzerInfo *models.JobImageAnalyzerInfo) *models.TaskImageScanData {
	taskImageScanData := new(models.TaskImageScanData)
	taskImageScanData.TaskID = jobImageAnalyzerInfo.TaskID
	taskImageScanData.ImageDigest = jobImageAnalyzerInfo.ImageDigest
	return taskImageScanData
}

func CreateImageVulnData(taskImageScanInfo *models.TaskImageScanInfo) *models.ImageVulnData {
	imageVulnData := new(models.ImageVulnData)
	imageVulnData.TaskID = taskImageScanInfo.TaskID
	imageVulnData.TaskStatus = taskImageScanInfo.TaskStatus
	imageVulnData.ImageName = taskImageScanInfo.ImageName
	imageVulnData.ImageDigest = taskImageScanInfo.ImageDigest
	imageVulnData.CreateTime = taskImageScanInfo.CreateTime
	imageVulnData.VulnData = make([]map[string]string, 0)
	return imageVulnData
}

func ConvertImageVulnData(taskImageScanData *models.TaskImageScanData, vulnData []map[string]string) *models.ImageVulnData {
	imageVulnData := new(models.ImageVulnData)
	imageVulnData.TaskID = taskImageScanData.TaskID
	imageVulnData.TaskStatus = taskImageScanData.TaskStatus
	imageVulnData.ImageName = taskImageScanData.ImageName
	imageVulnData.ImageDigest = taskImageScanData.ImageDigest
	imageVulnData.CreateTime = taskImageScanData.CreateTime
	imageVulnData.VulnData = vulnData
	return imageVulnData
}
