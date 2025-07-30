package handler

import (
	"context"
	"gateway-service/config"
	courierpb "gateway-service/pb/courier"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GatewayCourierHandler struct {
	GRPC *config.GRPCClients
}

func NewGatewayCourierHandler(grpcClients *config.GRPCClients) *GatewayCourierHandler {
	return &GatewayCourierHandler{GRPC: grpcClients}
}

// courier service start
func (h *GatewayCourierHandler) GetByIdCourier(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "courier ID cannot be empty"})
	}
	var req courierpb.GetByIdCourierRequest
	req.Id = id

	resp, err := h.GRPC.CourierClient.GetByIdCourier(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayCourierHandler) GetByLongLatCourier(c echo.Context) error {
	lat := c.Param("lat")
	long := c.Param("long")

	req := courierpb.GetByLongLatCourierRequest{
		Lat:  lat,
		Long: long,
	}

	resp, err := h.GRPC.CourierClient.GetByLongLatCourier(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayCourierHandler) CreateCourier(c echo.Context) error {
	var req courierpb.CreateCourierRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		return err
	}
	req.UserId = userID
	resp, err := h.GRPC.CourierClient.CreateCourier(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, resp)
}
func (h *GatewayCourierHandler) UpdateLongLatCourier(c echo.Context) error {
	var req courierpb.UpdateLongLatCourierRequest
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

	resp, err := h.GRPC.CourierClient.UpdateLongLatCourier(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayCourierHandler) DeleteCourier(c echo.Context) error {
	var req courierpb.DeleteCourierRequest
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "courier ID cannot be empty"})
	}
	req.Id = id
	resp, err := h.GRPC.CourierClient.DeleteCourier(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *GatewayCourierHandler) GetAllCouriers(c echo.Context) error {
	var req courierpb.Empty
	resp, err := h.GRPC.CourierClient.GetAllCouriers(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *GatewayCourierHandler) FindNearestCourier(c echo.Context) error {
	lat := c.Param("lat")
	long := c.Param("long")

	req := courierpb.FindNearestCourierRequest{
		Lat:  lat,
		Long: long,
	}
	resp, err := h.GRPC.CourierClient.FindNearestCourier(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

// courier service end
