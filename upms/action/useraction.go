// Package action provides in-memory implementations of all the domain repositories.
package action

import (
	"sync"
	"time"

	"github.com/geneseeq/authorize-system/db"
	"github.com/geneseeq/authorize-system/upms/user"
	"gopkg.in/mgo.v2/bson"
)

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

func newUserModel(new *user.UserModel, result user.UserModel) user.UserModel {
	if new.Type != 0 {
		result.Type = new.Type
	}

	if new.Number != "" {
		result.Number = new.Number
	}

	if new.Username != "" {
		result.Username = new.Username
	}

	if new.Tele != "" {
		result.Tele = new.Tele
	}

	if new.Gneder != false {
		result.Gneder = new.Gneder
	}

	if new.Status != 0 {
		result.Status = new.Status
	}

	if new.Validity != false {
		result.Validity = new.Validity
	}

	if new.Vip != false {
		result.Vip = new.Vip
	}

	if new.Buildin != false {
		result.Buildin = new.Buildin
	}

	if new.CreateUserID != "" {
		result.CreateUserID = new.CreateUserID
	}
	result.UpdateTime = user.TimeUtcToCst(time.Now())
	return result
}

func (r *userDBRepository) Update(id string, u *user.UserModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.UserModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	result = newUserModel(u, result)
	err = con.Update(bson.M{"id": id}, bson.M{"$set": result})
	return err
}

// NewUserDBRepository returns a new instance of a in-memory cargo repository.
func NewUserDBRepository(db string, collection string) user.Repository {
	return &userDBRepository{
		db:         db,
		collection: collection,
	}
}
