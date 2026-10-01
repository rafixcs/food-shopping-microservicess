package repository

import (
	"context"
	"uuid"

	"github.com/rafixcs/food-shopping-microservicess/services/user-service/internal/domain"
)

type InmemRepository struct {
	Users     map[string]*domain.UserModel
	Addresses map[string]*domain.AddressModel
}

func (i *InmemRepository) CreateUser(ctx context.Context, u *domain.UserModel) error {

}
func (i *InmemRepository) GetUserByEmail(ctx context.Context, email string) (*domain.UserModel, error) {

}
func (i *InmemRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.UserModel, error) {

}

func (i *InmemRepository) CreateAddress(ctx context.Context, a *domain.AddressModel) error {

}
func (i *InmemRepository) GetAddressByID(ctx context.Context, id uuid.UUID) (*domain.AddressModel, error) {
}
func (i *InmemRepository) ListAddressesByUser(ctx context.Context, userID uuid.UUID) ([]*domain.AddressModel, error) {
}
