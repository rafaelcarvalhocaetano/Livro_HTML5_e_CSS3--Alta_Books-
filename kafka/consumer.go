package kafka

import (
	"fmt"

	kfk "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	Topics    []string
	ConfigMap *kfk.ConfigMap
}

func NewConsumer(cm *kfk.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		Topics:    topics,
		ConfigMap: cm,
	}
}

func (c *Consumer) Consume(msgChannel chan *kfk.Message) error {
	consumer, err := kfk.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			msgChannel <- msg
		}
		fmt.Println("Error: ", err)
	}
}
