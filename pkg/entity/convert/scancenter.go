package convert

import (
	"time"
	
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
	"github.com/q8s-io/heimdall/pkg/models"
)

func CreateTaskImageScanInfo(imageRequestInfo *entity.ImageRequestInfo) *entity.TaskImageScanInfo {
	taskImageScanInfo := new(entity.TaskImageScanInfo)
	taskImageScanInfo.TaskID = distribution.GetUUID()
	taskImageScanInfo.TaskStatus = entity.StatusRunning
	taskImageScanInfo.ImageName = imageRequestInfo.ImageName
	taskImageScanInfo.ImageDigest = imageRequestInfo.ImageDigest
	taskImageScanInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return taskImageScanInfo
}

func ConvertTaskImageScanData(taskImageScanInfo *entity.TaskImageScanInfo, active int) *entity.TaskImageScanData {
	taskImageScanData := new(entity.TaskImageScanData)
	taskImageScanData.TaskID = taskImageScanInfo.TaskID
	taskImageScanData.TaskStatus = taskImageScanInfo.TaskStatus
	taskImageScanData.ImageName = taskImageScanInfo.ImageName
	taskImageScanData.ImageDigest = taskImageScanInfo.ImageDigest
	taskImageScanData.CreateTime = taskImageScanInfo.CreateTime
	taskImageScanData.Active = active
	return taskImageScanData
}

func ConvertTaskImageScanDataByAnalyzerInfo(jobImageAnalyzerInfo *entity.JobImageAnalyzerInfo) *entity.TaskImageScanData {
	taskImageScanData := new(entity.TaskImageScanData)
	taskImageScanData.TaskID = jobImageAnalyzerInfo.TaskID
	taskImageScanData.ImageDigest = jobImageAnalyzerInfo.ImageDigest
	return taskImageScanData
}

func CreateImageVulnData(taskImageScanInfo *entity.TaskImageScanInfo) *entity.ImageVulnData {
	imageVulnData := new(entity.ImageVulnData)
	imageVulnData.TaskID = taskImageScanInfo.TaskID
	imageVulnData.TaskStatus = taskImageScanInfo.TaskStatus
	imageVulnData.ImageName = taskImageScanInfo.ImageName
	imageVulnData.ImageDigest = taskImageScanInfo.ImageDigest
	imageVulnData.CreateTime = taskImageScanInfo.CreateTime
	imageVulnData.VulnData = make([]map[string]string, 0)
	return imageVulnData
}

func ConvertImageVulnData(taskImageScanData *entity.TaskImageScanData, vulnData []map[string]string) *entity.ImageVulnData {
	imageVulnData := new(entity.ImageVulnData)
	imageVulnData.TaskID = taskImageScanData.TaskID
	imageVulnData.TaskStatus = taskImageScanData.TaskStatus
	imageVulnData.ImageName = taskImageScanData.ImageName
	imageVulnData.ImageDigest = taskImageScanData.ImageDigest
	imageVulnData.CreateTime = taskImageScanData.CreateTime
	imageVulnData.VulnData = vulnData
	return imageVulnData
}
