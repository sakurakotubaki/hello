package repositories

import (
	"hello/internal/core/domain"
	"sync"
	"time"

	"github.com/google/uuid"
)

type InMemoryUserRepository struct {
	users map[string]*domain.User
	mutex sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *InMemoryUserRepository) Create(user *domain.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetByID(id string) (domain.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if user, exists := r.users[id]; exists {
		return *user, nil
	}
	return domain.User{}, nil
}

func (r *InMemoryUserRepository) GetAll() ([]*domain.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]*domain.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

func (r *InMemoryUserRepository) Update(user *domain.User) (domain.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return domain.User{}, nil
	}

	user.UpdatedAt = time.Now()
	r.users[user.ID] = user
	return *user, nil
}

func (r *InMemoryUserRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	delete(r.users, id)
	return nil
}
