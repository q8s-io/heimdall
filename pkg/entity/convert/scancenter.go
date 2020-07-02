package convert

import (
	"time"

	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
)

func TaskImageScan(taskImageScanInfo *model.TaskImageScanInfo, active int) *entity.TaskImageScan {
	taskImageScan := new(entity.TaskImageScan)
	taskImageScan.TaskID = taskImageScanInfo.TaskID
	taskImageScan.TaskStatus = taskImageScanInfo.TaskStatus
	taskImageScan.ImageName = taskImageScanInfo.ImageName
	taskImageScan.ImageDigest = taskImageScanInfo.ImageDigest
	taskImageScan.CreateTime = taskImageScanInfo.CreateTime
	taskImageScan.Active = active
	return taskImageScan
}

func TaskImageScanByAnalyzerInfo(jobImageAnalyzerInfo *model.JobImageAnalyzerInfo) *entity.TaskImageScan {
	taskImageScan := new(entity.TaskImageScan)
	taskImageScan.TaskID = jobImageAnalyzerInfo.TaskID
	taskImageScan.ImageDigest = jobImageAnalyzerInfo.ImageDigest
	return taskImageScan
}

func TaskImageScanInfo(taskImageScan *entity.TaskImageScan) *model.TaskImageScanInfo {
	taskImageScanInfo := new(model.TaskImageScanInfo)
	taskImageScanInfo.TaskID = taskImageScan.TaskID
	taskImageScanInfo.TaskStatus = taskImageScan.TaskStatus
	taskImageScanInfo.ImageName = taskImageScan.ImageName
	taskImageScanInfo.ImageDigest = taskImageScan.ImageDigest
	taskImageScanInfo.CreateTime = taskImageScan.CreateTime
	return taskImageScanInfo
}

func TaskImageScanInfoByRequestInfo(imageRequestInfo *model.ImageRequestInfo) *model.TaskImageScanInfo {
	taskImageScanInfo := new(model.TaskImageScanInfo)
	taskImageScanInfo.TaskID = distribution.GetUUID()
	taskImageScanInfo.TaskStatus = model.StatusRunning
	taskImageScanInfo.ImageName = imageRequestInfo.ImageName
	taskImageScanInfo.ImageDigest = imageRequestInfo.ImageDigest
	taskImageScanInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return taskImageScanInfo
}

func ImageVulnByScanInfo(taskImageScanInfo *model.TaskImageScanInfo, vulnData []map[string]string) *model.ImageVulnInfo {
	if vulnData == nil {
		vulnData = make([]map[string]string, 0)
	}
	imageVulnInfo := new(model.ImageVulnInfo)
	imageVulnInfo.TaskID = taskImageScanInfo.TaskID
	imageVulnInfo.TaskStatus = taskImageScanInfo.TaskStatus
	imageVulnInfo.ImageName = taskImageScanInfo.ImageName
	imageVulnInfo.ImageDigest = taskImageScanInfo.ImageDigest
	imageVulnInfo.CreateTime = taskImageScanInfo.CreateTime
	imageVulnInfo.VulnData = vulnData
	return imageVulnInfo
}
