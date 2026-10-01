package domain

import "errors"

var (
	ErrEmailAlreadyTaken  = errors.New("email já cadastrado")
	ErrUserNotFound       = errors.New("usuário não encontrado")
	ErrAddressNotFound    = errors.New("endereço não encontrado")
	ErrInvalidCredentials = errors.New("credenciais inválidas")
	ErrAlreadyCreatedUser = errors.New("user ja criado")
	ErrInvalidRole        = errors.New("role inválida")
)
