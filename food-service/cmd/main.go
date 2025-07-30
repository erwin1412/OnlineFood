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
	producer := infra.NewKafkaProducer("localhost:9092", "food-created")

	// Injek ke FoodApp
	foodApp := app.NewFoodApp(foodRepo, producer)

	// Init handler gRPC
	foodGRPC := grpcHandler.NewFoodHandler(foodApp)

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50052"
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
