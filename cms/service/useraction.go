// Package action provides in-memory implementations of all the domain repositories.
package action

import (
	"sync"
)

type userRepository struct {
	mtx   sync.RWMutex
	users map[user.TrackingID]*user.UserModel
}

func (r *userRepository) Store(c *users.UserModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.users[c.TrackingID] = c
	return nil
}

func (r *userRepository) Find(id users.TrackingID) (*users.UserModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.users[id]; ok {
		return val, nil
	}
	return nil, users.ErrUnknown
}

func (r *userRepository) FindAll() []*users.UserModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	c := make([]*users.UserModel, 0, len(r.users))
	for _, val := range r.users {
		c = append(c, val)
	}
	return c
}

// NewUserRepository returns a new instance of a in-memory cargo repository.
func NewUserRepository() user.Repository {
	return &userRepository{
		users: make(map[user.TrackingID]*user.UserModel),
	}
}
