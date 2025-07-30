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

	// 2️⃣ Midtrans
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		log.Fatal("MIDTRANS_SERVER_KEY not set")
	}
	isProduction := os.Getenv("MIDTRANS_ENV") == "production"
	midtransClient := payments.NewMidtransClient(serverKey, isProduction)

	// 3️⃣ DB
	db := config.PostgresInit()

	// 4️⃣ Repos
	cartRepo := infra.NewPgCartRepository(db)
	transactionRepo := infra.NewPgTransactionRepository(db)
	transactionDetailRepo := infra.NewPgTransactionDetailRepository(transactionRepo)

	// 5️⃣ 🔗 gRPC connection to Courier Service
	courierConn, err := grpc.Dial(os.Getenv("COURIER_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to Courier Service: %v", err)
	}
	defer courierConn.Close()

	courierClient := infra.NewCourierClient(courierConn)

	// 6️⃣ 🔗 gRPC connection to Merchant Service
	merchantConn, err := grpc.Dial(os.Getenv("MERCHANT_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to Merchant Service: %v", err)
	}
	defer merchantConn.Close()

	merchantClient := infra.NewMerchantClient(merchantConn)

	// 7️⃣ Apps
	cartApp := app.NewCartApp(cartRepo)
	transactionApp := app.NewTransactionApp(
		transactionRepo,
		transactionDetailRepo,
		cartRepo,
		midtransClient,
		courierClient,
		merchantClient,
	)

	// 8️⃣ Handlers
	cartHandler := cartGrpc.NewCartHandler(cartApp)
	transactionHandler := transactionGrpc.NewTransactionHandler(transactionApp)

	// 9️⃣ gRPC server
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50055"
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	cartPb.RegisterCartServiceServer(grpcServer, cartHandler)
	transactionPb.RegisterTransactionServiceServer(grpcServer, transactionHandler)

	fmt.Println("🚀 gRPC Transaction Service running at :" + grpcPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Clean shutdown
	log.Println("Gracefully shutting down...")
	db.Close()
}
