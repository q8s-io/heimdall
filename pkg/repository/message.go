package repository

import (
	"github.com/Shopify/sarama"
	
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
)

func SendMsgJobImageAnalyzer(jobScannerMsg *model.JobScannerMsg) {
	kafka.SyncProducerSendMsg("analyzer", jobScannerMsg)
}

func SendMsgJobAnchore(jobScannerMsg *model.JobScannerMsg) {
	kafka.SyncProducerSendMsg("anchore", jobScannerMsg)
}

func ConsumerMsgJobImageAnalyzer(queue chan *sarama.ConsumerMessage) {
	kafka.ConsumerMsg("analyzer", queue)
}

func ConsumerMsgJobAnchore(queue chan *sarama.ConsumerMessage) {
	kafka.ConsumerMsg("anchore", queue)
}
