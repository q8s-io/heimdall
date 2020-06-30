package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/domain/scanner"
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
)

func PreperJobAnchore(jobImageAnalyzerInfo *models.JobImageAnalyzerInfo) {
	// preper job scanner anchore
	jobAnchoreInfo := scanner.CreateJobAnchoreInfo(jobImageAnalyzerInfo)
	jobAnchoreData := scanner.ConvertJobAnchoreData(jobAnchoreInfo, 1)
	service.NewJobAnchore(*jobAnchoreData)
	// mark job status
	service.SetJobAnchoreStatus(jobAnchoreInfo.TaskID, models.StatusRunning)
	// send msg to mq
	jobAnchoreMsg := scanner.ConvertJobAnchoreMsg(jobAnchoreInfo)
	service.SendJobAnchoreMsg(jobAnchoreMsg)
}

func GetJobAnchore(taskID string) []map[string]string {
	jobAnchoreDataList := service.GetJobAnchore(taskID)
	jobAnchoreData := (*jobAnchoreDataList)[0]
	jobAnchoreInfo := scanner.ConvertJobAnchoreInfo(&jobAnchoreData)
	return jobAnchoreInfo.JobData
}

func UpdateJobAnchore(jobAnchoreInfo *models.JobAnchoreInfo) {
	// update job anchore
	jobAnchoreData := scanner.ConvertJobAnchoreData(jobAnchoreInfo, 1)
	service.UpdateJobAnchore(*jobAnchoreData)
	// mark job status
	service.SetJobAnchoreStatus(jobAnchoreData.TaskID, models.StatusSucceed)
}
