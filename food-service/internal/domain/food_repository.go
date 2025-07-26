package domain

import "context"

// type AuthRepository interface {
// 	GetByEmail(ctx context.Context, email string) (*User, error)
// 	Create(ctx context.Context, user *User) (*User, error)
// }

type FoodRepository interface {
	GetById(ctx context.Context, id string) (*Food, error)
	Create(ctx context.Context, food *Food) (*Food, error)
	GetAll(ctx context.Context) ([]*Food, error)
	Update(ctx context.Context, food *Food) (*Food, error)
	Delete(ctx context.Context, id string) error
}

// rpc CreateFood(CreateFoodRequest) returns (FoodResponse);
// rpc GetAllFood(Empty) returns (FoodResponse);
// rpc GetByIdFood(GetByIdFoodRequest) returns (FoodResponse);
// rpc UpdateFood(UpdateFoodRequest) returns (FoodResponse);
// rpc DeleteFood(DeleteFoodRequest) returns (DeleteFoodResponse);
