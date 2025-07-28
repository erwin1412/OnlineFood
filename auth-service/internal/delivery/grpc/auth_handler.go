package grpc

import (
	"auth-service/internal/app"
	"auth-service/pkg/utils"
	"fmt"

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
	// 1. Geocode dulu
	lat, lng, err := utils.GetLatLong(req.GetAlamat())
	if err != nil {
		return nil, fmt.Errorf("failed to geocode: %w", err)
	}

	// 2. Simpan user
	user, err := h.App.Register(
		ctx,
		req.GetName(),
		req.GetEmail(),
		req.GetPassword(),
		req.GetRole(),
		req.GetPhone(),
		req.GetAlamat(),
		lat,
		lng,
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to register: %v", err)
	}

	// 3. Response
	return &pb.AuthResponse{
		Id:        user.ID,
		Role:      user.Role,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Alamat:    user.Alamat,
		Latitude:  user.Latitude,
		Longitude: user.Longitude,
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
		Id:        user.ID,
		Role:      user.Role,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Alamat:    user.Alamat,
		Latitude:  user.Latitude,
		Longitude: user.Longitude,
		Token:     token,
	}

	return resp, nil
}
