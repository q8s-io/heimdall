package convert

import (
	"encoding/json"
	"strings"
	"time"
	
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
	"github.com/q8s-io/heimdall/pkg/models"
)

func CreateJobImageAnalyzerInfo(taskImageScanInfo *entity.TaskImageScanInfo) *entity.JobImageAnalyzerInfo {
	jobImageAnalyzerInfo := new(entity.JobImageAnalyzerInfo)
	jobImageAnalyzerInfo.TaskID = taskImageScanInfo.TaskID
	jobImageAnalyzerInfo.JobID = distribution.GetUUID()
	jobImageAnalyzerInfo.JobStatus = entity.StatusRunning
	jobImageAnalyzerInfo.ImageName = taskImageScanInfo.ImageName
	jobImageAnalyzerInfo.ImageDigest = taskImageScanInfo.ImageDigest
	jobImageAnalyzerInfo.ImageLayers = []string{}
	jobImageAnalyzerInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return jobImageAnalyzerInfo
}

func ConvertJobImageAnalyzerInfoByMsg(jobImageAnalyzerMsg *entity.JobImageAnalyzerMsg, digest, layers []string) *entity.JobImageAnalyzerInfo {
	digestList := strings.Split(digest[len(digest)-1], "@")
	jobImageAnalyzerInfo := new(entity.JobImageAnalyzerInfo)
	jobImageAnalyzerInfo.TaskID = jobImageAnalyzerMsg.TaskID
	jobImageAnalyzerInfo.JobID = jobImageAnalyzerMsg.JobID
	jobImageAnalyzerInfo.JobStatus = entity.StatusSucceed
	jobImageAnalyzerInfo.ImageName = jobImageAnalyzerMsg.ImageName
	jobImageAnalyzerInfo.ImageDigest = digestList[len(digestList)-1]
	jobImageAnalyzerInfo.ImageLayers = layers
	return jobImageAnalyzerInfo
}

func ConvertJobImageAnalyzerData(jobImageAnalyzerInfo *entity.JobImageAnalyzerInfo, active int) *entity.JobImageAnalyzerData {
	var layers string
	if len(jobImageAnalyzerInfo.ImageLayers) > 0 {
		layersByte, _ := json.Marshal(jobImageAnalyzerInfo.ImageLayers)
		layers = string(layersByte)
	}
	jobAnalyzerData := new(entity.JobImageAnalyzerData)
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

func ConvertJobImageAnalyzerMsg(jobImageAnalyzerInfo *entity.JobImageAnalyzerInfo) *entity.JobImageAnalyzerMsg {
	jobImageAnalyzerMsg := new(entity.JobImageAnalyzerMsg)
	jobImageAnalyzerMsg.TaskID = jobImageAnalyzerInfo.TaskID
	jobImageAnalyzerMsg.JobID = jobImageAnalyzerInfo.JobID
	jobImageAnalyzerMsg.ImageName = jobImageAnalyzerInfo.ImageName
	return jobImageAnalyzerMsg
}
