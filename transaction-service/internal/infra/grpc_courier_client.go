// infra/grpc_courier_client.go
package infra

import (
	"context"
	pb "transaction-service/internal/delivery/grpc/pb/courier"
	"transaction-service/internal/domain"

	"google.golang.org/grpc"
)

type grpcCourierClient struct {
	client pb.CourierServiceClient
}

func NewCourierClient(conn *grpc.ClientConn) domain.CourierClient {
	return &grpcCourierClient{
		client: pb.NewCourierServiceClient(conn),
	}
}

func (c *grpcCourierClient) FindNearest(ctx context.Context, lat, long string) (*domain.Courier, error) {
	res, err := c.client.FindNearestCourier(ctx, &pb.FindNearestCourierRequest{
		Lat:  lat,
		Long: long,
	})
	if err != nil {
		return nil, err
	}
	return &domain.Courier{
		ID:            res.Courier.Id,
		UserID:        res.Courier.UserId,
		Lat:           res.Courier.Lat,
		Long:          res.Courier.Long,
		VehicleNumber: res.Courier.VehicleNumber,
		Status:        res.Courier.Status,
		CreatedAt:     res.Courier.CreatedAt.AsTime(),
		UpdatedAt:     res.Courier.UpdatedAt.AsTime(),
	}, nil
}
