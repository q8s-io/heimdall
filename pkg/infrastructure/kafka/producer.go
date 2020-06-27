package kafka

import (
	"log"

	"github.com/Shopify/sarama"

	"github.com/q8s-io/heimdall/pkg/models"
)

var SyncProducer sarama.SyncProducer
var syncProducerErr interface{}

func InitSyncProducer() {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 1
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	kafkaConfig := models.Config.Kafka

	SyncProducer, syncProducerErr = sarama.NewSyncProducer(kafkaConfig.BrokerList, config)
	if syncProducerErr != nil {
		log.Println(syncProducerErr)
	}
	defer SyncProducer.Close()
}

func SyncProducerSendMsg(topic string, message sarama.Encoder) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: message,
	}
	partition, offset, err := SyncProducer.SendMessage(msg)

	if err != nil {
		log.Println("Error publish ", err.Error())
	}
	log.Println("partition", partition, " offset", offset)
}
