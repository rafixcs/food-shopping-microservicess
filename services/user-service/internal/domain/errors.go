package domain

import "errors"

var (
	ErrEmailAlreadyTaken  = errors.New("email já cadastrado")
	ErrUserNotFound       = errors.New("usuário não encontrado")
	ErrAddressNotFound    = errors.New("endereço não encontrado")
	ErrInvalidCredentials = errors.New("credenciais inválidas")
	ErrInvalidRole        = errors.New("role inválida")
)
