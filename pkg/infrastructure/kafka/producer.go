package kafka

import (
	"github.com/Shopify/sarama"

	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
)

var SyncProducer sarama.SyncProducer
var AsyncProducer sarama.AsyncProducer
var syncProducerErr error
var asyncProducerErr error

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

func InitAsyncProducer() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	kafkaConfig := model.Config.Kafka

	AsyncProducer, asyncProducerErr = sarama.NewAsyncProducer(kafkaConfig.BrokerList, config)
	if asyncProducerErr != nil {
		xray.ErrMini(asyncProducerErr)
		panic(asyncProducerErr)
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

func AsyncProducerSendMsg(topic string, message sarama.Encoder) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: message,
	}
	AsyncProducer.Input() <- msg
}
