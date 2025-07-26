package grpc

import (
	"auth-service/internal/app"

	pb "auth-service/internal/delivery/grpc/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	App *app.AuthApp
}

func NewAuthHandler(app *app.AuthApp) *AuthHandler {
	return &AuthHandler{App: app}
}

func (h *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	user, err := h.App.Register(
		ctx,
		req.GetName(),
		req.GetEmail(),
		req.GetPassword(),
		req.GetRole(),
		req.GetPhone(),
		req.GetAlamat(),
		req.GetLat(),
		req.GetLong(),
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to register: %v", err)
	}

	return &pb.AuthResponse{
		Id:     user.ID,
		Role:   user.Role,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
		Alamat: user.Address,
		Lat:    user.Latitude,
		Long:   user.Longitude,
	}, nil
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	token, err := h.App.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "login failed: %v", err)
	}

	// Dapatkan user untuk di-include di response
	user, _ := h.App.UserRepo.GetByEmail(ctx, req.GetEmail())

	resp := &pb.AuthResponse{
		Id:     user.ID,
		Role:   user.Role,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
		Alamat: user.Address,
		Lat:    user.Latitude,
		Long:   user.Longitude,
		Token:  token,
	}

	return resp, nil
}
