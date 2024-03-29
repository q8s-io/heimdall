package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func PrepareJobAnalyzer(taskImageScanInfo *model.TaskImageScanInfo) {
	// prepare job analyzer
	jobImageAnalyzerInfo := convert.JobImageAnalyzerInfoByScan(taskImageScanInfo)
	jobImageAnalyzer := convert.JobImageAnalyzer(jobImageAnalyzerInfo, 1)
	repository.NewJobImageAnalyzer(*jobImageAnalyzer)
	// mark job status
	repository.SetJobImageAnalyzerStatus(jobImageAnalyzerInfo.TaskID, model.StatusRunning)
	// send msg to mq
	jobMsg := convert.JobScannerMsg(jobImageAnalyzerInfo.TaskID, jobImageAnalyzerInfo.JobID, jobImageAnalyzerInfo.ImageName, jobImageAnalyzerInfo.ImageDigest)
	repository.SendMsgJobImageAnalyzer(jobMsg)
}

func UpdateJobImageAnalyzer(jobImageAnalyzerInfo *model.JobImageAnalyzerInfo) {
	// update job analyzer
	jobImageAnalyzer := convert.JobImageAnalyzer(jobImageAnalyzerInfo, 1)
	repository.UpdateJobImageAnalyzer(*jobImageAnalyzer)
	// mark job status
	repository.SetJobImageAnalyzerStatus(jobImageAnalyzer.TaskID, jobImageAnalyzerInfo.JobStatus)
}
