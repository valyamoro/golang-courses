package repository

import (
	"errors"
	"sync"
)

type User struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

type Repository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id int) (User, error)
	UpdateUser(id int, updatedUser User) (User, error)
	DeleteUser(id int) error
}

type InMemoryRepository struct {
	users  map[int]User
	mu     sync.Mutex
	nextID int
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (repo *InMemoryRepository) CreateUser(user User) (User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	user.ID = repo.nextID
	repo.NextID++
	repo.users[user.ID] = user
	return user, nil
}

func (repo *InMemoryRepository) GetUserByID(id int) (User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	user, ok := repo.users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}

	return user, nil
}

func (repo *InMemoryRepository) UpdateUser(id int, updatedUser User) (User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, ok := repo.users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}
	updatedUser.ID = id
	repo.users[id] = updatedUser
	return updatedUser, nil
}

func (repo *InMemoryRepository) DeleteUser(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, ok := repo.users[id]; !ok {
		return errors.New("user not found")
	}

	delete(repo.users, id)

	return nil
}

func (repo *InMemoryRepository) DeleteUser(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, ok := repo.users[id]; !ok {
		return errors.New("user not found")
	}

	delete(repo.users, id)

	return nil
}
