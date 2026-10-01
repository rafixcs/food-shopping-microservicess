package repository

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/rafixcs/food-shopping-microservicess/services/user-service/internal/domain"
)

type InmemRepository struct {
	mu        sync.RWMutex
	Users     map[string]*domain.UserModel
	Addresses map[string]*domain.AddressModel
}

func NewInmemRepository() *InmemRepository {
	return &InmemRepository{
		Users:     make(map[string]*domain.UserModel),
		Addresses: make(map[string]*domain.AddressModel),
	}
}

func (i *InmemRepository) CreateUser(ctx context.Context, u *domain.UserModel) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	val, exists := i.Users[u.ID.String()]
	if exists {
		if val.Email == u.Email {
			return domain.ErrEmailAlreadyTaken
		}
	}

	i.Users[u.ID.String()] = u
	return nil
}

func (i *InmemRepository) GetUserByEmail(ctx context.Context, email string) (*domain.UserModel, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	for _, u := range i.Users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

func (i *InmemRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.UserModel, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	u, ok := i.Users[id.String()]
	if !ok {
		return nil, domain.ErrUserNotFound
	}
	return u, nil
}

func (i *InmemRepository) CreateAddress(ctx context.Context, a *domain.AddressModel) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.Addresses[a.ID.String()] = a
	return nil
}

func (i *InmemRepository) GetAddressByID(ctx context.Context, id uuid.UUID) (*domain.AddressModel, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	a, ok := i.Addresses[id.String()]
	if !ok {
		return nil, domain.ErrAddressNotFound
	}
	return a, nil
}

func (i *InmemRepository) ListAddressesByUser(ctx context.Context, userID uuid.UUID) ([]*domain.AddressModel, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	addresses := make([]*domain.AddressModel, 0)
	for _, a := range i.Addresses {
		if a.UserID == userID {
			addresses = append(addresses, a)
		}
	}
	return addresses, nil
}
