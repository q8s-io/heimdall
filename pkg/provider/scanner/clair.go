package scanner

import (
	"encoding/json"
	"fmt"
	"github.com/q8s-io/heimdall/pkg/infrastructure/docker"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
	"log"

	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
	"github.com/q8s-io/heimdall/pkg/infrastructure/net"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func JobClair() {
	dockerConfig := model.Config.Docker
	docker.Init(dockerConfig.Host, dockerConfig.Version, nil, nil)
	// consumer msg from mq
	repository.ConsumerMsgJobClair()
	jobScannerMsg := new(model.JobScannerMsg)

	for msg := range kafka.Queue {
		log.Printf("consumer msg from kafka: %s", msg)
		_ = json.Unmarshal(msg, &jobScannerMsg)

		// prepare clair data
		imageName := jobScannerMsg.ImageName

		// get scanning data
		vulnData, getErr := ClairScan(imageName)
		if getErr != nil {
			xray.ErrTaskInfo(getErr, jobScannerMsg.TaskID, jobScannerMsg.JobID)
		}

		// prepare clair scan result
		jobClairInfo := PrepareClairScanResult(jobScannerMsg, &vulnData, getErr)

		// send to scancenter
		requestJSON, _ := json.Marshal(jobClairInfo)
		log.Printf("anchore process succeed %s", imageName)
		_ = net.HTTPPUT(model.Config.ScanCenter.ClairURL, string(requestJSON))
	}
}

func PrepareClairScanResult(jobScannerMsg *model.JobScannerMsg, vulnData *model.ClairScanResult, runErr error) *model.JobScannerInfo {
	var cveList []map[string]string

	vulnerabilities := vulnData.Vulnerabilities
	ScanGrade(&cveList, vulnerabilities.Unknown)
	ScanGrade(&cveList, vulnerabilities.High)
	ScanGrade(&cveList, vulnerabilities.Medium)
	ScanGrade(&cveList, vulnerabilities.Low)
	ScanGrade(&cveList, vulnerabilities.Critical)
	ScanGrade(&cveList, vulnerabilities.Defcon1)
	ScanGrade(&cveList, vulnerabilities.Negligible)

	jobScannerInfo := new(model.JobScannerInfo)
	jobScannerInfo.TaskID = jobScannerMsg.TaskID
	jobScannerInfo.JobID = jobScannerMsg.JobID
	if runErr != nil {
		jobScannerInfo.JobStatus = model.StatusFailed
	} else {
		jobScannerInfo.JobStatus = model.StatusSucceed
	}
	jobScannerInfo.JobData = cveList

	return jobScannerInfo
}

func ScanGrade(cveList *[]map[string]string, grade []model.Grade) {
	if grade == nil || len(grade) == 0 {
		return
	}
	for _, vulnInfo := range grade {
		cve := make(map[string]string)
		cve["package_name"] = vulnInfo.FeatureName
		cve["package_version"] = vulnInfo.FeatureVersion
		cve["package_full_nale"] = fmt.Sprintf("%s-%s", vulnInfo.FeatureName, vulnInfo.FeatureVersion)
		cve["cve"] = vulnInfo.Name
		// process url
		cve["cve_url"] = fmt.Sprintf("http://cve.mitre.org/cgi-bin/cvename.cgi?name=%s", cve["cve"])
		*cveList = append(*cveList, cve)
	}
}
