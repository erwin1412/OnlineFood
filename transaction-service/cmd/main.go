package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"transaction-service/internal/app"
	"transaction-service/internal/config"
	cartGrpc "transaction-service/internal/delivery/grpc/cart"
	cartPb "transaction-service/internal/delivery/grpc/pb/cart"
	transactionPb "transaction-service/internal/delivery/grpc/pb/transaction"
	transactionGrpc "transaction-service/internal/delivery/grpc/transaction"
	"transaction-service/internal/infra"
	"transaction-service/pkg/payments"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// 1️⃣ Load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 2️⃣ Get Midtrans server key from environment
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		log.Fatal("MIDTRANS_SERVER_KEY environment variable is not set")
	}
	isProduction := os.Getenv("MIDTRANS_ENV") == "production" // Fixed: true only for production

	// 3️⃣ Initialize Midtrans client
	midtransClient := payments.NewMidtransClient(serverKey, isProduction)

	// 4️⃣ Init DB
	db := config.PostgresInit()

	// 5️⃣ Init repos
	cartRepo := infra.NewPgCartRepository(db)
	transactionRepo := infra.NewPgTransactionRepository(db)
	transactionDetailRepo := infra.NewPgTransactionDetailRepository(transactionRepo)

	// 6️⃣ Init apps
	cartApp := app.NewCartApp(cartRepo)
	transactionApp := app.NewTransactionApp(transactionRepo, transactionDetailRepo, cartRepo, midtransClient)

	// 7️⃣ Init handlers
	cartHandler := cartGrpc.NewCartHandler(cartApp)
	transactionHandler := transactionGrpc.NewTransactionHandler(transactionApp)

	// 8️⃣ gRPC server
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50055" // default port transaction-service
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// ✅ Register multiple services
	cartPb.RegisterCartServiceServer(grpcServer, cartHandler)
	transactionPb.RegisterTransactionServiceServer(grpcServer, transactionHandler)

	fmt.Println("🚀 gRPC Transaction Service running at :" + grpcPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// 9️⃣ Clean shutdown
	log.Println("Gracefully shutting down...")
	db.Close()
}
