package scanner

import (
	"encoding/json"
	"errors"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
	"log"
	"time"

	"github.com/q8s-io/heimdall/pkg/entity/convert"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
	"github.com/q8s-io/heimdall/pkg/infrastructure/net"
	"github.com/q8s-io/heimdall/pkg/repository"
)

func JobAnchore(scanTime int) {
	// consumer msg from mq
	repository.ConsumerMsgJobAnchore()
	jobScannerMsg := new(model.JobScannerMsg)

	for msg := range kafka.Queue {
		log.Printf("consumer msg from kafka %s", msg)
		_ = json.Unmarshal(msg, &jobScannerMsg)

		// prepare anchore data
		anchoreRequestInfo := convert.AnchoreRequestInfo(jobScannerMsg)

		// trigger anchore scan
		jobAnchoreInfo := new(model.JobScannerInfo)
		timeoutStatus, scanError := TriggerAnchoreScan(anchoreRequestInfo, scanTime)
		if scanError != nil {
			xray.ErrTaskInfo(scanError, jobScannerMsg.TaskID, jobScannerMsg.JobID)
			jobAnchoreInfo = PrepareAnchoreScanResult(jobScannerMsg, nil, scanError)
		} else {
			// handle timeout problem
			if timeoutStatus {
				jobAnchoreInfo = PrepareAnchoreScanResult(jobScannerMsg, nil, nil)
				jobAnchoreInfo.JobStatus = model.StatusTimeout
			} else {
				// get anchore scan data
				vulnRequestURL := model.Config.Anchore.AnchoreURL + "/v1/images/" + anchoreRequestInfo.ImageDigest + "/vuln/all"
				vulnData, getErr := AnchoreGET(vulnRequestURL)
				if getErr != nil {
					xray.ErrTaskInfo(getErr, jobScannerMsg.TaskID, jobScannerMsg.JobID)
				}
				// prepare anchore scan result data
				jobAnchoreInfo = PrepareAnchoreScanResult(jobScannerMsg, vulnData, getErr)
			}
		}
		// send data to scancenter
		requestJSON, _ := json.Marshal(jobAnchoreInfo)
		log.Printf("anchore process  %s \t %s", jobAnchoreInfo.ImageName, jobAnchoreInfo.JobStatus)
		_ = net.HTTPPUT(model.Config.ScanCenter.AnchoreURL, string(requestJSON))
	}
}

// handle timout result
// return true represent timeout
// return false represent not timeout
func TriggerAnchoreScan(anchoreRequestInfo *model.AnchoreRequestInfo, scanTime int) (bool, error) {
	triggerRequest, _ := json.Marshal(anchoreRequestInfo)
	triggerURL := model.Config.Anchore.AnchoreURL + "/v1/images"
	startTime := time.Now()
RETRY:
	// handle timeout
	if time.Since(startTime) > time.Minute*time.Duration(scanTime) {
		return true, nil
	}
	anchoreData := AnchorePOST(triggerURL, string(triggerRequest))
	if len(anchoreData) == 0 {
		return false, xray.ErrMiniInfo(errors.New("anchoreData len is 0"))
	}
	analysisStatus, ok := anchoreData[0]["analysis_status"].(string)
	if !ok {
		return false, xray.ErrMiniInfo(errors.New("alert status error"))
	}
	if analysisStatus != "analyzed" {
		time.Sleep(3 * time.Second)
		goto RETRY
	}
	return false, nil
}

func PrepareAnchoreScanResult(jobScannerMsg *model.JobScannerMsg, vulnData map[string]interface{}, runErr error) *model.JobScannerInfo {
	var cveList []map[string]string
	if vulnData != nil && vulnData["vulnerabilities"] != nil {
		for _, vulnInfo := range vulnData["vulnerabilities"].([]interface{}) {
			cve := make(map[string]string)
			cve["package_full_nale"] = vulnInfo.(map[string]interface{})["package"].(string)
			cve["package_name"] = vulnInfo.(map[string]interface{})["package_name"].(string)
			cve["package_version"] = vulnInfo.(map[string]interface{})["package_version"].(string)
			cve["cve"] = vulnInfo.(map[string]interface{})["vuln"].(string)
			cve["cve_url"] = vulnInfo.(map[string]interface{})["url"].(string)
			cveList = append(cveList, cve)
		}
	}
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
