package config

import (
	"log"
	"os"

	authpb "gateway-service/pb"
	cartpb "gateway-service/pb"
	courierpb "gateway-service/pb"
	foodpb "gateway-service/pb"
	merchantpb "gateway-service/pb"
	transactionpb "gateway-service/pb"

	"google.golang.org/grpc"
)

type GRPCClients struct {
	AuthClient        authpb.AuthServiceClient
	CourierClient     courierpb.CourierServiceClient
	FoodClient        foodpb.FoodServiceClient
	MerchantClient    merchantpb.MerchantServiceClient
	CartClient        cartpb.CartServiceClient
	TransactionClient transactionpb.TransactionServiceClient
}

func NewGRPCClients() *GRPCClients {
	authConn, err := grpc.Dial(os.Getenv("AUTH_SERVICE_URL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Auth Service: %v", err)
	}
	courierConn, err := grpc.Dial(os.Getenv("COURIER_SERVICE_URL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Courier Service: %v", err)
	}
	foodConn, err := grpc.Dial(os.Getenv("FOOD_SERVICE_URL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Food Service: %v", err)
	}
	merchantConn, err := grpc.Dial(os.Getenv("MERCHANT_SERVICE_URL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Merchant Service: %v", err)
	}
	transactionConn, err := grpc.Dial(os.Getenv("TRANSACTION_SERVICE_URL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Transaction Service: %v", err)
	}
	return &GRPCClients{
		AuthClient:        authpb.NewAuthServiceClient(authConn),
		CourierClient:     courierpb.NewCourierServiceClient(courierConn),
		FoodClient:        foodpb.NewFoodServiceClient(foodConn),
		MerchantClient:    merchantpb.NewMerchantServiceClient(merchantConn),
		CartClient:        cartpb.NewCartServiceClient(transactionConn),
		TransactionClient: transactionpb.NewTransactionServiceClient(transactionConn),
	}
}
