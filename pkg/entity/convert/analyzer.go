package convert

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/distribution"
)

func JobImageAnalyzer(jobImageAnalyzerInfo *model.JobImageAnalyzerInfo, active int) *entity.JobImageAnalyzer {
	var layers string
	if len(jobImageAnalyzerInfo.ImageLayers) > 0 {
		layersByte, _ := json.Marshal(jobImageAnalyzerInfo.ImageLayers)
		layers = string(layersByte)
	}
	jobImageAnalyzer := new(entity.JobImageAnalyzer)
	jobImageAnalyzer.TaskID = jobImageAnalyzerInfo.TaskID
	jobImageAnalyzer.JobID = jobImageAnalyzerInfo.JobID
	jobImageAnalyzer.JobStatus = jobImageAnalyzerInfo.JobStatus
	jobImageAnalyzer.ImageName = jobImageAnalyzerInfo.ImageName
	jobImageAnalyzer.ImageDigest = jobImageAnalyzerInfo.ImageDigest
	jobImageAnalyzer.ImageLayers = layers
	jobImageAnalyzer.CreateTime = jobImageAnalyzerInfo.CreateTime
	jobImageAnalyzer.Active = active
	return jobImageAnalyzer
}

func JobImageAnalyzerInfoByScan(taskImageScanInfo *model.TaskImageScanInfo) *model.JobImageAnalyzerInfo {
	jobImageAnalyzerInfo := new(model.JobImageAnalyzerInfo)
	jobImageAnalyzerInfo.TaskID = taskImageScanInfo.TaskID
	jobImageAnalyzerInfo.JobID = distribution.GetUUID()
	jobImageAnalyzerInfo.JobStatus = model.StatusRunning
	jobImageAnalyzerInfo.ImageName = taskImageScanInfo.ImageName
	jobImageAnalyzerInfo.ImageDigest = taskImageScanInfo.ImageDigest
	jobImageAnalyzerInfo.ImageLayers = []string{}
	jobImageAnalyzerInfo.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return jobImageAnalyzerInfo
}

func JobImageAnalyzerInfoByMsg(jobScannerMsg *model.JobScannerMsg, digest, layers []string) *model.JobImageAnalyzerInfo {
	digestList := strings.Split(digest[len(digest)-1], "@")
	jobImageAnalyzerInfo := new(model.JobImageAnalyzerInfo)
	jobImageAnalyzerInfo.TaskID = jobScannerMsg.TaskID
	jobImageAnalyzerInfo.JobID = jobScannerMsg.JobID
	jobImageAnalyzerInfo.JobStatus = model.StatusSucceed
	jobImageAnalyzerInfo.ImageName = jobScannerMsg.ImageName
	jobImageAnalyzerInfo.ImageDigest = digestList[len(digestList)-1]
	jobImageAnalyzerInfo.ImageLayers = layers
	return jobImageAnalyzerInfo
}
