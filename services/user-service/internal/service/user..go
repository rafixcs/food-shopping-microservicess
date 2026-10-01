package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rafixcs/food-shopping-microservicess/services/user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo   domain.UserRepository
	tokens *TokenManager
}

func NewUserService(repo domain.UserRepository, tokens *TokenManager) *UserService {
	return &UserService{
		repo:   repo,
		tokens: tokens,
	}
}

func (s *UserService) Register(ctx context.Context, in domain.RegisterInput) (*domain.UserModel, error) {
	if !in.Role.Valid() {
		return nil, domain.ErrInvalidRole
	}

	_, err := s.repo.GetUserByEmail(ctx, in.Email)
	if err == nil {
		return nil, domain.ErrEmailAlreadyTaken
	}
	if !errors.Is(err, domain.ErrUserNotFound) {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &domain.UserModel{
		ID:       uuid.New(),
		Name:     in.Name,
		Email:    in.Email,
		Password: string(hash),
		Role:     string(in.Role),
		Phone:    in.Phone,
	}

	if err := s.repo.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		// não vaza se o problema foi email ou senha -- sempre a mesma mensagem
		return "", domain.ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return "", domain.ErrInvalidCredentials
	}
	return s.tokens.Generate(u.ID, string(u.Role))
}

func (s *UserService) ValidateToken(ctx context.Context, token string) (*Claims, error) {
	return s.tokens.Validate(token)
}

func (s *UserService) GetAddress(ctx context.Context, id uuid.UUID) (*domain.AddressModel, error) {
	return s.repo.GetAddressByID(ctx, id)
}
