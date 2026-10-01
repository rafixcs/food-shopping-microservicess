package grpc

import (
	"context"

	"github.com/rafixcs/food-shopping-microservicess/services/user-service/internal/domain"
	"github.com/rafixcs/food-shopping-microservicess/services/user-service/internal/service"
	pb "github.com/rafixcs/food-shopping-microservicess/shared/proto"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewGrpcHandler(server *grpc.Server, service *service.UserService) *grpcHandler {
	handler := &grpcHandler{
		service: service,
	}

	pb.RegisterUserServiceServer(server, handler)

	return handler
}

func (h *grpcHandler) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.UserResponse, error) {
	user := domain.RegisterInput{
		Name:     in.Name,
		Email:    in.Email,
		Role:     domain.Role(in.Role),
		Phone:    in.Phone,
		Password: in.Password,
	}

	data, err := h.service.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:    data.ID.String(),
		Name:  data.Name,
		Email: data.Email,
		Role:  data.Role,
	}, nil
}

func (h *grpcHandler) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (h *grpcHandler) ValidateToken(ctx context.Context, in *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	return nil, nil
}

func (h *grpcHandler) GetAddress(ctx context.Context, in *pb.GetAddressRequest) (*pb.AddressResponse, error) {
	return nil, nil
}
