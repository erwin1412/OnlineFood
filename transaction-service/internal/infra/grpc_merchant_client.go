package infra

import (
	"context"
	pb "transaction-service/internal/delivery/grpc/pb/merchant"
	"transaction-service/internal/domain"

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

func (m *grpcMerchantClient) GetById(ctx context.Context, id string) (*domain.Merchant, error) {
	res, err := m.client.GetByIdMerchant(ctx, &pb.GetByIdMerchantRequest{Id: id})
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
