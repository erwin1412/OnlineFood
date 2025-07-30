package handler

import (
	"context"
	"gateway-service/config"
	cartpb "gateway-service/pb/cart"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GatewayCartHandler struct {
	GRPC *config.GRPCClients
}

func NewGatewayCartHandler(grpcClients *config.GRPCClients) *GatewayCartHandler {
	return &GatewayCartHandler{GRPC: grpcClients}
}

// cart
func (h *GatewayCartHandler) CreateCart(c echo.Context) error {
	var req cartpb.CreateCartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		return err
	}
	req.UserId = userID
	resp, err := h.GRPC.CartClient.CreateCart(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, resp)
}
func (h *GatewayCartHandler) GetAllCart(c echo.Context) error {
	var req cartpb.GetAllCartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		return err
	}
	req.UserId = userID
	resp, err := h.GRPC.CartClient.GetAllCart(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayCartHandler) GetByIdCart(c echo.Context) error {
	var req cartpb.GetByIdCartRequest
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
	resp, err := h.GRPC.CartClient.GetByIdCart(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayCartHandler) UpdateCart(c echo.Context) error {
	var req cartpb.UpdateCartRequest
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
	resp, err := h.GRPC.CartClient.UpdateCart(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayCartHandler) DeleteCart(c echo.Context) error {
	var req cartpb.DeleteCartRequest
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
	resp, err := h.GRPC.CartClient.DeleteCart(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayCartHandler) DeleteAllCart(c echo.Context) error {
	var req cartpb.DeleteAllCartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		return err
	}
	req.UserId = userID
	resp, err := h.GRPC.CartClient.DeleteAllCart(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
