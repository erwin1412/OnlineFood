package grpc

import (
	"context"
	"fmt"
	"merchant-service/internal/app"
	pb "merchant-service/internal/delivery/grpc/pb"
	"merchant-service/internal/domain"
	"merchant-service/pkg/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MerchantHandler struct {
	pb.UnimplementedMerchantServiceServer
	App *app.MerchantApp
}

func NewMerchantHandler(app *app.MerchantApp) *MerchantHandler {
	return &MerchantHandler{App: app}
}
func (h *MerchantHandler) GetByIdMerchant(ctx context.Context, req *pb.GetByIdMerchantRequest) (*pb.MerchantResponse, error) {
	merchant, err := h.App.GetById(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get merchant by id: %v", err)
	}
	if merchant == nil {
		return nil, status.Errorf(codes.NotFound, "merchant not found")
	}
	return &pb.MerchantResponse{
		Merchant: &pb.Merchant{
			Id:           merchant.ID,
			UserId:       merchant.UserID,
			NameMerchant: merchant.NameMerchant,
			Alamat:       merchant.Alamat,
			Lat:          merchant.Lat,
			Long:         merchant.Long,
			OpenHour:     merchant.OpenHour,
			CloseHour:    merchant.CloseHour,
			Status:       merchant.Status,
			CreatedAt:    timestamppb.New(merchant.CreatedAt),
			UpdatedAt:    timestamppb.New(merchant.UpdatedAt),
		},
	}, nil
}
func (h *MerchantHandler) CreateMerchant(ctx context.Context, req *pb.CreateMerchantRequest) (*pb.MerchantResponse, error) {
	lat, long, err := utils.GetLatLong(req.GetAlamat())
	if err != nil {
		return nil, fmt.Errorf("failed to geocode: %w", err)
	}

	merchant := &domain.Merchant{
		UserID:       req.GetUserId(),
		NameMerchant: req.GetNameMerchant(),
		Alamat:       req.GetAlamat(),
		Lat:          lat,
		Long:         long,
		OpenHour:     req.GetOpenHour(),
		CloseHour:    req.GetCloseHour(),
		Status:       req.GetStatus(),
	}
	createdMerchant, err := h.App.Create(ctx, merchant)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create merchant: %v", err)
	}
	return &pb.MerchantResponse{
		Merchant: &pb.Merchant{
			Id:           createdMerchant.ID,
			UserId:       createdMerchant.UserID,
			NameMerchant: createdMerchant.NameMerchant,
			Alamat:       createdMerchant.Alamat,
			Lat:          createdMerchant.Lat,
			Long:         createdMerchant.Long,
			OpenHour:     createdMerchant.OpenHour,
			CloseHour:    createdMerchant.CloseHour,
			Status:       createdMerchant.Status,
			CreatedAt:    timestamppb.New(createdMerchant.CreatedAt),
			UpdatedAt:    timestamppb.New(createdMerchant.UpdatedAt),
		},
	}, nil
}
func (h *MerchantHandler) GetAllMerchant(ctx context.Context, req *pb.Empty) (*pb.MerchantListResponse, error) {
	merchants, err := h.App.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all merchants: %v", err)
	}
	var merchantList []*pb.Merchant
	for _, merchant := range merchants {
		merchantList = append(merchantList, &pb.Merchant{
			Id:           merchant.ID,
			UserId:       merchant.UserID,
			NameMerchant: merchant.NameMerchant,
			Alamat:       merchant.Alamat,
			Lat:          merchant.Lat,
			Long:         merchant.Long,
			OpenHour:     merchant.OpenHour,
			CloseHour:    merchant.CloseHour,
			Status:       merchant.Status,
			CreatedAt:    timestamppb.New(merchant.CreatedAt),
			UpdatedAt:    timestamppb.New(merchant.UpdatedAt),
		})
	}
	return &pb.MerchantListResponse{Merchants: merchantList}, nil
}

func (h *MerchantHandler) UpdateMerchant(ctx context.Context, req *pb.UpdateMerchantRequest) (*pb.MerchantResponse, error) {
	lat, long, err := utils.GetLatLong(req.GetAlamat())
	if err != nil {
		return nil, fmt.Errorf("failed to geocode: %w", err)
	}

	merchant := &domain.Merchant{
		ID:           req.GetId(),
		UserID:       req.GetUserId(),
		NameMerchant: req.GetNameMerchant(),
		Alamat:       req.GetAlamat(),
		Lat:          lat,
		Long:         long,
		OpenHour:     req.GetOpenHour(),
		CloseHour:    req.GetCloseHour(),
		Status:       req.GetStatus(),
	}
	updatedMerchant, err := h.App.Update(ctx, merchant)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update merchant: %v", err)
	}
	return &pb.MerchantResponse{
		Merchant: &pb.Merchant{
			Id:           updatedMerchant.ID,
			UserId:       updatedMerchant.UserID,
			NameMerchant: updatedMerchant.NameMerchant,
			Alamat:       updatedMerchant.Alamat,
			Lat:          updatedMerchant.Lat,
			Long:         updatedMerchant.Long,
			OpenHour:     updatedMerchant.OpenHour,
			CloseHour:    updatedMerchant.CloseHour,
			Status:       updatedMerchant.Status,
			CreatedAt:    timestamppb.New(updatedMerchant.CreatedAt),
			UpdatedAt:    timestamppb.New(updatedMerchant.UpdatedAt),
		},
	}, nil
}
func (h *MerchantHandler) DeleteMerchant(ctx context.Context, req *pb.DeleteMerchantRequest) (*pb.DeleteMerchantResponse, error) {
	err := h.App.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete merchant: %v", err)
	}
	return &pb.DeleteMerchantResponse{
		Message: "Merchant deleted successfully",
	}, nil
}

func (h *MerchantHandler) GetMerchantByUserId(ctx context.Context, req *pb.GetMerchantByUserIdRequest) (*pb.GetMerchantByUserIdResponse, error) {

	merchant, err := h.App.GetMerchantByUserId(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get merchant by user id: %v", err)
	}
	if merchant == nil {
		return nil, status.Errorf(codes.NotFound, "merchant not found for user id: %s", req.GetUserId())
	}
	return &pb.GetMerchantByUserIdResponse{
		Merchant: &pb.Merchant{
			Id:           merchant.ID,
			UserId:       merchant.UserID,
			NameMerchant: merchant.NameMerchant,
			Alamat:       merchant.Alamat,
			Lat:          merchant.Lat,
			Long:         merchant.Long,
			OpenHour:     merchant.OpenHour,
			CloseHour:    merchant.CloseHour,
			Status:       merchant.Status,
			CreatedAt:    timestamppb.New(merchant.CreatedAt),
			UpdatedAt:    timestamppb.New(merchant.UpdatedAt),
		},
	}, nil
}
