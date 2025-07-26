package main

import (
	"fmt"
	"log"
	"merchant-service/internal/app"
	"merchant-service/internal/config"
	grpcHandler "merchant-service/internal/delivery/grpc"
	"merchant-service/internal/delivery/grpc/pb"
	"merchant-service/internal/infra"
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
	merchantRepo := infra.NewPgMerchantRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set in .env")
	}
	merchantApp := app.NewMerchantApp(merchantRepo)

	// gRPC handler
	merchantGRPC := grpcHandler.NewMerchantHandler(merchantApp)
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50052"
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMerchantServiceServer(grpcServer, merchantGRPC)

	fmt.Println("🚀 gRPC running at : " + grpcPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Gracefully shutting down...")
	db.Close()
}
