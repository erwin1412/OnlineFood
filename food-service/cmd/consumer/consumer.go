package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
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
