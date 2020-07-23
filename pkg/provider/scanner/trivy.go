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

func JobTrivy() {
	dockerConfig := model.Config.Docker
	docker.Init(dockerConfig.Host, dockerConfig.Version, nil, nil)
	// consumer msg from mq
	repository.ConsumerMsgJobTrivy()
	jobScannerMsg := new(model.JobScannerMsg)

	for msg := range kafka.Queue {
		log.Printf("consumer msg from kafka %s", msg)
		_ = json.Unmarshal(msg, &jobScannerMsg)

		// prepare trivy data
		imageName := jobScannerMsg.ImageName

		// get scanning data
		vulnData, getErr := TrivyScan(imageName)
		if getErr != nil {
			xray.ErrTaskInfo(getErr, jobScannerMsg.TaskID, jobScannerMsg.JobID)
		}

		// prepare trivy scan result
		jobTrivyInfo := PrepareTrivyScanResult(jobScannerMsg, &vulnData, getErr)

		// send to scancenter
		requestJSON, _ := json.Marshal(jobTrivyInfo)
		log.Printf("trivy process succeed %s", imageName)
		_ = net.HTTPPUT(model.Config.ScanCenter.TrivyURL, string(requestJSON))
	}
}

func PrepareTrivyScanResult(jobScannerMsg *model.JobScannerMsg, vulnData *model.TrivyScanResult, runErr error) *model.JobScannerInfo {
	var cveList []map[string]string
	for _, vulnInfo := range vulnData.Vulnerabilities {
		cve := make(map[string]string)

		cve["package_name"] = vulnInfo.PkgName
		cve["package_version"] = vulnInfo.InstalledVersion
		cve["package_full_nale"] = fmt.Sprintf("%s-%s", vulnInfo.PkgName, vulnInfo.InstalledVersion)
		cve["cve"] = vulnInfo.VulnerabilityID
		// process url
		cve["cve_url"] = fmt.Sprintf("http://cve.mitre.org/cgi-bin/cvename.cgi?name=%s", cve["cve"])
		cveList = append(cveList, cve)
	}

	jobScannerInfo := new(model.JobScannerInfo)
	jobScannerInfo.TaskID = jobScannerMsg.TaskID
	jobScannerInfo.JobID = jobScannerMsg.JobID

	// 判断运行过程是否异常
	if runErr != nil {
		jobScannerInfo.JobStatus = model.StatusFailed
	} else {
		jobScannerInfo.JobStatus = model.StatusSucceed
	}
	jobScannerInfo.JobData = cveList

	return jobScannerInfo
}
