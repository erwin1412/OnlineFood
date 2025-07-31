package main

import (
	"auth-service/internal/config" // ⬅️ Tambah ini!
	"auth-service/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func main() {
	// Load .env kalau ada
	_ = godotenv.Load()

	// PENTING: Init SMTP config dulu
	config.InitSMTP()
	// log.Println("SMTP CONFIG:", config.SMTP)
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		log.Println("⚠️  KAFKA_BROKER not set, using default localhost:9092")
		kafkaBroker = "localhost:9092"
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   "user-registered",
		GroupID: "auth-consumer-group",
	})

	fmt.Println("Consumer started...")

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

		email, _ := event["email"].(string)
		name, _ := event["name"].(string)

		subject := "Welcome!"
		body := fmt.Sprintf("Hello %s, thank you for registering!", name)

		if err := utils.SendMail(email, subject, body); err != nil {
			log.Printf("failed to send email: %v", err)
		} else {
			log.Printf("✅ Email sent to %s", email)
		}
	}
}
