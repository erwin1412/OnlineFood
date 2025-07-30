package grpc

import (
	"context"
	"courier-service/internal/app"
	pb "courier-service/internal/delivery/grpc/pb"
	"courier-service/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CourierHandler struct {
	pb.UnimplementedCourierServiceServer
	App *app.CourierApp
}

func NewCourierHandler(app *app.CourierApp) *CourierHandler {
	return &CourierHandler{App: app}
}

func (h *CourierHandler) GetByIdCourier(ctx context.Context, req *pb.GetByIdCourierRequest) (*pb.CourierResponse, error) {
	courier, err := h.App.GetById(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get courier by id: %v", err)
	}
	if courier == nil {
		return nil, status.Errorf(codes.NotFound, "courier not found")
	}
	return &pb.CourierResponse{
		Courier: &pb.Courier{
			Id:            courier.ID,
			UserId:        courier.UserID,
			Lat:           courier.Lat,
			Long:          courier.Long,
			VehicleNumber: courier.VehicleNumber,
			Status:        courier.Status,
			CreatedAt:     timestamppb.New(courier.CreatedAt),
			UpdatedAt:     timestamppb.New(courier.UpdatedAt),
		},
	}, nil
}
func (h *CourierHandler) CreateCourier(ctx context.Context, req *pb.CreateCourierRequest) (*pb.CourierResponse, error) {
	courier := &domain.Courier{
		UserID:        req.GetUserId(),
		Lat:           req.GetLat(),
		Long:          req.GetLong(),
		VehicleNumber: req.GetVehicleNumber(),
		Status:        req.GetStatus(),
	}
	createdCourier, err := h.App.Create(ctx, courier)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create courier: %v", err)
	}
	return &pb.CourierResponse{
		Courier: &pb.Courier{
			Id:            createdCourier.ID,
			UserId:        createdCourier.UserID,
			Lat:           createdCourier.Lat,
			Long:          createdCourier.Long,
			VehicleNumber: createdCourier.VehicleNumber,
			Status:        createdCourier.Status,
			CreatedAt:     timestamppb.New(createdCourier.CreatedAt),
			UpdatedAt:     timestamppb.New(createdCourier.UpdatedAt),
		},
	}, nil
}
func (h *CourierHandler) GetByLongLatCourier(ctx context.Context, req *pb.GetByLongLatCourierRequest) (*pb.CourierResponse, error) {
	courier, err := h.App.GetByLongLat(ctx, req.GetLat(), req.GetLong())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get courier by location: %v", err)
	}
	if courier == nil {
		return nil, status.Errorf(codes.NotFound, "courier not found at the specified location")
	}
	return &pb.CourierResponse{
		Courier: &pb.Courier{
			Id:            courier.ID,
			UserId:        courier.UserID,
			Lat:           courier.Lat,
			Long:          courier.Long,
			VehicleNumber: courier.VehicleNumber,
			Status:        courier.Status,
			CreatedAt:     timestamppb.New(courier.CreatedAt),
			UpdatedAt:     timestamppb.New(courier.UpdatedAt),
		},
	}, nil
}
func (h *CourierHandler) UpdateLongLatCourier(ctx context.Context, req *pb.UpdateLongLatCourierRequest) (*pb.CourierResponse, error) {
	courier, err := h.App.UpdateLongLat(ctx, req.GetId(), req.GetLat(), req.GetLong())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update courier location: %v", err)
	}
	if courier == nil {
		return nil, status.Errorf(codes.NotFound, "courier not found")
	}
	return &pb.CourierResponse{
		Courier: &pb.Courier{
			Id:            courier.ID,
			UserId:        courier.UserID,
			Lat:           courier.Lat,
			Long:          courier.Long,
			VehicleNumber: courier.VehicleNumber,
			Status:        courier.Status,
			CreatedAt:     timestamppb.New(courier.CreatedAt),
			UpdatedAt:     timestamppb.New(courier.UpdatedAt),
		},
	}, nil
}
func (h *CourierHandler) DeleteCourier(ctx context.Context, req *pb.DeleteCourierRequest) (*pb.DeleteCourierResponse, error) {
	err := h.App.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete courier: %v", err)
	}
	return &pb.DeleteCourierResponse{
		Message: "Courier deleted successfully",
	}, nil
}
func (h *CourierHandler) GetAllCouriers(ctx context.Context, req *pb.Empty) (*pb.CourierListResponse, error) {
	couriers, err := h.App.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all couriers: %v", err)
	}
	var courierList []*pb.Courier
	for _, courier := range couriers {
		courierList = append(courierList, &pb.Courier{
			Id:            courier.ID,
			UserId:        courier.UserID,
			Lat:           courier.Lat,
			Long:          courier.Long,
			VehicleNumber: courier.VehicleNumber,
			Status:        courier.Status,
			CreatedAt:     timestamppb.New(courier.CreatedAt),
			UpdatedAt:     timestamppb.New(courier.UpdatedAt),
		})
	}
	return &pb.CourierListResponse{Couriers: courierList}, nil
}

func (h *CourierHandler) FindNearestCourier(ctx context.Context, req *pb.FindNearestCourierRequest) (*pb.CourierResponse, error) {
	courier, err := h.App.FindNearest(ctx, req.GetLat(), req.GetLong())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find nearest courier: %v", err)
	}
	if courier == nil {
		return nil, status.Errorf(codes.NotFound, "no courier found near the specified location")
	}
	return &pb.CourierResponse{
		Courier: &pb.Courier{
			Id:            courier.ID,
			UserId:        courier.UserID,
			Lat:           courier.Lat,
			Long:          courier.Long,
			VehicleNumber: courier.VehicleNumber,
			Status:        courier.Status,
			CreatedAt:     timestamppb.New(courier.CreatedAt),
			UpdatedAt:     timestamppb.New(courier.UpdatedAt),
		},
	}, nil
}
