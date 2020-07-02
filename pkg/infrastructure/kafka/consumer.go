package kafka

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Shopify/sarama"

	"github.com/q8s-io/heimdall/pkg/entity/model"
)

type Consumer struct {
	ready chan bool
}

var Client sarama.ConsumerGroup
var clientErr interface{}

var Queue chan []byte

func InitConsumer() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange

	kafkaConfig := model.Config.Kafka

	Client, clientErr = sarama.NewConsumerGroup(kafkaConfig.BrokerList, "heimdall", config)
	if clientErr != nil {
		log.Println(clientErr)
	}

	Queue = make(chan []byte, 1)
}

func ConsumerMsg(topic string) {
	ctx, cancel := context.WithCancel(context.Background())
	consumer := Consumer{
		ready: make(chan bool),
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			err := Client.Consume(ctx, []string{topic}, &consumer)
			if err != nil {
				log.Println(err)
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("context cancelled")
	case <-sigterm:
		log.Println("signal cancelled")
		time.Sleep(2 * time.Second)
		close(Queue)
		os.Exit(0)
	}
	cancel()
	wg.Wait()
	err := Client.Close()
	if err != nil {
		log.Println(err)
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(c.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		Queue <- message.Value
		session.MarkMessage(message, "")
	}
	return nil
}
