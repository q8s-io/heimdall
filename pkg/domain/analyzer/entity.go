package analyzer

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
	"github.com/q8s-io/heimdall/pkg/models"
)

func CreateJobImageAnalyzerInfo(imageVulnInfo *models.ImageVulnInfo) *models.JobImageAnalyzerInfo {
	jobImageAnalyzerInfo := new(models.JobImageAnalyzerInfo)
	jobImageAnalyzerInfo.TaskID = imageVulnInfo.TaskID
	jobImageAnalyzerInfo.JobID = distribution.GetUUID()
	jobImageAnalyzerInfo.JobStatus = models.StatusRunning
	jobImageAnalyzerInfo.ImageName = imageVulnInfo.ImageName
	jobImageAnalyzerInfo.ImageDigest = imageVulnInfo.ImageDigest
	jobImageAnalyzerInfo.ImageLayers = []string{}
	jobImageAnalyzerInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return jobImageAnalyzerInfo
}

func ConvertJobImageAnalyzerInfoByMsg(jobImageAnalyzerMsg *models.JobImageAnalyzerMsg, digest, layers []string) *models.JobImageAnalyzerInfo {
	digestList := strings.Split(digest[len(digest)-1], "@")
	jobImageAnalyzerInfo := new(models.JobImageAnalyzerInfo)
	jobImageAnalyzerInfo.TaskID = jobImageAnalyzerMsg.TaskID
	jobImageAnalyzerInfo.JobID = jobImageAnalyzerMsg.JobID
	jobImageAnalyzerInfo.JobStatus = models.StatusSucceed
	jobImageAnalyzerInfo.ImageName = jobImageAnalyzerMsg.ImageName
	jobImageAnalyzerInfo.ImageDigest = digestList[len(digestList)-1]
	jobImageAnalyzerInfo.ImageLayers = layers
	return jobImageAnalyzerInfo
}
func ConvertJobImageAnalyzerData(jobImageAnalyzerInfo *models.JobImageAnalyzerInfo, active int) *models.JobImageAnalyzerData {
	var layers string
	if len(jobImageAnalyzerInfo.ImageLayers) > 0 {
		layersByte, _ := json.Marshal(jobImageAnalyzerInfo.ImageLayers)
		layers = string(layersByte)
	}
	jobAnalyzerData := new(models.JobImageAnalyzerData)
	jobAnalyzerData.TaskID = jobImageAnalyzerInfo.TaskID
	jobAnalyzerData.JobID = jobImageAnalyzerInfo.JobID
	jobAnalyzerData.JobStatus = jobImageAnalyzerInfo.JobStatus
	jobAnalyzerData.ImageName = jobImageAnalyzerInfo.ImageName
	jobAnalyzerData.ImageDigest = jobImageAnalyzerInfo.ImageDigest
	jobAnalyzerData.ImageLayers = layers
	jobAnalyzerData.CreateTime = jobImageAnalyzerInfo.CreateTime
	jobAnalyzerData.Active = active
	return jobAnalyzerData
}

func ConvertJobImageAnalyzerMsg(jobImageAnalyzerInfo *models.JobImageAnalyzerInfo) *models.JobImageAnalyzerMsg {
	jobImageAnalyzerMsg := new(models.JobImageAnalyzerMsg)
	jobImageAnalyzerMsg.TaskID = jobImageAnalyzerInfo.TaskID
	jobImageAnalyzerMsg.JobID = jobImageAnalyzerInfo.JobID
	jobImageAnalyzerMsg.ImageName = jobImageAnalyzerInfo.ImageName
	return jobImageAnalyzerMsg
}
