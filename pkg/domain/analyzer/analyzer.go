package analyzer

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"

	"github.com/q8s-io/heimdall/pkg/service"
)

func JobAnalyzer() {
	var queue chan *sarama.ConsumerMessage
	queue = make(chan *sarama.ConsumerMessage, 100)

	service.ConsumerImageAnalyzerMsg(queue)

	for msg := range queue {
		log.Println(msg.Topic, msg.Partition, msg.Offset)
		msgInfo := make(map[string]interface{})
		_ = json.Unmarshal(msg.Value, &msgInfo)
		log.Println(msgInfo["cluster"], msgInfo["app"], msgInfo["controller_kind"], msgInfo["deployment"], msgInfo["container_name"], msgInfo["hostname"])
		log.Println(msgInfo["log"])
	}

	close(queue)
}
