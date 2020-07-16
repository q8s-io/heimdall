package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func PrepareJobAnchore(jobImageAnalyzerInfo *model.JobImageAnalyzerInfo) {
	// prepare job scanner anchore
	jobAnchoreInfo := convert.JobScannerInfoByAnalyzerInfo(jobImageAnalyzerInfo)
	jobAnchore := convert.JobScanner(jobAnchoreInfo, 1)
	repository.NewJobAnchore(*jobAnchore)
	// mark job status
	repository.SetJobAnchoreStatus(jobAnchoreInfo.TaskID, model.StatusRunning)
	// send msg to mq
	jobAnchoreMsg := convert.JobScannerMsg(jobAnchoreInfo.TaskID, jobAnchoreInfo.JobID, jobAnchoreInfo.ImageName, jobAnchoreInfo.ImageDigest)
	repository.SendMsgJobAnchore(jobAnchoreMsg)
}

func GetJobAnchore(taskID string) []map[string]string {
	jobAnchoreList := repository.GetJobAnchore(taskID)
	if len(*jobAnchoreList) == 0 {
		return make([]map[string]string, 0)
	}
	jobAnchore := (*jobAnchoreList)[0]
	jobAnchoreInfo := convert.JobScannerInfo(&jobAnchore)
	return jobAnchoreInfo.JobData
}

func UpdateJobAnchore(jobScannerInfo *model.JobScannerInfo) {
	// update job anchore
	jobAnchore := convert.JobScanner(jobScannerInfo, 1)
	repository.UpdateJobAnchore(*jobAnchore)
	// mark job status
	repository.SetJobAnchoreStatus(jobAnchore.TaskID, jobScannerInfo.JobStatus)
}
