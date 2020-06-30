package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
)

func CreateTaskImageScan(imageRequestInfo *models.ImageRequestInfo) (*models.ImageVulnInfo, error) {
	// preper task
	imageVulnInfo := CreateImageVulnInfo(imageRequestInfo)
	imageVulnData := ConvertImageVulnData(imageVulnInfo, 1)
	err := service.NewTaskImageScan(*imageVulnData)
	if err != nil {
		return nil, err
	}
	return imageVulnInfo, nil
}
