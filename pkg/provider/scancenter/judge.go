package scancenter

import (
	"strings"

	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func JudgeTask(imageRequestInfo *model.ImageRequestInfo) (interface{}, error) {
	// get data by ImageName
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

		// if Image tag is latest or empty, run scan center
		imageSlice := strings.Split(imageRequestInfo.ImageName, ":")
		imageTag := imageSlice[len(imageSlice)-1]
		if imageTag == "latest" || len(imageSlice) == 1 {
			UpdateTaskImageScanActive(imageRequestInfo)
			return TaskImageScanRotaryCreate(imageRequestInfo)

			// if ImageDigest is db.ImageDigest, return data
		} else if imageRequestInfo.ImageDigest == taskImageScan.ImageDigest {
			return TaskImageScanMerger(&taskImageScan)

			// if ImageDigest not is db.ImageDigest, mark old data, run scan center
		} else {
			UpdateTaskImageScanActive(imageRequestInfo)
			return TaskImageScanRotaryCreate(imageRequestInfo)

		}
	}
	return taskImageScanList, nil
}

func JudgeTaskRotary(taskID string) {
	// judge status
	currentStatus := GetTaskCurrentStatus(taskID)
	// mark task status
	if currentStatus != model.StatusRunning {
		repository.UpdateTaskImageScanStatus(taskID, currentStatus)
		repository.DeleteTask(taskID)
	}
}

func GetTaskCurrentStatus(taskID string) string {
	taskStatus := repository.GetTaskStatus(taskID)
	// count success num
	succeedNum := 0

	for _, v := range taskStatus {
		// 若还有在运行的引擎，直接返回运行态
		if v == model.StatusRunning {
			return model.StatusRunning
		}
		if v == model.StatusSucceed {
			succeedNum++
		}
	}
	// 大于等于 1 就成功，否则失败
	if succeedNum > 1 {
		return model.StatusSucceed
	}
	return model.StatusFailed
}
