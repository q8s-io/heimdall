package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
)

func CreateTaskImageScan(imageRequestInfo *models.ImageRequestInfo) (*models.TaskImageScanInfo, error) {
	// preper task
	taskImageScanInfo := CreateTaskImageScanInfo(imageRequestInfo)
	taskImageScanData := ConvertTaskImageScanData(taskImageScanInfo, 1)
	err := service.NewTaskImageScan(*taskImageScanData)
	if err != nil {
		return nil, err
	}
	return taskImageScanInfo, nil
}

func GetTaskImageScan(imageRequestInfo *models.ImageRequestInfo) *[]models.TaskImageScanData {
	taskImageScanData := service.GetTaskImageScan(*imageRequestInfo)
	return taskImageScanData
}

func UpdateTaskImageScanDigest(jobImageAnalyzerInfo *models.JobImageAnalyzerInfo) {
	taskImageScanData := ConvertTaskImageScanDataByAnalyzerInfo(jobImageAnalyzerInfo)
	service.UpdateTaskImageScanDigest(taskImageScanData.TaskID, taskImageScanData.ImageDigest)
}

func UpdateTaskImageScanActive(imageRequestInfo *models.ImageRequestInfo) {
	service.UpdateTaskImageScanActive(imageRequestInfo.ImageName, 0)
}
