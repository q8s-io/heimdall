package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/domain/analyzer"
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/persistence"
	
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
)

func PreperJobAnalyzer(taskImageScanInfo *entity.TaskImageScanInfo) {
	// preper job analyzer
	jobImageAnalyzerInfo := analyzer.CreateJobImageAnalyzerInfo(taskImageScanInfo)
	jobImageAnalyzerData := analyzer.ConvertJobImageAnalyzerData(jobImageAnalyzerInfo, 1)
	persistence.NewJobImageAnalyzer(*jobImageAnalyzerData)
	// mark job status
	persistence.SetJobImageAnalyzerStatus(jobImageAnalyzerInfo.TaskID, entity.StatusRunning)
	// send msg to mq
	jobImageAnalyzerMsg := analyzer.ConvertJobImageAnalyzerMsg(jobImageAnalyzerInfo)
	persistence.SendJobImageAnalyzerMsg(jobImageAnalyzerMsg)
}

func UpdateJobImageAnalyzer(jobImageAnalyzerInfo *entity.JobImageAnalyzerInfo) {
	// update job analyzer
	jobImageAnalyzerData := analyzer.ConvertJobImageAnalyzerData(jobImageAnalyzerInfo, 1)
	persistence.UpdateJobImageAnalyzer(*jobImageAnalyzerData)
	// mark job status
	persistence.SetJobImageAnalyzerStatus(jobImageAnalyzerData.TaskID, entity.StatusSucceed)
}
