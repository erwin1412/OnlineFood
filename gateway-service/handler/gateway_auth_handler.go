package handler

import (
	"context"
	"gateway-service/config"
	authpb "gateway-service/pb/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GatewayAuthHandler struct {
	GRPC *config.GRPCClients
}

func NewGatewayAuthHandler(grpcClients *config.GRPCClients) *GatewayAuthHandler {
	return &GatewayAuthHandler{GRPC: grpcClients}
}

// auth service start
func (h *GatewayAuthHandler) Register(c echo.Context) error {
	var req authpb.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.AuthClient.Register(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h *GatewayAuthHandler) Login(c echo.Context) error {
	var req authpb.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.AuthClient.Login(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

// auth service end
