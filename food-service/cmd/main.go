package main

import (
	"fmt"
	"food-service/internal/app"
	"food-service/internal/config"
	grpcHandler "food-service/internal/delivery/grpc"
	"food-service/internal/delivery/grpc/pb"
	"food-service/internal/infra"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load .env kalau ada
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Init DB
	db := config.PostgresInit()
	foodRepo := infra.NewPgFoodRepository(db)

	// Init Kafka producer (harus nyala Zookeeper & Kafka broker!)
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		log.Println("⚠️  KAFKA_BROKER not set, using default localhost:9092")
		kafkaBroker = "localhost:9092"
	}
	producer := infra.NewKafkaProducer(kafkaBroker, "food-created")

	merchantConn, err := grpc.Dial(os.Getenv("MERCHANT_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to Merchant Service: %v", err)
	}
	defer merchantConn.Close()

	merchantClient := infra.NewMerchantClient(merchantConn)

	// Injek ke FoodApp
	foodApp := app.NewFoodApp(foodRepo, producer, merchantClient)

	// Init handler gRPC

	foodGRPC := grpcHandler.NewFoodHandler(foodApp)

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50053"
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFoodServiceServer(grpcServer, foodGRPC)

	fmt.Println("🚀 gRPC running at :", grpcPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Tutup DB connection kalau server stop
	defer db.Close()
}
