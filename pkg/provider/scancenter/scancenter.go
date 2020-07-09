package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"log"
)

func TaskImageScanRotaryCreate(imageRequestInfo *model.ImageRequestInfo) (*model.ImageVulnInfo, error) {
	taskImageScanInfo, err := CreateTaskImageScan(imageRequestInfo)
	if err != nil {
		return nil, err
	}
	PrepareJobAnalyzer(taskImageScanInfo)
	imageVulnInfo := convert.ImageVulnByScanInfo(taskImageScanInfo, nil)
	return imageVulnInfo, nil
}

func TaskImageScanRotaryAnalyzer(jobImageAnalyzerInfo *model.JobImageAnalyzerInfo) {
	UpdateTaskImageScanDigest(jobImageAnalyzerInfo)
	UpdateJobImageAnalyzer(jobImageAnalyzerInfo)
	PrepareJobAnchore(jobImageAnalyzerInfo)
	PrepareJobTrivy(jobImageAnalyzerInfo)
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
	jobAnchoreVuln := GetJobAnchore(taskID)
	jobTrivyVuln := GetJobTrivy(taskID)
	imageVulnData := MergerImageVulnData(taskImageScan, jobAnchoreVuln, jobTrivyVuln)
	// imageVulnData := MergerImageVulnData(taskImageScan, jobTrivyVuln)
	return imageVulnData, nil
}

func MergerImageVulnData(taskImageScan *entity.TaskImageScan, jobAnchoreVuln []map[string]string, jobTrivyVuln []map[string]string) *model.ImageVulnInfo {
	var vulnData []map[string]interface{}
	cveMap := make(map[string]int)

	merge(&vulnData, &cveMap, jobTrivyVuln)
	merge(&vulnData, &cveMap, jobAnchoreVuln)

	taskImageScanInfo := convert.TaskImageScanInfo(taskImageScan)
	imageVulnInfo := convert.ImageVulnByScanInfo(taskImageScanInfo, vulnData)
	return imageVulnInfo
}

// Aggregate engine scan results
func merge(vulnData *[]map[string]interface{}, cveMap *map[string]int, jobVuln []map[string]string) {

	for _, cveData := range jobVuln {
		idx, exist := (*cveMap)[cveData["cve"]]
		if !exist { // 不存在
			// 每次添加元素都需要重新分配内存，否则都是浅拷贝，会导致切片中的元素都一样。
			curMap := make(map[string]interface{})

			curMap["cve"] = cveData["cve"]
			curMap["cve_url"] = cveData["cve_url"]
			curMap["package_info"] = []map[string]string{
				{"package_name": cveData["package_name"],
					"package_version":   cveData["package_version"],
					"package_full_nale": cveData["package_full_nale"]},
			}

			*vulnData = append(*vulnData, curMap)
			(*cveMap)[cveData["cve"]] = len(*vulnData) - 1
		} else { // 存在
			value := (*vulnData)[idx]["package_info"]

			switch value.(type) {
			case []map[string]string:
				pkgList := value.([]map[string]string)
				repeat := false // 是否重复包

				for _, pkg := range pkgList {
					// 全名相等
					if pkg["package_full_nale"] == cveData["package_full_nale"] {
						repeat = true
						break
					}
				}

				if !repeat {
					newPkg := map[string]string{"package_name": cveData["package_name"],
						"package_version":   cveData["package_version"],
						"package_full_nale": cveData["package_full_nale"]}
					pkgList = append(pkgList, newPkg)
					(*vulnData)[idx]["package_info"] = pkgList
				}
			default:
				log.Print("process pkg list err !!!")
			}
		}
	}
}
