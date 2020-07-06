package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
)

func TaskImageScanRotaryCreate(imageRequestInfo *model.ImageRequestInfo) (*model.ImageVulnInfo, error) {
	taskImageScanInfo, err := CreateTaskImageScan(imageRequestInfo)
	if err != nil {
		return nil, err
	}
	PreperJobAnalyzer(taskImageScanInfo)
	imageVulnInfo := convert.ImageVulnByScanInfo(taskImageScanInfo, nil)
	return imageVulnInfo, nil
}

func TaskImageScanRotaryAnalyzer(jobImageAnalyzerInfo *model.JobImageAnalyzerInfo) {
	UpdateTaskImageScanDigest(jobImageAnalyzerInfo)
	UpdateJobImageAnalyzer(jobImageAnalyzerInfo)
	PreperJobAnchore(jobImageAnalyzerInfo)
	PreperJobTrivy(jobImageAnalyzerInfo)
}

func TaskImageScanRotaryAnchore(jobScannerInfo *model.JobScannerInfo) {
	UpdateJobAnchore(jobScannerInfo)
	JudgeTaskRotary(jobScannerInfo.TaskID)
}

func TaskImageScanRotaryTrivy(jobScannerInfo *model.JobScannerInfo) {
	UpdateJobTrivy(jobScannerInfo)
	JudgeTaskRotary(jobScannerInfo.TaskID)
}

func TaskImageScanMerger(taskImageScan *entity.TaskImageScan) (interface{}, error) {
	taskID := taskImageScan.TaskID
	// jobAnchoreVuln := GetJobAnchore(taskID)
	jobTrivyVuln := GetJobTrivy(taskID)

	// imageVulnData := MergerImageVulnData(taskImageScan, jobAnchoreVuln)
	imageVulnData := MergerImageVulnData(taskImageScan, jobTrivyVuln)
	return imageVulnData, nil
}

func MergerImageVulnData(taskImageScan *entity.TaskImageScan, jobAnchoreVuln []map[string]string) *model.ImageVulnInfo {
	vulnData := jobAnchoreVuln
	taskImageScanInfo := convert.TaskImageScanInfo(taskImageScan)
	imageVulnInfo := convert.ImageVulnByScanInfo(taskImageScanInfo, vulnData)
	return imageVulnInfo
}
