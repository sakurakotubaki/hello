package ports

import "hello/internal/core/domain"

type UserService interface {
	CreateUser(name, email string) (*domain.User, error)
	GetUser(id string) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
	UpdateUser(id, name, email string) (*domain.User, error)
	DeleteUser(id string) error
}
