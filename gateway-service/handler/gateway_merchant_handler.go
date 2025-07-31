package handler

import (
	"context"
	"gateway-service/config"
	merchantpb "gateway-service/pb/merchant"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GatewayMerchantHandler struct {
	GRPC *config.GRPCClients
}

func NewGatewayMerchantHandler(grpcClients *config.GRPCClients) *GatewayMerchantHandler {
	return &GatewayMerchantHandler{GRPC: grpcClients}
}

// merchant service start
func (h *GatewayMerchantHandler) GetByIdMerchant(c echo.Context) error {
	var req merchantpb.GetByIdMerchantRequest
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "courier ID cannot be empty"})
	}
	req.Id = id
	resp, err := h.GRPC.MerchantClient.GetByIdMerchant(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayMerchantHandler) CreateMerchant(c echo.Context) error {
	var req merchantpb.CreateMerchantRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	userID, err := GetUserIDFromContext(c)
	if err != nil {
		return err
	}
	req.UserId = userID
	resp, err := h.GRPC.MerchantClient.CreateMerchant(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, resp)
}
func (h *GatewayMerchantHandler) GetAllMerchant(c echo.Context) error {
	var req merchantpb.Empty
	resp, err := h.GRPC.MerchantClient.GetAllMerchant(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayMerchantHandler) UpdateMerchant(c echo.Context) error {
	var req merchantpb.UpdateMerchantRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "courier ID cannot be empty"})
	}
	req.Id = id
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		return err
	}
	req.UserId = userID
	resp, err := h.GRPC.MerchantClient.UpdateMerchant(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayMerchantHandler) DeleteMerchant(c echo.Context) error {
	var req merchantpb.DeleteMerchantRequest
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "courier ID cannot be empty"})
	}
	req.Id = id
	resp, err := h.GRPC.MerchantClient.DeleteMerchant(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayMerchantHandler) GetMerchantByUserId(c echo.Context) error {
	var req merchantpb.GetMerchantByUserIdRequest
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		return err
	}
	req.UserId = userID
	resp, err := h.GRPC.MerchantClient.GetMerchantByUserId(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

// merchant service end
