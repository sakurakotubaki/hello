package services

import (
	"hello/internal/core/domain"
	"hello/internal/core/ports"
	"time"

	"github.com/google/uuid"
)

type userService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(name, email string) (*domain.User, error) {
	user := &domain.User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUser(id string) (*domain.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) GetAllUsers() ([]*domain.User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) UpdateUser(id, name, email string) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = name
	user.Email = email
	user.UpdatedAt = time.Now()

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}
