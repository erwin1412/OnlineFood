package handler

import (
	"context"
	"gateway-service/config"
	transactionpb "gateway-service/pb/transaction"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GatewayTransactionHandler struct {
	GRPC *config.GRPCClients
}

func NewGatewayTransactionHandler(grpcClients *config.GRPCClients) *GatewayTransactionHandler {
	return &GatewayTransactionHandler{GRPC: grpcClients}
}

// transaction
func (h *GatewayTransactionHandler) CreateTransaction(c echo.Context) error {
	var req transactionpb.CreateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		return err
	}
	req.UserId = userID
	resp, err := h.GRPC.TransactionClient.CreateTransaction(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h *GatewayTransactionHandler) GetAllTransaction(c echo.Context) error {
	var req transactionpb.GetAllTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		return err
	}
	req.UserId = userID
	resp, err := h.GRPC.TransactionClient.GetAllTransaction(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayTransactionHandler) GetByIdTransaction(c echo.Context) error {
	var req transactionpb.GetByIdTransactionRequest
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
	resp, err := h.GRPC.TransactionClient.GetByIdTransaction(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayTransactionHandler) UpdateTransaction(c echo.Context) error {
	var req transactionpb.UpdateTransactionRequest
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
	resp, err := h.GRPC.TransactionClient.UpdateTransaction(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayTransactionHandler) DeleteTransaction(c echo.Context) error {
	var req transactionpb.DeleteTransactionRequest
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
	resp, err := h.GRPC.TransactionClient.DeleteTransaction(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

// transaction service end
