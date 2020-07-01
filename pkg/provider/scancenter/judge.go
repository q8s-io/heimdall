package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
	
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/persistence"
)

func JudgeTask(imageRequestInfo *entity.ImageRequestInfo) (interface{}, error) {
	// get data by ImageName & ImageDigest
	taskImageScanDataList := GetTaskImageScan(imageRequestInfo)
	// if data is empty, run scan center
	if len(*taskImageScanDataList) == 0 {
		return TaskImageScanRotaryCreate(imageRequestInfo)
	}
	taskImageScanData := (*taskImageScanDataList)[0]
	// if status is running, return data
	if taskImageScanData.TaskStatus == entity.StatusRunning {
		return TaskImageScanMerger(&taskImageScanData)
	}
	// if status is succeed
	if taskImageScanData.TaskStatus == entity.StatusSucceed {
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
	if currentStatus == entity.StatusSucceed {
		persistence.UpdateTaskImageScanStatus(taskID, entity.StatusSucceed)
		persistence.DeleteTask(taskID)
	}
}

func GetTaskCurrentStatus(taskID string) string {
	taskStatus := persistence.GetTaskStatus(taskID)
	for _, v := range taskStatus {
		if v != "succeed" {
			return v
		}
	}
	return entity.StatusSucceed
}
