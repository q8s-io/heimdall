package scancenter

import (
	"time"

	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
	"github.com/q8s-io/heimdall/pkg/models"
)

func CreateImageVulnInfo(imageRequestInfo *models.ImageRequestInfo) *models.ImageVulnInfo {
	imageVulnInfo := new(models.ImageVulnInfo)
	imageVulnInfo.TaskID = distribution.GetUUID()
	imageVulnInfo.TaskStatus = models.StatusRunning
	imageVulnInfo.ImageName = imageRequestInfo.ImageName
	imageVulnInfo.ImageDigest = imageRequestInfo.ImageDigest
	imageVulnInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return imageVulnInfo
}

func ConvertImageVulnData(imageVulnInfo *models.ImageVulnInfo, active int) *models.ImageVulnData {
	imageVulnData := new(models.ImageVulnData)
	imageVulnData.TaskID = imageVulnInfo.TaskID
	imageVulnData.TaskStatus = imageVulnInfo.TaskStatus
	imageVulnData.ImageName = imageVulnInfo.ImageName
	imageVulnData.ImageDigest = imageVulnInfo.ImageDigest
	imageVulnData.CreateTime = imageVulnInfo.CreateTime
	imageVulnData.Active = active
	return imageVulnData
}
