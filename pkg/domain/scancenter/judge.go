package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
)

func JudgeTask(imageRequestInfo *models.ImageRequestInfo) (interface{}, error) {
	// get data by imageName & imageDigest

	// if data is empty, run scan center
	return TaskImageScanRotaryCreate(imageRequestInfo)

	// if status is running, return data

	// if status is done
	// if imageDigest is empty, return data

	// if imageDigest is db.imageDigest, return data

	// if imageDigest not is db.imageDigest, mark old data, run scan center
}

func JudgeTaskRotary(taskID string) {
	// judge status
	currentStatus := GetTaskCurrentStatus(taskID)

	// mark task status
	if currentStatus == models.StatusSucceed {
		service.UpdateTaskImageScan(taskID, models.StatusSucceed)
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
