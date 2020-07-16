package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func PrepareJobClair(jobImageAnalyzerInfo *model.JobImageAnalyzerInfo) {
	// preper job scanner clair
	jobClairInfo := convert.JobScannerInfoByAnalyzerInfo(jobImageAnalyzerInfo)
	jobClair := convert.JobScanner(jobClairInfo, 1)
	repository.NewJobClair(*jobClair)

	// mark job status
	repository.SetJobClairStatus(jobClairInfo.TaskID, model.StatusRunning)
	// send msg to mq
	jobClairMsg := convert.JobScannerMsg(jobClairInfo.TaskID, jobClairInfo.JobID, jobClairInfo.ImageName, jobClairInfo.ImageDigest)
	repository.SendMsgJobClair(jobClairMsg)
}

func GetJobClair(taskID string) []map[string]string {
	jobClairList := repository.GetJobClair(taskID)
	if len(*jobClairList) == 0 {
		return make([]map[string]string, 0)
	}
	jobClair := (*jobClairList)[0]
	jobClairInfo := convert.JobScannerInfo(&jobClair)
	return jobClairInfo.JobData
}

func UpdateJobClair(jobScannerInfo *model.JobScannerInfo) {
	// update job clair
	jobClair := convert.JobScanner(jobScannerInfo, 1)
	repository.UpdateJobClair(*jobClair)

	// mark job status
	repository.SetJobClairStatus(jobClair.TaskID, jobScannerInfo.JobStatus)
}
