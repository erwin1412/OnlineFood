package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		log.Println("⚠️  KAFKA_BROKER not set, using default localhost:9092")
		kafkaBroker = "localhost:9092"
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   "food-created",
		GroupID: "food-consumer-group",
	})

	fmt.Println("Food Consumer started...")

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		var event map[string]interface{}
		if err := json.Unmarshal(m.Value, &event); err != nil {
			log.Printf("failed to unmarshal: %v", err)
			continue
		}

		fmt.Printf("✅ New food created: %+v\n", event)
	}
}
