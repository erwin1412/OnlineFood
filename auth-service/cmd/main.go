package main

import (
	"auth-service/internal/app"
	"auth-service/internal/config"
	grpcHandler "auth-service/internal/delivery/grpc"
	"auth-service/internal/delivery/grpc/pb"
	"auth-service/internal/infra"
	"auth-service/pkg/hasher"
	"auth-service/pkg/jwt"
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
	log.Println("✅ GOOGLE_API_KEY =", os.Getenv("GOOGLE_API_KEY"))

	config.InitSMTP()
	// log.Println("SMTP CONFIG:", config.SMTP)

	db := config.PostgresInit()
	userRepo := infra.NewPgAuthRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set in .env")
	}
	jwtManager := jwt.NewManager(jwtSecret)
	passwordHasher := hasher.NewBcrypt()

	// producer := infra.NewKafkaProducer("localhost:9092", "user-registered")

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		log.Println("⚠️  KAFKA_BROKER not set, using default localhost:9092")
		kafkaBroker = "localhost:9092"
	}
	producer := infra.NewKafkaProducer(kafkaBroker, "user-registered")

	authApp := app.NewAuthApp(userRepo, passwordHasher, jwtManager, producer)

	// gRPC handler
	authGRPC := grpcHandler.NewAuthHandler(authApp)
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authGRPC)

	fmt.Println("gRPC running at : " + grpcPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Gracefully shutting down...")
	db.Close()
}
