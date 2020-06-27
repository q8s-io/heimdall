package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
)

func CreateTaskImageScan(imageRequestInfo *models.ImageRequestInfo) (interface{}, error) {
	// preper task data
	imageVulnInfo := ConvertTaskImageScan(imageRequestInfo)
	err := service.NewImageScan(*imageVulnInfo)
	if err != nil {
		return nil, err
	}
	// preper job analyzer
	jobAnalyzerInfo := ConvertJobImageAnalyzer(imageVulnInfo)
	_ = service.NewImageAnalyzer(*jobAnalyzerInfo)
	// mark job status
	service.SetImageAnalyzerStatus((*jobAnalyzerInfo).TaskID, models.StatusRunning)
	// send msg to mq
	jobAnalyzerMsg := ConvertImageAnalyzerMsg(jobAnalyzerInfo)
	service.SendImageAnalyzerMsg(jobAnalyzerMsg)
	return imageVulnInfo, nil
}
