package storage

import (
	"context"
	"fmt"
	"sync"

	"github.com/eeQuillibrium/posts/graph/model"
)

const preparedUsersVolume = 10000

var (
	errNoUser = "user with this userID doesn't exist"
)

type usersStorage struct {
	users    map[int]*model.User
	mu       *sync.Mutex
	idSerial int
	mediator Mediator
}

func NewUsersStorage() *usersStorage {
	return &usersStorage{
		users:    make(map[int]*model.User, preparedUsersVolume),
		mu:       &sync.Mutex{},
		idSerial: 0,
	}
}

func (us *usersStorage) CreateUser(
	ctx context.Context,
	user *model.NewUser,
) (int, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	us.idSerial++

	us.users[us.idSerial] = &model.User{
		ID:       us.idSerial,
		Login:    user.Login,
		Password: user.Password,
		Name:     user.Name,
	}

	return us.idSerial, nil
}

func (us *usersStorage) isUserExist(userID int) error {
	us.mu.Lock()
	defer us.mu.Unlock()

	if _, ok := us.users[userID]; !ok {
		return fmt.Errorf("isUserExist(): %w", errNoUser)
	}
	return nil
}

func (us *usersStorage) SetMediator(mediator Mediator) {
	us.mediator = mediator
}
