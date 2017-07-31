// Package action provides in-memory implementations of all the domain repositories.
package action

import (
	"sync"

	"github.com/geneseeq/authorize-system/cms/user"
	"github.com/geneseeq/authorize-system/db"
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

func (r *userRepository) Remove(id string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
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

func (r *userRepository) Update(id string, u *user.UserModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	return nil
}

// NewUserRepository returns a new instance of a in-memory cargo repository.
func NewUserRepository() user.Repository {
	return &userRepository{
		users: make(map[string]*user.UserModel),
	}
}

type userDBRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *userDBRepository) Store(c *user.UserModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(c)
	return err
}

func (r *userDBRepository) Find(id string) (*user.UserModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.UserModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, user.ErrUnknown
	}
	return &result, nil
}

func (r *userDBRepository) FindAll() []*user.UserModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*user.UserModel
	con.Find(nil).All(&result)
	return result
}

func (r *userDBRepository) Remove(id string) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Remove(bson.M{"id": id})
	if err != nil {
		return user.ErrUnknown
	}
	return nil
}

func (r *userDBRepository) Update(id string, u *user.UserModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Update(bson.M{"id": id}, u)
	return err
}

// NewUserDBRepository returns a new instance of a in-memory cargo repository.
func NewUserDBRepository(db string, collection string) user.Repository {
	return &userDBRepository{
		db:         db,
		collection: collection,
	}
}
