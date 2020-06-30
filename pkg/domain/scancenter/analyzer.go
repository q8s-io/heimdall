package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/domain/analyzer"
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
)

func PreperJobAnalyzer(imageVulnInfo *models.ImageVulnInfo) {
	// preper job analyzer
	jobImageAnalyzerInfo := analyzer.CreateJobImageAnalyzerInfo(imageVulnInfo)
	jobImageAnalyzerData := analyzer.ConvertJobImageAnalyzerData(jobImageAnalyzerInfo, 1)
	service.NewJobImageAnalyzer(*jobImageAnalyzerData)
	// mark job status
	service.SetJobImageAnalyzerStatus(jobImageAnalyzerInfo.TaskID, models.StatusRunning)
	// send msg to mq
	jobImageAnalyzerMsg := analyzer.ConvertJobImageAnalyzerMsg(jobImageAnalyzerInfo)
	service.SendJobImageAnalyzerMsg(jobImageAnalyzerMsg)
}

func UpdateJobImageAnalyzer(jobImageAnalyzerInfo *models.JobImageAnalyzerInfo) {
	// update job analyzer
	jobImageAnalyzerData := analyzer.ConvertJobImageAnalyzerData(jobImageAnalyzerInfo, 1)
	service.UpdateJobImageAnalyzer(*jobImageAnalyzerData)
	// mark job status
	service.SetJobImageAnalyzerStatus(jobImageAnalyzerData.TaskID, models.StatusSucceed)
}
