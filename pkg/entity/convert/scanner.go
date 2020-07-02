package convert

import (
	"encoding/json"
	"time"

	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
)

func JobScanner(jobScannerInfo *model.JobScannerInfo, active int) *entity.JobScanner {
	var jobData string
	if len(jobScannerInfo.JobData) > 0 {
		jobDataByte, _ := json.Marshal(jobScannerInfo.JobData)
		jobData = string(jobDataByte)
	}
	JobScanner := new(entity.JobScanner)
	JobScanner.TaskID = jobScannerInfo.TaskID
	JobScanner.JobID = jobScannerInfo.JobID
	JobScanner.JobStatus = jobScannerInfo.JobStatus
	JobScanner.JobData = jobData
	JobScanner.ImageName = jobScannerInfo.ImageName
	JobScanner.ImageDigest = jobScannerInfo.ImageDigest
	JobScanner.CreateTime = jobScannerInfo.CreateTime
	JobScanner.Active = active
	return JobScanner
}

func JobScannerInfo(jobScanner *entity.JobScanner) *model.JobScannerInfo {
	jobData := make([]map[string]string, 0)
	_ = json.Unmarshal([]byte(jobScanner.JobData), &jobData)
	jobScannerInfo := new(model.JobScannerInfo)
	jobScannerInfo.TaskID = jobScanner.TaskID
	jobScannerInfo.JobID = jobScanner.JobID
	jobScannerInfo.JobStatus = jobScanner.JobStatus
	jobScannerInfo.JobData = jobData
	jobScannerInfo.ImageName = jobScanner.ImageName
	jobScannerInfo.ImageDigest = jobScanner.ImageDigest
	jobScannerInfo.CreateTime = jobScanner.CreateTime
	return jobScannerInfo
}

func JobScannerInfoByAnalyzerInfo(jobImageAnalyzerInfo *model.JobImageAnalyzerInfo) *model.JobScannerInfo {
	jobScannerInfo := new(model.JobScannerInfo)
	jobScannerInfo.TaskID = jobImageAnalyzerInfo.TaskID
	jobScannerInfo.JobID = distribution.GetUUID()
	jobScannerInfo.JobStatus = model.StatusRunning
	jobScannerInfo.ImageName = jobImageAnalyzerInfo.ImageName
	jobScannerInfo.ImageDigest = jobImageAnalyzerInfo.ImageDigest
	jobScannerInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return jobScannerInfo
}

func JobScannerMsg(taskID, jobID, imageName, imageDigest string) *model.JobScannerMsg {
	jobScannerMsg := new(model.JobScannerMsg)
	jobScannerMsg.TaskID = taskID
	jobScannerMsg.JobID = jobID
	jobScannerMsg.ImageName = imageName
	jobScannerMsg.ImageDigest = imageDigest
	return jobScannerMsg
}
