package kafka

import (
	"log"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"

	"github.com/q8s-io/heimdall/pkg/models"
)

var Consumer *cluster.Consumer
var consumerErr interface{}

func InitConsumer(topic string) {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = false
	config.Group.PartitionStrategy = "range"
	config.Group.Return.Notifications = false
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	kafkaConfig := models.Config.Kafka

	Consumer, consumerErr = cluster.NewConsumer(kafkaConfig.BrokerList, "devsecops", []string{topic}, config)
	if consumerErr != nil {
		log.Println(consumerErr)
	}
	defer Consumer.Close()

	// go func() {
	// 	for {
	// 		select {
	// 		case err, more := <-Consumer.Errors():
	// 			if more {
	// 				fmt.Println(err.Error())
	// 			}
	// 		case ntf, more := <-Consumer.Notifications():
	// 			if more {
	// 				fmt.Println(ntf)
	// 			}
	// 		}
	// 	}
	// }()
}

func ConsumerMsg(queue chan *sarama.ConsumerMessage) {
	for msg := range Consumer.Messages() {
		queue <- msg
		Consumer.MarkOffset(msg, "")
	}
}
