package kafka

import (
	"log"

	"github.com/Shopify/sarama"

	"github.com/q8s-io/heimdall/pkg/models"
)

var Consumer sarama.Consumer
var consumerErr interface{}

func InitConsumer() {
	config := sarama.NewConfig()

	kafkaConfig := models.Config.Kafka

	Consumer, consumerErr = sarama.NewConsumer(kafkaConfig.BrokerList, config)
	if consumerErr != nil {
		log.Println("consumer error", consumerErr)
	}
}

func ConsumerMsg(topic string, queue chan *sarama.ConsumerMessage) {
	partitionList, err := Consumer.Partitions(topic)
	if err != nil {
		log.Println("fail to get list of partition", err)
		return
	}

	for partition := range partitionList {
		partitionConsumer, err := Consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Println(err)
		}
		go func() {
			for {
				select {
				case err := <-partitionConsumer.Errors():
					log.Println(err)
				case msg := <-partitionConsumer.Messages():
					queue <- msg
				}
			}
		}()
	}
}
