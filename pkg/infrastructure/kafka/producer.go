package kafka

import (
	"github.com/Shopify/sarama"

	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
)

var SyncProducer sarama.SyncProducer
var syncProducerErr error

func InitSyncProducer() {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 1
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	kafkaConfig := model.Config.Kafka

	SyncProducer, syncProducerErr = sarama.NewSyncProducer(kafkaConfig.BrokerList, config)
	if syncProducerErr != nil {
		xray.ErrMini(syncProducerErr)
	}
}

func SyncProducerSendMsg(topic string, message sarama.Encoder) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: message,
	}
	_, _, err := SyncProducer.SendMessage(msg)

	if err != nil {
		xray.ErrMini(err)
	}
}
