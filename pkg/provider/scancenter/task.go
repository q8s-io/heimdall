package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
	
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/persistence"
)

func CreateTaskImageScan(imageRequestInfo *entity.ImageRequestInfo) (*entity.TaskImageScanInfo, error) {
	// preper task
	taskImageScanInfo := CreateTaskImageScanInfo(imageRequestInfo)
	taskImageScanData := ConvertTaskImageScanData(taskImageScanInfo, 1)
	err := persistence.NewTaskImageScan(*taskImageScanData)
	if err != nil {
		return nil, err
	}
	return taskImageScanInfo, nil
}

func GetTaskImageScan(imageRequestInfo *entity.ImageRequestInfo) *[]entity.TaskImageScanData {
	taskImageScanData := persistence.GetTaskImageScan(*imageRequestInfo)
	return taskImageScanData
}

func UpdateTaskImageScanDigest(jobImageAnalyzerInfo *entity.JobImageAnalyzerInfo) {
	taskImageScanData := ConvertTaskImageScanDataByAnalyzerInfo(jobImageAnalyzerInfo)
	persistence.UpdateTaskImageScanDigest(taskImageScanData.TaskID, taskImageScanData.ImageDigest)
}

func UpdateTaskImageScanActive(imageRequestInfo *entity.ImageRequestInfo) {
	persistence.UpdateTaskImageScanActive(imageRequestInfo.ImageName, 0)
}
