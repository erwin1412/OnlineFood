package handler

import (
	"context"
	"gateway-service/config"
	foodpb "gateway-service/pb/food"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GatewayFoodHandler struct {
	GRPC *config.GRPCClients
}

func NewGatewayFoodHandler(grpcClients *config.GRPCClients) *GatewayFoodHandler {
	return &GatewayFoodHandler{GRPC: grpcClients}
}

// food service start
func (h *GatewayFoodHandler) GetByIdFood(c echo.Context) error {
	var req foodpb.GetByIdFoodRequest
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "courier ID cannot be empty"})
	}
	req.Id = id
	resp, err := h.GRPC.FoodClient.GetByIdFood(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayFoodHandler) CreateFood(c echo.Context) error {
	var req foodpb.CreateFoodRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.FoodClient.CreateFood(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h *GatewayFoodHandler) GetAllFood(c echo.Context) error {
	var req foodpb.EmptyFood
	resp, err := h.GRPC.FoodClient.GetAllFood(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayFoodHandler) UpdateFood(c echo.Context) error {
	var req foodpb.UpdateFoodRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "courier ID cannot be empty"})
	}
	req.Id = id
	resp, err := h.GRPC.FoodClient.UpdateFood(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayFoodHandler) DeleteFood(c echo.Context) error {
	var req foodpb.DeleteFoodRequest
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "courier ID cannot be empty"})
	}
	req.Id = id
	resp, err := h.GRPC.FoodClient.DeleteFood(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

// food service end
