package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/models"
)

func TaskImageScanRotaryCreate(imageRequestInfo *models.ImageRequestInfo) (interface{}, error) {
	imageVulnInfo, err := CreateTaskImageScan(imageRequestInfo)
	if err != nil {
		return nil, err
	}
	PreperJobAnalyzer(imageVulnInfo)
	return imageVulnInfo, nil
}

func TaskImageScanRotaryAnalyzer(jobImageAnalyzerInfo *models.JobImageAnalyzerInfo) {
	UpdateJobImageAnalyzer(jobImageAnalyzerInfo)
	PreperJobAnchore(jobImageAnalyzerInfo)
}

func TaskImageScanRotaryAnchore(jobAnchoreInfo *models.JobAnchoreInfo) {
	UpdateJobAnchore(jobAnchoreInfo)
	JudgeTaskRotary(jobAnchoreInfo.TaskID)
}
