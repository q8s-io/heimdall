package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/domain/scanner"
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/persistence"
	
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
)

func PreperJobAnchore(jobImageAnalyzerInfo *entity.JobImageAnalyzerInfo) {
	// preper job scanner anchore
	jobAnchoreInfo := scanner.CreateJobAnchoreInfo(jobImageAnalyzerInfo)
	jobAnchoreData := scanner.ConvertJobAnchoreData(jobAnchoreInfo, 1)
	persistence.NewJobAnchore(*jobAnchoreData)
	// mark job status
	persistence.SetJobAnchoreStatus(jobAnchoreInfo.TaskID, entity.StatusRunning)
	// send msg to mq
	jobAnchoreMsg := scanner.ConvertJobAnchoreMsg(jobAnchoreInfo)
	persistence.SendJobAnchoreMsg(jobAnchoreMsg)
}

func GetJobAnchore(taskID string) []map[string]string {
	jobAnchoreDataList := persistence.GetJobAnchore(taskID)
	jobAnchoreData := (*jobAnchoreDataList)[0]
	jobAnchoreInfo := scanner.ConvertJobAnchoreInfo(&jobAnchoreData)
	return jobAnchoreInfo.JobData
}

func UpdateJobAnchore(jobAnchoreInfo *entity.JobAnchoreInfo) {
	// update job anchore
	jobAnchoreData := scanner.ConvertJobAnchoreData(jobAnchoreInfo, 1)
	persistence.UpdateJobAnchore(*jobAnchoreData)
	// mark job status
	persistence.SetJobAnchoreStatus(jobAnchoreData.TaskID, entity.StatusSucceed)
}
