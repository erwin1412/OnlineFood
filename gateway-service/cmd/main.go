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
	h := handler.NewGatewayHandler(grpcClients)

	e := echo.New()

	// AUTH endpoints (proxy ke Auth Service via gRPC)
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)

	// COURIER endpoints (proxy ke Courier Service via gRPC, protected by JWT)
	courierGroup := e.Group("/couriers", middleware.JWTAuth)
	courierGroup.GET("", h.GetAllCouriers)
	courierGroup.GET("/:id", h.GetByIdCourier)
	courierGroup.GET("/longlat", h.GetByLongLatCourier)
	courierGroup.GET("/longlat/:lat/:long", h.GetByLongLatCourier)
	courierGroup.POST("", h.CreateCourier)
	courierGroup.PUT("/:id", h.UpdateLongLatCourier)
	courierGroup.DELETE("/:id", h.DeleteCourier)

	// FOOD endpoints (proxy ke Food Service via gRPC, protected by JWT)
	foodGroup := e.Group("/foods", middleware.JWTAuth)
	foodGroup.GET("/:id", h.GetByIdFood)
	foodGroup.POST("", h.CreateFood)
	foodGroup.GET("", h.GetAllFood)
	foodGroup.PUT("/:id", h.UpdateFood)
	foodGroup.DELETE("/:id", h.DeleteFood)

	// MERCHANT endpoints (proxy ke Merchant Service via gRPC, protected by JWT)
	merchantGroup := e.Group("/merchants", middleware.JWTAuth)
	merchantGroup.GET("/:id", h.GetByIdMerchant)
	merchantGroup.POST("", h.CreateMerchant)
	merchantGroup.GET("", h.GetAllMerchant)
	merchantGroup.PUT("/:id", h.UpdateMerchant)
	merchantGroup.DELETE("/:id", h.DeleteMerchant)

	// CART endpoints (proxy ke Cart Service via gRPC, protected by JWT)
	cartGroup := e.Group("/carts", middleware.JWTAuth)
	cartGroup.POST("", h.CreateCart)
	cartGroup.GET("", h.GetAllCart)
	cartGroup.GET("/:id", h.GetByIdCart)
	cartGroup.PUT("/:id", h.UpdateCart)
	cartGroup.DELETE("/:id", h.DeleteCart)
	cartGroup.DELETE("", h.DeleteAllCart)

	// TRANSACTION endpoints (proxy ke Transaction Service via gRPC, protected by JWT)
	transactionGroup := e.Group("/transactions", middleware.JWTAuth)
	transactionGroup.POST("", h.CreateTransaction)
	transactionGroup.GET("", h.GetAllTransaction)
	transactionGroup.GET("/:id", h.GetByIdTransaction)
	transactionGroup.PUT("/:id", h.UpdateTransaction)
	transactionGroup.DELETE("/:id", h.DeleteTransaction)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Gateway Service running at http://localhost:" + port)
	e.Logger.Fatal(e.Start(":" + port))
}
