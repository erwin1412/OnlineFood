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

	// 1. Load env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db := config.PostgresInit()
	foodRepo := infra.NewPgFoodRepository(db)

	producer := infra.NewKafkaProducer("localhost:9092", "food-created")

	authApp := app.NewFoodApp(foodRepo, producer)

	// gRPC handler
	authGRPC := grpcHandler.NewFoodHandler(authApp)
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50052"
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFoodServiceServer(grpcServer, authGRPC)

	fmt.Println("🚀 gRPC running at : " + grpcPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Gracefully shutting down...")
	db.Close()
}
