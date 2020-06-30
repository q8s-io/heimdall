package scanner

import (
	"encoding/json"
	"time"

	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
	"github.com/q8s-io/heimdall/pkg/models"
)

func CreateJobAnchoreInfo(jobImageAnalyzerInfo *models.JobImageAnalyzerInfo) *models.JobAnchoreInfo {
	jobAnchoreInfo := new(models.JobAnchoreInfo)
	jobAnchoreInfo.TaskID = jobImageAnalyzerInfo.TaskID
	jobAnchoreInfo.JobID = distribution.GetUUID()
	jobAnchoreInfo.JobStatus = models.StatusRunning
	jobAnchoreInfo.ImageName = jobImageAnalyzerInfo.ImageName
	jobAnchoreInfo.ImageDigest = jobImageAnalyzerInfo.ImageDigest
	jobAnchoreInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return jobAnchoreInfo
}

func ConvertJobAnchoreInfo(jobAnchoreData *models.JobAnchoreData) *models.JobAnchoreInfo {
	jobData := make([]map[string]string, 0)
	_ = json.Unmarshal([]byte(jobAnchoreData.JobData), &jobData)
	jobAnchoreInfo := new(models.JobAnchoreInfo)
	jobAnchoreInfo.TaskID = jobAnchoreData.TaskID
	jobAnchoreInfo.JobID = jobAnchoreData.JobID
	jobAnchoreInfo.JobStatus = jobAnchoreData.JobStatus
	jobAnchoreInfo.JobData = jobData
	jobAnchoreInfo.ImageName = jobAnchoreData.ImageName
	jobAnchoreInfo.ImageDigest = jobAnchoreData.ImageDigest
	jobAnchoreInfo.CreateTime = jobAnchoreData.CreateTime
	return jobAnchoreInfo
}

func ConvertJobAnchoreData(jobAnchoreInfo *models.JobAnchoreInfo, active int) *models.JobAnchoreData {
	var jobData string
	if len(jobAnchoreInfo.JobData) > 0 {
		jobDataByte, _ := json.Marshal(jobAnchoreInfo.JobData)
		jobData = string(jobDataByte)
	}
	jobAnchoreData := new(models.JobAnchoreData)
	jobAnchoreData.TaskID = jobAnchoreInfo.TaskID
	jobAnchoreData.JobID = jobAnchoreInfo.JobID
	jobAnchoreData.JobStatus = jobAnchoreInfo.JobStatus
	jobAnchoreData.JobData = jobData
	jobAnchoreData.ImageName = jobAnchoreInfo.ImageName
	jobAnchoreData.ImageDigest = jobAnchoreInfo.ImageDigest
	jobAnchoreData.CreateTime = jobAnchoreInfo.CreateTime
	jobAnchoreData.Active = active
	return jobAnchoreData
}

func ConvertJobAnchoreMsg(jobAnchoreInfo *models.JobAnchoreInfo) *models.JobAnchoreMsg {
	jobAnchoreMsg := new(models.JobAnchoreMsg)
	jobAnchoreMsg.TaskID = jobAnchoreInfo.TaskID
	jobAnchoreMsg.JobID = jobAnchoreInfo.JobID
	jobAnchoreMsg.ImageName = jobAnchoreInfo.ImageName
	jobAnchoreMsg.ImageDigest = jobAnchoreInfo.ImageDigest
	return jobAnchoreMsg
}

func CreateAnchoreRequestInfo(jobAnchoreMsg *models.JobAnchoreMsg) *models.AnchoreRequestInfo {
	anchoreRequestInfo := new(models.AnchoreRequestInfo)
	anchoreRequestInfo.ImageName = jobAnchoreMsg.ImageName
	anchoreRequestInfo.ImageDigest = jobAnchoreMsg.ImageDigest
	anchoreRequestInfo.CreateTime = time.Now().Format("2006-01-02T15:04:05Z")
	return anchoreRequestInfo
}
