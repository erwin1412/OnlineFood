package grpc

import (
	"context"
	"fmt"
	"food-service/internal/app"
	pb "food-service/internal/delivery/grpc/pb"
	"food-service/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// type AuthHandler struct {
// 	pb.UnimplementedAuthServiceServer
// 	App *app.AuthApp
// }

// func NewAuthHandler(app *app.AuthApp) *AuthHandler {
// 	return &AuthHandler{App: app}
// }

type FoodHandler struct {
	pb.UnimplementedFoodServiceServer
	App *app.FoodApp
}

func NewFoodHandler(app *app.FoodApp) *FoodHandler {
	return &FoodHandler{App: app}
}

func (h *FoodHandler) GetByIdFood(ctx context.Context, req *pb.GetByIdFoodRequest) (*pb.FoodResponse, error) {
	food, err := h.App.GetById(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get food by id: %v", err)
	}
	if food == nil {
		return nil, status.Errorf(codes.NotFound, "food not found")
	}

	return &pb.FoodResponse{
		Food: &pb.Food{
			Id:           food.ID,
			MerchantId:   food.MerchantID,
			NameFoods:    food.Name,
			Price:        food.Price,
			Availability: food.Availability,
		},
	}, nil
}

// func (h *FoodHandler) CreateFood(ctx context.Context, req *pb.CreateFoodRequest) (*pb.FoodResponse, error) {
// 	food := &domain.Food{
// 		MerchantID:   req.GetMerchantId(),
// 		Name:         req.GetNameFoods(),
// 		Price:        req.GetPrice(),
// 		Availability: req.GetAvailability(),
// 	}

// 	createdFood, err := h.App.Create(ctx, food)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to create food: %v", err)
// 	}

// 	return &pb.FoodResponse{
// 		Food: &pb.Food{
// 			Id:           createdFood.ID,
// 			MerchantId:   createdFood.MerchantID,
// 			NameFoods:    createdFood.Name,
// 			Price:        createdFood.Price,
// 			Availability: createdFood.Availability,
// 			// kalau mau isi created_at, updated_at => mapping ke timestamp di sini
// 		},
// 	}, nil
// }

func (h *FoodHandler) CreateFood(ctx context.Context, req *pb.CreateFoodRequest) (*pb.FoodResponse, error) {
	// Baca user_id dari metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated: metadata not found")
	}
	userIDs := md.Get("user_id")
	if len(userIDs) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated: user_id not found in metadata")
	}

	userID := userIDs[0]
	fmt.Println("UserID:", userID)

	// get merchant ID from user GetMerchantByUserId
	merchant, err := h.App.MerchantClient.GetMerchantByUserId(ctx, userID)
	if err != nil {
		if err == domain.ErrMerchantNotFound {
			return nil, status.Errorf(codes.NotFound, "merchant not found for user ID: %s", userID)
		}
		return nil, status.Errorf(codes.Internal, "failed to get merchant by user ID: %v", err)
	}

	food := &domain.Food{
		MerchantID:   merchant.ID,
		Name:         req.GetNameFoods(),
		Price:        req.GetPrice(),
		Availability: req.GetAvailability(),
	}

	createdFood, err := h.App.Create(ctx, food, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create food: %v", err)
	}

	return &pb.FoodResponse{
		Food: &pb.Food{
			Id:           createdFood.ID,
			MerchantId:   createdFood.MerchantID,
			NameFoods:    createdFood.Name,
			Price:        createdFood.Price,
			Availability: createdFood.Availability,
		},
	}, nil
}

func (h *FoodHandler) GetAllFood(ctx context.Context, req *pb.Empty) (*pb.FoodListResponse, error) {
	foods, err := h.App.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all foods: %v", err)
	}

	var foodList []*pb.Food
	for _, food := range foods {
		foodList = append(foodList, &pb.Food{
			Id:           food.ID,
			MerchantId:   food.MerchantID,
			NameFoods:    food.Name,
			Price:        food.Price,
			Availability: food.Availability,
		})
	}

	return &pb.FoodListResponse{Foods: foodList}, nil
}
func (h *FoodHandler) UpdateFood(ctx context.Context, req *pb.UpdateFoodRequest) (*pb.FoodResponse, error) {
	food := &domain.Food{
		ID:           req.GetId(),
		MerchantID:   req.GetMerchantId(),
		Name:         req.GetNameFoods(),
		Price:        req.GetPrice(),
		Availability: req.GetAvailability(),
	}
	updatedFood, err := h.App.Update(ctx, food)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update food: %v", err)
	}
	return &pb.FoodResponse{
		Food: &pb.Food{
			Id:           updatedFood.ID,
			MerchantId:   updatedFood.MerchantID,
			NameFoods:    updatedFood.Name,
			Price:        updatedFood.Price,
			Availability: updatedFood.Availability,
			// kalau mau isi created_at, updated_at => mapping ke timestamp di sini
		},
	}, nil
}

func (h *FoodHandler) DeleteFood(ctx context.Context, req *pb.DeleteFoodRequest) (*pb.DeleteFoodResponse, error) {
	err := h.App.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete food: %v", err)
	}
	return &pb.DeleteFoodResponse{
		Message: "Food deleted successfully",
	}, nil
}
