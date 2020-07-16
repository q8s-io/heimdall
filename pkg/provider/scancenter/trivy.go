package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func PrepareJobTrivy(jobImageAnalyzerInfo *model.JobImageAnalyzerInfo) {
	// preper job scanner trivy
	jobTrivyInfo := convert.JobScannerInfoByAnalyzerInfo(jobImageAnalyzerInfo)
	jobTrivy := convert.JobScanner(jobTrivyInfo, 1)
	repository.NewJobTrivy(*jobTrivy)

	// mark job status
	repository.SetJobTrivyStatus(jobTrivyInfo.TaskID, model.StatusRunning)
	// send msg to mq
	jobTrivyMsg := convert.JobScannerMsg(jobTrivyInfo.TaskID, jobTrivyInfo.JobID, jobTrivyInfo.ImageName, jobTrivyInfo.ImageDigest)
	repository.SendMsgJobTrivy(jobTrivyMsg)
}

func GetJobTrivy(taskID string) []map[string]string {
	jobTrivyList := repository.GetJobTrivy(taskID)
	if len(*jobTrivyList) == 0 {
		return make([]map[string]string, 0)
	}
	jobTrivy := (*jobTrivyList)[0]
	jobTrivyInfo := convert.JobScannerInfo(&jobTrivy)
	return jobTrivyInfo.JobData
}

func UpdateJobTrivy(jobScannerInfo *model.JobScannerInfo) {
	// update job trivy
	jobTrivy := convert.JobScanner(jobScannerInfo, 1)
	repository.UpdateJobTrivy(*jobTrivy)

	// mark job status
	repository.SetJobTrivyStatus(jobTrivy.TaskID, jobScannerInfo.JobStatus)
}
