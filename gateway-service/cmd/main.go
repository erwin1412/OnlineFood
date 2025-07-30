package main

import (
	"gateway-service/config"
	"gateway-service/handler"
	"gateway-service/middleware"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	config.LoadEnv()

	grpcClients := config.NewGRPCClients()
	ah := handler.NewGatewayAuthHandler(grpcClients)
	crh := handler.NewGatewayCourierHandler(grpcClients)
	fh := handler.NewGatewayFoodHandler(grpcClients)
	mh := handler.NewGatewayMerchantHandler(grpcClients)
	cth := handler.NewGatewayCartHandler(grpcClients)
	th := handler.NewGatewayTransactionHandler(grpcClients)

	e := echo.New()

	// AUTH endpoints (proxy ke Auth Service via gRPC)
	e.POST("/register", ah.Register)
	e.POST("/login", ah.Login)

	// COURIER endpoints (proxy ke Courier Service via gRPC, protected by JWT)
	courierGroup := e.Group("/couriers", middleware.JWTAuth)
	courierGroup.GET("", crh.GetAllCouriers)
	courierGroup.GET("/:id", crh.GetByIdCourier)
	courierGroup.GET("/longlat/:lat/:long", crh.GetByLongLatCourier)
	courierGroup.POST("", crh.CreateCourier)
	courierGroup.PUT("/:id", crh.UpdateLongLatCourier)
	courierGroup.DELETE("/:id", crh.DeleteCourier)
	courierGroup.GET("/near/:lat/:long", crh.FindNearestCourier)

	// FOOD endpoints (proxy ke Food Service via gRPC, protected by JWT)
	foodGroup := e.Group("/foods", middleware.JWTAuth)
	foodGroup.GET("/:id", fh.GetByIdFood)
	foodGroup.POST("", fh.CreateFood)
	foodGroup.GET("", fh.GetAllFood)
	foodGroup.PUT("/:id", fh.UpdateFood)
	foodGroup.DELETE("/:id", fh.DeleteFood)

	// MERCHANT endpoints (proxy ke Merchant Service via gRPC, protected by JWT)
	merchantGroup := e.Group("/merchants", middleware.JWTAuth)
	merchantGroup.GET("/:id", mh.GetByIdMerchant)
	merchantGroup.POST("", mh.CreateMerchant)
	merchantGroup.GET("", mh.GetAllMerchant)
	merchantGroup.PUT("/:id", mh.UpdateMerchant)
	merchantGroup.DELETE("/:id", mh.DeleteMerchant)

	// CART endpoints (proxy ke Cart Service via gRPC, protected by JWT)
	cartGroup := e.Group("/carts", middleware.JWTAuth)
	cartGroup.POST("", cth.CreateCart)
	cartGroup.GET("", cth.GetAllCart)
	cartGroup.GET("/:id", cth.GetByIdCart)
	cartGroup.PUT("/:id", cth.UpdateCart)
	cartGroup.DELETE("/:id", cth.DeleteCart)
	cartGroup.DELETE("", cth.DeleteAllCart)

	// TRANSACTION endpoints (proxy ke Transaction Service via gRPC, protected by JWT)
	transactionGroup := e.Group("/transactions", middleware.JWTAuth)
	transactionGroup.POST("", th.CreateTransaction)
	transactionGroup.GET("", th.GetAllTransaction)
	transactionGroup.GET("/:id", th.GetByIdTransaction)
	transactionGroup.PUT("/:id", th.UpdateTransaction)
	transactionGroup.DELETE("/:id", th.DeleteTransaction)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8032"
	}
	log.Println("Gateway Service running at http://localhost:" + port)
	e.Logger.Fatal(e.Start(":" + port))
}
