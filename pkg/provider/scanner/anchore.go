package scanner

import (
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
	
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/infrastructure/net"
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
)

func JobAnchore() {
	var queue chan *sarama.ConsumerMessage
	queue = make(chan *sarama.ConsumerMessage, 1000)

	// consumer msg from mq
	persistence.ConsumerJobAnchoreMsg(queue)
	jobAnchoreMsg := new(entity.JobAnchoreMsg)
	for msg := range queue {
		_ = json.Unmarshal(msg.Value, &jobAnchoreMsg)

		// preper anchore data
		anchoreRequestInfo := CreateAnchoreRequestInfo(jobAnchoreMsg)

		// trigger anchore scan
		TriggerAnchoreScan(anchoreRequestInfo)

		// get anchore scan data
		vulnURL := entity.Config.Anchore.AnchoreURL + "/v1/images/" + anchoreRequestInfo.ImageDigest + "/vuln/all"
		vulnData := AnchoreGET(vulnURL)

		// preper anchore scan result data
		jobAnchoreInfo := PreperAnchoreScanResult(jobAnchoreMsg, vulnData)

		// send data to scancenter
		requestJSON, _ := json.Marshal(jobAnchoreInfo)
		_ = net.HTTPPUT(entity.Config.ScanCenter.AnchoreURL, string(requestJSON))
	}

	close(queue)
}

func TriggerAnchoreScan(anchoreRequestInfo *entity.AnchoreRequestInfo) {
	triggerRequest, _ := json.Marshal(anchoreRequestInfo)
	triggerURL := entity.Config.Anchore.AnchoreURL + "/v1/images"
RETYR:
	anchoreData := AnchorePOST(triggerURL, string(triggerRequest))
	analysisStatus := anchoreData[0]["analysis_status"].(string)
	if analysisStatus != "analyzed" {
		time.Sleep(3 * time.Second)
		goto RETYR
	}
}

func PreperAnchoreScanResult(jobAnchoreMsg *entity.JobAnchoreMsg, vulnData map[string]interface{}) *entity.JobAnchoreInfo {
	var cveList []map[string]string
	for _, vulnInfo := range vulnData["vulnerabilities"].([]interface{}) {
		anchoreScanInfo := make(map[string]string)
		anchoreScanInfo["package_full_nale"] = vulnInfo.(map[string]interface{})["package"].(string)
		anchoreScanInfo["package_name"] = vulnInfo.(map[string]interface{})["package_name"].(string)
		anchoreScanInfo["package_version"] = vulnInfo.(map[string]interface{})["package_version"].(string)
		anchoreScanInfo["cve"] = vulnInfo.(map[string]interface{})["vuln"].(string)
		anchoreScanInfo["cve_url"] = vulnInfo.(map[string]interface{})["url"].(string)
		cveList = append(cveList, anchoreScanInfo)
	}
	jobAnchoreInfo := new(entity.JobAnchoreInfo)
	jobAnchoreInfo.TaskID = jobAnchoreMsg.TaskID
	jobAnchoreInfo.JobID = jobAnchoreMsg.JobID
	jobAnchoreInfo.JobStatus = entity.StatusSucceed
	jobAnchoreInfo.JobData = cveList
	return jobAnchoreInfo
}
