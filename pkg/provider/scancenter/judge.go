package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
	"github.com/q8s-io/heimdall/pkg/repository"
	"strings"
	"time"
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
			creatTaskTime, timeErr := time.Parse("2006-01-02 15:04:05", taskImageScan.CreateTime)
			if timeErr != nil {
				xray.ErrMini(timeErr)
			}
			// 注意时区，默认是UTC标准时间（北京时间减去八小时）。
			// latest Tag 大于一个小时触发扫描
			if time.Now().Sub(creatTaskTime) > time.Hour {
				UpdateTaskImageScanActive(imageRequestInfo)
				return TaskImageScanRotaryCreate(imageRequestInfo)
			}
		}
		// 异步写入kafka
		imageVulnData, getScanErr := TaskImageScanMerger(&taskImageScan)
		if getScanErr == nil {
			result := PrepareKafkaData(imageVulnData)
			kafka.AsyncProducerSendMsg("imageCVE", result)
		}
		return imageVulnData, getScanErr
	}
	taskImageScanInfo := convert.TaskImageScanInfo(&taskImageScan)
	imageVulnInfo := convert.ImageVulnByScanInfo(taskImageScanInfo, nil)
	return imageVulnInfo, nil
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
	// 镜像分析失败则任务失败
	if taskStatus["analyzed"] == model.StatusFailed || taskStatus["analyzed"] == model.StatusTimeout {
		return taskStatus["analyzed"]
	}
	// count success and timeout num
	succeedNum := 0
	timeoutNum := 0
	for _, v := range taskStatus {
		// 若还有在运行的引擎，直接返回运行态
		switch v {
		case model.StatusRunning:
			return model.StatusRunning
		case model.StatusSucceed:
			succeedNum++
		case model.StatusTimeout:
			timeoutNum++
		}
	}
	// 结果状态优先级：成功、超时、失败。
	// 大于等于 1 就成功
	if succeedNum > 1 {
		return model.StatusSucceed
	}
	// 没有成功的，有超时的状态就返回超时。
	if timeoutNum > 0 {
		return model.StatusTimeout
	}
	// 都是失败，返回失败
	return model.StatusFailed
}
