package memory

import (
	"errors"
	"hello/internal/core/domain"
	"hello/internal/core/ports"
	"sync"
)

type userRepository struct {
	users map[string]*domain.User
	mutex sync.RWMutex
}

func NewUserRepository() ports.UserRepository {
	return &userRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *userRepository) Create(user *domain.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	r.users[user.ID] = user
	return nil
}

func (r *userRepository) GetByID(id string) (*domain.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *userRepository) GetAll() ([]*domain.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]*domain.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) Update(user *domain.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}

	r.users[user.ID] = user
	return nil
}

func (r *userRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}
