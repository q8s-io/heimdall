package scancenter

import (
	"log"

	"github.com/pkg/errors"

	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
)

func TaskImageScanRotaryCreate(imageRequestInfo *model.ImageRequestInfo) (*model.ImageVulnInfo, error) {
	taskImageScanInfo := CreateTaskImageScan(imageRequestInfo)
	PrepareJobAnalyzer(taskImageScanInfo)
	imageVulnInfo := convert.ImageVulnByScanInfo(taskImageScanInfo, nil)
	return imageVulnInfo, nil
}

func TaskImageScanRotaryAnalyzer(jobImageAnalyzerInfo *model.JobImageAnalyzerInfo) {
	UpdateTaskImageScanDigest(jobImageAnalyzerInfo)
	UpdateJobImageAnalyzer(jobImageAnalyzerInfo)
	PrepareJobAnchore(jobImageAnalyzerInfo)
	PrepareJobTrivy(jobImageAnalyzerInfo)
	PrepareJobClair(jobImageAnalyzerInfo)
}

func TaskImageScanRotaryAnchore(jobScannerInfo *model.JobScannerInfo) {
	UpdateJobAnchore(jobScannerInfo)
	JudgeTaskRotary(jobScannerInfo.TaskID)
}

func TaskImageScanRotaryTrivy(jobScannerInfo *model.JobScannerInfo) {
	UpdateJobTrivy(jobScannerInfo)
	JudgeTaskRotary(jobScannerInfo.TaskID)
}

func TaskImageScanRotaryClair(jobScannerInfo *model.JobScannerInfo) {
	UpdateJobClair(jobScannerInfo)
	JudgeTaskRotary(jobScannerInfo.TaskID)
}

func TaskImageScanMerger(taskImageScan *entity.TaskImageScan) (interface{}, error) {
	taskID := taskImageScan.TaskID
	jobAnchoreVuln := GetJobAnchore(taskID)
	jobTrivyVuln := GetJobTrivy(taskID)
	jobClairVuln := GetJobClair(taskID)
	imageVulnData := MergerImageVulnData(taskImageScan, jobAnchoreVuln, jobTrivyVuln, jobClairVuln)
	return imageVulnData, nil
}

func MergerImageVulnData(taskImageScan *entity.TaskImageScan, jobAnchoreVuln, jobTrivyVuln, jobClairVuln []map[string]string) *model.ImageVulnInfo {
	var vulnData []map[string]interface{}
	// key: cveID  value: index in vulnData。
	cveMap := make(map[string]int)

	merge(&vulnData, &cveMap, "Trivy", jobTrivyVuln)
	merge(&vulnData, &cveMap, "Anchore", jobAnchoreVuln)
	merge(&vulnData, &cveMap, "Clair", jobClairVuln)

	taskImageScanInfo := convert.TaskImageScanInfo(taskImageScan)
	imageVulnInfo := convert.ImageVulnByScanInfo(taskImageScanInfo, vulnData)
	return imageVulnInfo
}

// Aggregate scan results
func merge(vulnData *[]map[string]interface{}, cveMap *map[string]int, engineName string, jobVuln []map[string]string) {
	// 输出结果为空的引擎
	if jobVuln == nil || len(jobVuln) == 0 {
		log.Printf("scanner %s result empty", engineName)
		return
	}

	for _, cveData := range jobVuln {
		idx, exist := (*cveMap)[cveData["cve"]]

		// 不存在
		if !exist {
			packageElement := make(map[string]string, 3)
			packageElement["package_name"] = cveData["package_name"]
			packageElement["package_version"] = cveData["package_version"]
			packageElement["package_full_nale"] = cveData["package_full_nale"]

			packageInfo := make([]map[string]string, 0)
			packageInfo = append(packageInfo, packageElement)

			// 每次添加元素都需要重新分配内存，否则都是浅拷贝，会导致切片中的元素都一样。
			curMap := make(map[string]interface{}, 3)
			curMap["cve"] = cveData["cve"]
			curMap["cve_url"] = cveData["cve_url"]
			curMap["package_info"] = packageInfo

			*vulnData = append(*vulnData, curMap)
			(*cveMap)[cveData["cve"]] = len(*vulnData) - 1

			// 存在
		} else {
			value := (*vulnData)[idx]["package_info"]

			switch value.(type) {

			case []map[string]string:
				pkgList := value.([]map[string]string)
				// 是否重复包
				repeat := false

				for _, pkg := range pkgList {
					// 全名相等
					if pkg["package_full_nale"] == cveData["package_full_nale"] {
						repeat = true
						break
					}
				}

				if !repeat {
					newPkg := make(map[string]string, 3)
					newPkg["package_name"] = cveData["package_name"]
					newPkg["package_version"] = cveData["package_version"]
					newPkg["package_full_nale"] = cveData["package_full_nale"]

					pkgList = append(pkgList, newPkg)
					(*vulnData)[idx]["package_info"] = pkgList
				}

			default:
				xray.ErrMini(errors.New("process pkg list failed"))
			}
		}
	}
}
