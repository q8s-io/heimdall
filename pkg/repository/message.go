package repository

import (
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
)

func SendMsgJobImageAnalyzer(jobScannerMsg *model.JobScannerMsg) {
	kafka.SyncProducerSendMsg("analyzer", jobScannerMsg)
}

func SendMsgJobAnchore(jobScannerMsg *model.JobScannerMsg) {
	kafka.SyncProducerSendMsg("anchore", jobScannerMsg)
}

func ConsumerMsgJobImageAnalyzer() {
	go kafka.ConsumerMsg("analyzer")
}

func ConsumerMsgJobAnchore() {
	go kafka.ConsumerMsg("anchore")
}
