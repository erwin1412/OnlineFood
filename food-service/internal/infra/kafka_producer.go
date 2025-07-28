package infra

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

func NewKafkaProducer(brokerAddress string, topic string) *KafkaProducer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
		Async:        false,
	}

	return &KafkaProducer{Writer: w}
}

func (p *KafkaProducer) Publish(message interface{}) error {
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return p.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(time.Now().Format(time.RFC3339)),
			Value: value,
		},
	)
}

func (p *KafkaProducer) Close() error {
	return p.Writer.Close()
}
