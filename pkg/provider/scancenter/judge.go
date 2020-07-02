package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func JudgeTask(imageRequestInfo *model.ImageRequestInfo) (interface{}, error) {
	// get data by ImageName & ImageDigest
	taskImageScanList := GetTaskImageScan(imageRequestInfo)
	// if data is empty, run scan center
	if len(*taskImageScanList) == 0 {
		return TaskImageScanRotaryCreate(imageRequestInfo)
	}
	taskImageScan := (*taskImageScanList)[0]
	// if status is running, return data
	if taskImageScan.TaskStatus == model.StatusRunning {
		return TaskImageScanMerger(&taskImageScan)
	}
	// if status is succeed
	if taskImageScan.TaskStatus == model.StatusSucceed {
		// if ImageDigest is empty, run scan center
		if imageRequestInfo.ImageDigest == "" {
			return TaskImageScanMerger(&taskImageScan)
			// if ImageDigest is db.ImageDigest, return data
		} else if imageRequestInfo.ImageDigest == taskImageScan.ImageDigest {
			return TaskImageScanMerger(&taskImageScan)
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
	if currentStatus == model.StatusSucceed {
		repository.UpdateTaskImageScanStatus(taskID, model.StatusSucceed)
		repository.DeleteTask(taskID)
	}
}

func GetTaskCurrentStatus(taskID string) string {
	taskStatus := repository.GetTaskStatus(taskID)
	for _, v := range taskStatus {
		if v != "succeed" {
			return v
		}
	}
	return model.StatusSucceed
}
