package scanner

import (
	"encoding/json"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
	"github.com/q8s-io/heimdall/pkg/infrastructure/net"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func JobTrivy() {
	// consumer msg from mq
	repository.ConsumerMsgJobTrivy()
	jobScannerMsg := new(model.JobScannerMsg)

	for msg := range kafka.Queue {
		_ = json.Unmarshal(msg, &jobScannerMsg)

		// prepare trivy data
		imageName := jobScannerMsg.ImageName

		// get scanning data
		vulnData := TrivyScan(imageName)

		// prepare trivy scan result
		jobTrivyInfo := PreperTrivyScanResult(jobScannerMsg, &vulnData)

		// send to scancenter
		requestJSON, _ := json.Marshal(jobTrivyInfo)
		_ = net.HTTPPUT(model.Config.ScanCenter.TrivyURL, string(requestJSON))
	}
}

func PreperTrivyScanResult(jobScannerMsg *model.JobScannerMsg, vulnData *model.TrivyScanResult) *model.JobScannerInfo {
	var cveList []map[string]string
	for _, vulnInfo := range vulnData.Vulnerabilities {
		cve := make(map[string]string)

		cve["package_name"] = vulnInfo.PkgName
		cve["package_version"] = vulnInfo.InstalledVersion
		cve["package_full_nale"] = vulnInfo.PkgName + "-" + vulnInfo.InstalledVersion
		cve["cve"] = vulnInfo.VulnerabilityID
		// process url
		cve["cve_url"] = "http://cve.mitre.org/cgi-bin/cvename.cgi?name=" + cve["cve"]
		cveList = append(cveList, cve)
	}
	jobScannerInfo := new(model.JobScannerInfo)
	jobScannerInfo.TaskID = jobScannerMsg.TaskID
	jobScannerInfo.JobID = jobScannerMsg.JobID
	jobScannerInfo.JobStatus = model.StatusSucceed
	jobScannerInfo.JobData = cveList
	return jobScannerInfo
}
