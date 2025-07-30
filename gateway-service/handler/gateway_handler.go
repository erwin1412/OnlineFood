package handler

import (
	"context"
	"gateway-service/config"
	authpb "gateway-service/pb/auth"
	cartpb "gateway-service/pb/cart"
	courierpb "gateway-service/pb/courier"
	foodpb "gateway-service/pb/food"
	merchantpb "gateway-service/pb/merchant"
	transactionpb "gateway-service/pb/transaction"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GatewayHandler struct {
	GRPC *config.GRPCClients
}

func NewGatewayHandler(grpcClients *config.GRPCClients) *GatewayHandler {
	return &GatewayHandler{GRPC: grpcClients}
}

// auth service start
func (h *GatewayHandler) Register(c echo.Context) error {
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

func (h *GatewayHandler) Login(c echo.Context) error {
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

// courier service start
func (h *GatewayHandler) GetByIdCourier(c echo.Context) error {
	var req courierpb.GetByIdCourierRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.CourierClient.GetByIdCourier(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayHandler) GetByLongLatCourier(c echo.Context) error {
	var req courierpb.GetByLongLatCourierRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.CourierClient.GetByLongLatCourier(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayHandler) CreateCourier(c echo.Context) error {
	var req courierpb.CreateCourierRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.CourierClient.CreateCourier(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, resp)
}
func (h *GatewayHandler) UpdateLongLatCourier(c echo.Context) error {
	var req courierpb.UpdateLongLatCourierRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.CourierClient.UpdateLongLatCourier(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayHandler) DeleteCourier(c echo.Context) error {
	var req courierpb.DeleteCourierRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.CourierClient.DeleteCourier(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayHandler) GetAllCouriers(c echo.Context) error {
	var req courierpb.Empty
	resp, err := h.GRPC.CourierClient.GetAllCouriers(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayHandler) FindNearestCourier(c echo.Context) error {
	var req courierpb.FindNearestCourierRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.CourierClient.FindNearestCourier(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

// courier service end

// food service start
func (h *GatewayHandler) GetByIdFood(c echo.Context) error {
	var req foodpb.GetByIdFoodRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.FoodClient.GetByIdFood(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayHandler) CreateFood(c echo.Context) error {
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

func (h *GatewayHandler) GetAllFood(c echo.Context) error {
	var req foodpb.EmptyFood
	resp, err := h.GRPC.FoodClient.GetAllFood(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayHandler) UpdateFood(c echo.Context) error {
	var req foodpb.UpdateFoodRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.FoodClient.UpdateFood(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayHandler) DeleteFood(c echo.Context) error {
	var req foodpb.DeleteFoodRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.FoodClient.DeleteFood(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

// food service end

// merchant service start
func (h *GatewayHandler) GetByIdMerchant(c echo.Context) error {
	var req merchantpb.GetByIdMerchantRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.MerchantClient.GetByIdMerchant(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayHandler) CreateMerchant(c echo.Context) error {
	var req merchantpb.CreateMerchantRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.MerchantClient.CreateMerchant(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, resp)
}
func (h *GatewayHandler) GetAllMerchant(c echo.Context) error {
	var req merchantpb.EmptyMerchant
	resp, err := h.GRPC.MerchantClient.GetAllMerchant(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayHandler) UpdateMerchant(c echo.Context) error {
	var req merchantpb.UpdateMerchantRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.MerchantClient.UpdateMerchant(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayHandler) DeleteMerchant(c echo.Context) error {
	var req merchantpb.DeleteMerchantRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.MerchantClient.DeleteMerchant(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

// merchant service end

// transaction service start
// cart
func (h *GatewayHandler) CreateCart(c echo.Context) error {
	var req cartpb.CreateCartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.CartClient.CreateCart(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, resp)
}
func (h *GatewayHandler) GetAllCart(c echo.Context) error {
	var req cartpb.GetAllCartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.CartClient.GetAllCart(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayHandler) GetByIdCart(c echo.Context) error {
	var req cartpb.GetByIdCartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.CartClient.GetByIdCart(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayHandler) UpdateCart(c echo.Context) error {
	var req cartpb.UpdateCartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.CartClient.UpdateCart(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayHandler) DeleteCart(c echo.Context) error {
	var req cartpb.DeleteCartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.CartClient.DeleteCart(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayHandler) DeleteAllCart(c echo.Context) error {
	var req cartpb.DeleteAllCartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.CartClient.DeleteAllCart(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

// transaction
func (h *GatewayHandler) CreateTransaction(c echo.Context) error {
	var req transactionpb.CreateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.TransactionClient.CreateTransaction(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h *GatewayHandler) GetAllTransaction(c echo.Context) error {
	var req transactionpb.GetAllTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.TransactionClient.GetAllTransaction(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayHandler) GetByIdTransaction(c echo.Context) error {
	var req transactionpb.GetByIdTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.TransactionClient.GetByIdTransaction(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayHandler) UpdateTransaction(c echo.Context) error {
	var req transactionpb.UpdateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.TransactionClient.UpdateTransaction(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayHandler) DeleteTransaction(c echo.Context) error {
	var req transactionpb.DeleteTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	resp, err := h.GRPC.TransactionClient.DeleteTransaction(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

// transaction service end
