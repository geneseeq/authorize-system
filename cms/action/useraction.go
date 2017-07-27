// Package action provides in-memory implementations of all the domain repositories.
package action

import (
	"sync"

	"github.com/geneseeq/authorize-system/cms/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

type userDBRepository struct {
	mtx  sync.RWMutex
	coll *mgo.Collection
}

func (r *userDBRepository) Store(c *user.UserModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	err := r.coll.Insert(c)
	return err
}

func (r *userDBRepository) Find(id string) (*user.UserModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	result := user.UserModel{}
	err := r.coll.Find(bson.M{"ID": id}).One(&result)
	if err != nil {
		return &result, err
	}
	return nil, user.ErrUnknown
}

func (r *userDBRepository) FindAll() []*user.UserModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return nil
}

// NewUserDBRepository returns a new instance of a in-memory cargo repository.
func NewUserDBRepository(collection mgo.Collection) user.Repository {
	return &userDBRepository{
		coll: &collection,
	}
}
