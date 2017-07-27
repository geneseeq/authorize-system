// Package action provides in-memory implementations of all the domain repositories.
package action

import (
	"sync"

	"github.com/geneseeq/authorize-system/cms/user"
)

type userRepository struct {
	mtx   sync.RWMutex
	users map[string]*user.UserModel
}

func (r *userRepository) Store(c *user.UserModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.users[c.ID] = c
	return nil
}

func (r *userRepository) Find(id string) (*user.UserModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.users[id]; ok {
		return val, nil
	}
	return nil, user.ErrUnknown
}

func (r *userRepository) FindAll() []*user.UserModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	c := make([]*user.UserModel, 0, len(r.users))
	for _, val := range r.users {
		c = append(c, val)
	}
	return c
}

// NewUserRepository returns a new instance of a in-memory cargo repository.
func NewUserRepository() user.Repository {
	return &userRepository{
		users: make(map[string]*user.UserModel),
	}
}
