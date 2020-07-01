package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/models"
	
	"github.com/q8s-io/heimdall/pkg/entity"
)

func TaskImageScanRotaryCreate(imageRequestInfo *entity.ImageRequestInfo) (*entity.ImageVulnData, error) {
	taskImageScanInfo, err := CreateTaskImageScan(imageRequestInfo)
	if err != nil {
		return nil, err
	}
	PreperJobAnalyzer(taskImageScanInfo)
	imageVulnData := CreateImageVulnData(taskImageScanInfo)
	return imageVulnData, nil
}

func TaskImageScanRotaryAnalyzer(jobImageAnalyzerInfo *entity.JobImageAnalyzerInfo) {
	UpdateTaskImageScanDigest(jobImageAnalyzerInfo)
	UpdateJobImageAnalyzer(jobImageAnalyzerInfo)
	PreperJobAnchore(jobImageAnalyzerInfo)
}

func TaskImageScanRotaryAnchore(jobAnchoreInfo *entity.JobAnchoreInfo) {
	UpdateJobAnchore(jobAnchoreInfo)
	JudgeTaskRotary(jobAnchoreInfo.TaskID)
}

func TaskImageScanMerger(taskImageScanData *entity.TaskImageScanData) (interface{}, error) {
	taskID := taskImageScanData.TaskID
	jobAnchoreVuln := GetJobAnchore(taskID)
	imageVulnData := ImageVulnDataMerger(taskImageScanData, jobAnchoreVuln)
	return imageVulnData, nil
}

func ImageVulnDataMerger(taskImageScanData *entity.TaskImageScanData, jobAnchoreVuln []map[string]string) *entity.ImageVulnData {
	vulnData := jobAnchoreVuln
	imageVulnData := ConvertImageVulnData(taskImageScanData, vulnData)
	return imageVulnData
}
