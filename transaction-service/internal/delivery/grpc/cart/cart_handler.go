package cart

import (
	"context"
	"transaction-service/internal/app"
	pb "transaction-service/internal/delivery/grpc/pb/cart"
	"transaction-service/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CartHandler struct {
	pb.UnimplementedCartServiceServer
	App *app.CartApp
}

func NewCartHandler(app *app.CartApp) *CartHandler {
	return &CartHandler{App: app}
}

// ✅ CreateCart
func (h *CartHandler) CreateCart(ctx context.Context, req *pb.CreateCartRequest) (*pb.CartResponse, error) {
	cart := &domain.Cart{
		MerchantID: req.GetMerchantId(),
		FoodID:     req.GetFoodId(),
		UserID:     req.GetUserId(),
		Qty:        req.GetQty(),
	}

	createdCart, err := h.App.Create(ctx, cart)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create cart: %v", err)
	}

	return &pb.CartResponse{
		Cart: mapToPbCart(createdCart),
	}, nil
}

// ✅ GetAllCart
func (h *CartHandler) GetAllCart(ctx context.Context, req *pb.GetAllCartRequest) (*pb.CartListResponse, error) {
	carts, err := h.App.GetAll(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all carts: %v", err)
	}

	var pbCarts []*pb.Cart
	for _, cart := range carts {
		pbCarts = append(pbCarts, mapToPbCart(cart))
	}

	return &pb.CartListResponse{
		Carts: pbCarts,
	}, nil
}

// ✅ GetByIdCart
func (h *CartHandler) GetByIdCart(ctx context.Context, req *pb.GetByIdCartRequest) (*pb.CartResponse, error) {
	cart, err := h.App.GetById(ctx, req.GetId(), req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart by id: %v", err)
	}
	if cart == nil {
		return nil, status.Errorf(codes.NotFound, "cart not found")
	}

	return &pb.CartResponse{
		Cart: mapToPbCart(cart),
	}, nil
}

// ✅ UpdateCart
func (h *CartHandler) UpdateCart(ctx context.Context, req *pb.UpdateCartRequest) (*pb.CartResponse, error) {
	cart := &domain.Cart{
		ID:         req.GetId(),
		MerchantID: req.GetMerchantId(),
		FoodID:     req.GetFoodId(),
		UserID:     req.GetUserId(),
		Qty:        req.GetQty(),
	}

	updatedCart, err := h.App.Update(ctx, cart)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update cart: %v", err)
	}

	return &pb.CartResponse{
		Cart: mapToPbCart(updatedCart),
	}, nil
}

// ✅ DeleteCart
func (h *CartHandler) DeleteCart(ctx context.Context, req *pb.DeleteCartRequest) (*pb.DeleteCartResponse, error) {
	err := h.App.Delete(ctx, req.GetId(), req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete cart: %v", err)
	}

	return &pb.DeleteCartResponse{
		Message: "Cart deleted successfully",
	}, nil
}

// ✅ Helper mapper
func mapToPbCart(cart *domain.Cart) *pb.Cart {
	return &pb.Cart{
		Id:         cart.ID,
		MerchantId: cart.MerchantID,
		FoodId:     cart.FoodID,
		UserId:     cart.UserID,
		Qty:        cart.Qty,
		CreatedAt:  timestamppb.New(cart.CreatedAt),
		UpdatedAt:  timestamppb.New(cart.UpdatedAt),
	}
}

func (h *CartHandler) DeleteAllCart(ctx context.Context, req *pb.DeleteAllCartRequest) (*pb.DeleteAllCartResponse, error) {
	err := h.App.DeleteAll(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete all carts: %v", err)
	}

	return &pb.DeleteAllCartResponse{
		Message: "All carts deleted successfully",
	}, nil
}
