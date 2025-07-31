package infra

import (
	"context"
	pb "food-service/internal/delivery/grpc/pb/merchant"
	"food-service/internal/domain"

	// pb "courier-service/internal/delivery/grpc/pb/merchant"
	// "courier-service/internal/domain"

	"google.golang.org/grpc"
)

type grpcMerchantClient struct {
	client pb.MerchantServiceClient
}

func NewMerchantClient(conn *grpc.ClientConn) domain.MerchantClient {
	return &grpcMerchantClient{
		client: pb.NewMerchantServiceClient(conn),
	}
}

// GetMerchantByUserId
func (m *grpcMerchantClient) GetMerchantByUserId(ctx context.Context, userId string) (*domain.Merchant, error) {
	res, err := m.client.GetMerchantByUserId(ctx, &pb.GetMerchantByUserIdRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	merchant := res.Merchant

	return &domain.Merchant{
		ID:           merchant.Id,
		UserID:       merchant.UserId,
		NameMerchant: merchant.NameMerchant,
		Alamat:       merchant.Alamat,
		Lat:          merchant.Lat,
		Long:         merchant.Long,
		OpenHour:     merchant.OpenHour,
		CloseHour:    merchant.CloseHour,
		Status:       merchant.Status,
		CreatedAt:    merchant.CreatedAt.AsTime(),
		UpdatedAt:    merchant.UpdatedAt.AsTime(),
	}, nil
}
