package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
)

func JudgeTask(imageRequestInfo *models.ImageRequestInfo) (interface{}, error) {
	// get data by ImageName & ImageDigest
	taskImageScanDataList := GetTaskImageScan(imageRequestInfo)
	// if data is empty, run scan center
	if len(*taskImageScanDataList) == 0 {
		return TaskImageScanRotaryCreate(imageRequestInfo)
	}
	taskImageScanData := (*taskImageScanDataList)[0]
	// if status is running, return data
	if taskImageScanData.TaskStatus == models.StatusRunning {
		return TaskImageScanMerger(&taskImageScanData)
	}
	// if status is succeed
	if taskImageScanData.TaskStatus == models.StatusSucceed {
		// if ImageDigest is empty, run scan center
		if imageRequestInfo.ImageDigest == "" {
			return TaskImageScanMerger(&taskImageScanData)
			// if ImageDigest is db.ImageDigest, return data
		} else if imageRequestInfo.ImageDigest == taskImageScanData.ImageDigest {
			return TaskImageScanMerger(&taskImageScanData)
			// if ImageDigest not is db.ImageDigest, mark old data, run scan center
		} else {
			UpdateTaskImageScanActive(imageRequestInfo)
			return TaskImageScanRotaryCreate(imageRequestInfo)
		}
	}
	return nil, nil
}

func JudgeTaskRotary(taskID string) {
	// judge status
	currentStatus := GetTaskCurrentStatus(taskID)
	// mark task status
	if currentStatus == models.StatusSucceed {
		service.UpdateTaskImageScanStatus(taskID, models.StatusSucceed)
		service.DeleteTask(taskID)
	}
}

func GetTaskCurrentStatus(taskID string) string {
	taskStatus := service.GetTaskStatus(taskID)
	for _, v := range taskStatus {
		if v != "succeed" {
			return v
		}
	}
	return models.StatusSucceed
}
