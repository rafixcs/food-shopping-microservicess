package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleCustomer   Role = "customer"
	RoleRestaurant Role = "restaurant"
	RoleCourier    Role = "courier"
)

func (r Role) Valid() bool {
	switch r {
	case RoleCustomer, RoleRestaurant, RoleCourier:
		return true
	default:
		return false
	}
}

type UserModel struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type AddressModel struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Label     string    `json:"label"`
	Street    string    `json:"street"`
	Number    string    `json:"number"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Zip       string    `json:"zip"`
	Lat       *float64  `json:"lat,omitempty"`
	Lng       *float64  `json:"lng,omitempty"`
	IsDefault bool      `json:"is_default"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
	Phone    string `json:"phone"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, u *UserModel) error
	GetUserByEmail(ctx context.Context, email string) (*UserModel, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*UserModel, error)

	CreateAddress(ctx context.Context, a *AddressModel) error
	GetAddressByID(ctx context.Context, id uuid.UUID) (*AddressModel, error)
	ListAddressesByUser(ctx context.Context, userID uuid.UUID) ([]*AddressModel, error)
}

type UserService interface {
	Register(ctx context.Context, user RegisterInput) (*UserModel, error)
	Login(ctx context.Context, email, password string) (string, error)
}
