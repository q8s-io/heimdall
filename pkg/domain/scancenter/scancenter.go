package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/models"
)

func TaskImageScanRotaryCreate(imageRequestInfo *models.ImageRequestInfo) (*models.ImageVulnData, error) {
	taskImageScanInfo, err := CreateTaskImageScan(imageRequestInfo)
	if err != nil {
		return nil, err
	}
	PreperJobAnalyzer(taskImageScanInfo)
	imageVulnData := CreateImageVulnData(taskImageScanInfo)
	return imageVulnData, nil
}

func TaskImageScanRotaryAnalyzer(jobImageAnalyzerInfo *models.JobImageAnalyzerInfo) {
	UpdateTaskImageScanDigest(jobImageAnalyzerInfo)
	UpdateJobImageAnalyzer(jobImageAnalyzerInfo)
	PreperJobAnchore(jobImageAnalyzerInfo)
}

func TaskImageScanRotaryAnchore(jobAnchoreInfo *models.JobAnchoreInfo) {
	UpdateJobAnchore(jobAnchoreInfo)
	JudgeTaskRotary(jobAnchoreInfo.TaskID)
}

func TaskImageScanMerger(taskImageScanData *models.TaskImageScanData) (interface{}, error) {
	taskID := taskImageScanData.TaskID
	jobAnchoreVuln := GetJobAnchore(taskID)
	imageVulnData := ImageVulnDataMerger(taskImageScanData, jobAnchoreVuln)
	return imageVulnData, nil
}

func ImageVulnDataMerger(taskImageScanData *models.TaskImageScanData, jobAnchoreVuln []map[string]string) *models.ImageVulnData {
	vulnData := jobAnchoreVuln
	imageVulnData := ConvertImageVulnData(taskImageScanData, vulnData)
	return imageVulnData
}
