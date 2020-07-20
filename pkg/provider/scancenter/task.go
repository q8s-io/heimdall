package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func CreateTaskImageScan(imageRequestInfo *model.ImageRequestInfo) *model.TaskImageScanInfo {
	// prepare task
	taskImageScanInfo := convert.TaskImageScanInfoByRequestInfo(imageRequestInfo)
	taskImageScan := convert.TaskImageScan(taskImageScanInfo, 1)
	repository.NewTaskImageScan(*taskImageScan)
	return taskImageScanInfo
}

func GetTaskImageScan(imageRequestInfo *model.ImageRequestInfo) *[]entity.TaskImageScan {
	taskImageScanList := repository.GetTaskImageScan(*imageRequestInfo)
	return taskImageScanList
}

func UpdateTaskImageScanDigest(jobImageAnalyzerInfo *model.JobImageAnalyzerInfo) {
	taskImageScan := convert.TaskImageScanByAnalyzerInfo(jobImageAnalyzerInfo)
	repository.UpdateTaskImageScanDigest(taskImageScan.TaskID, taskImageScan.ImageDigest)
}

func UpdateTaskImageScanActive(imageRequestInfo *model.ImageRequestInfo) {
	repository.UpdateTaskImageScanActive(imageRequestInfo.ImageName, 0)
}
