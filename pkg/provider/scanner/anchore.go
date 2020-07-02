package scanner

import (
	"encoding/json"
	"time"

	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
	"github.com/q8s-io/heimdall/pkg/infrastructure/net"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func JobAnchore() {
	// consumer msg from mq
	repository.ConsumerMsgJobAnchore()
	jobScannerMsg := new(model.JobScannerMsg)
	for msg := range kafka.Queue {
		_ = json.Unmarshal(msg, &jobScannerMsg)

		// preper anchore data
		anchoreRequestInfo := convert.AnchoreRequestInfo(jobScannerMsg)

		// trigger anchore scan
		TriggerAnchoreScan(anchoreRequestInfo)

		// get anchore scan data
		vulnRequestURL := model.Config.Anchore.AnchoreURL + "/v1/images/" + anchoreRequestInfo.ImageDigest + "/vuln/all"
		vulnData := AnchoreGET(vulnRequestURL)

		// preper anchore scan result data
		jobAnchoreInfo := PreperAnchoreScanResult(jobScannerMsg, vulnData)

		// send data to scancenter
		requestJSON, _ := json.Marshal(jobAnchoreInfo)
		_ = net.HTTPPUT(model.Config.ScanCenter.AnchoreURL, string(requestJSON))
	}
}

func TriggerAnchoreScan(anchoreRequestInfo *model.AnchoreRequestInfo) {
	triggerRequest, _ := json.Marshal(anchoreRequestInfo)
	triggerURL := model.Config.Anchore.AnchoreURL + "/v1/images"
RETRY:
	anchoreData := AnchorePOST(triggerURL, string(triggerRequest))
	analysisStatus := anchoreData[0]["analysis_status"].(string)
	if analysisStatus != "analyzed" {
		time.Sleep(3 * time.Second)
		goto RETRY
	}
}

func PreperAnchoreScanResult(jobScannerMsg *model.JobScannerMsg, vulnData map[string]interface{}) *model.JobScannerInfo {
	var cveList []map[string]string
	for _, vulnInfo := range vulnData["vulnerabilities"].([]interface{}) {
		cve := make(map[string]string)
		cve["package_full_nale"] = vulnInfo.(map[string]interface{})["package"].(string)
		cve["package_name"] = vulnInfo.(map[string]interface{})["package_name"].(string)
		cve["package_version"] = vulnInfo.(map[string]interface{})["package_version"].(string)
		cve["cve"] = vulnInfo.(map[string]interface{})["vuln"].(string)
		cve["cve_url"] = vulnInfo.(map[string]interface{})["url"].(string)
		cveList = append(cveList, cve)
	}
	jobScannerInfo := new(model.JobScannerInfo)
	jobScannerInfo.TaskID = jobScannerMsg.TaskID
	jobScannerInfo.JobID = jobScannerMsg.JobID
	jobScannerInfo.JobStatus = model.StatusSucceed
	jobScannerInfo.JobData = cveList
	return jobScannerInfo
}
