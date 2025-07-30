package main

import (
	"courier-service/internal/app"
	"courier-service/internal/config"
	grpcHandler "courier-service/internal/delivery/grpc"
	"courier-service/internal/delivery/grpc/pb"
	"courier-service/internal/infra"
	"fmt"
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
	courierRepo := infra.NewPgCourierRepository(db)

	courierApp := app.NewCourierApp(courierRepo)

	// gRPC handler
	courierGRPC := grpcHandler.NewCourierHandler(courierApp)
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50052"
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCourierServiceServer(grpcServer, courierGRPC)

	fmt.Println("🚀 gRPC running at : " + grpcPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Gracefully shutting down...")
	db.Close()
}
