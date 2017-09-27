// Package action provides in-memory implementations of all the domain repositories.
package action

import (
	"sync"
	"time"

	"github.com/geneseeq/authorize-system/db"
	"github.com/geneseeq/authorize-system/upms/user"
	"gopkg.in/mgo.v2/bson"
)

type serviceDBRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *serviceDBRepository) Store(s *user.ServicesModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(s)
	return err
}

func (r *serviceDBRepository) FindService(id string) (*user.ServicesModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.ServicesModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, user.ErrUnknown
	}
	return &result, nil
}

func (r *serviceDBRepository) FindAllService() []*user.ServicesModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*user.ServicesModel
	con.Find(nil).All(&result)
	return result
}

func (r *serviceDBRepository) Remove(id string) error {
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

func newServicesModel(new *user.ServicesModel, result user.ServicesModel) user.ServicesModel {
	if new.Parent != "" {
		result.Parent = new.Parent
	}

	if len(new.Depend) != 0 {
		result.Depend = new.Depend
	}

	if len(new.Owner) != 0 {
		result.Owner = new.Owner
	}

	if new.Name != "" {
		result.Name = new.Name
	}

	if new.Level != "" {
		result.Level = new.Level
	}

	if new.Path != "" {
		result.Path = new.Path
	}

	if new.Status != "" {
		result.Status = new.Status
	}

	if new.Validity != false {
		result.Validity = new.Validity
	}

	if new.Buildin != false {
		result.Buildin = new.Buildin
	}

	if new.CreateUserID != "" {
		result.CreateUserID = new.CreateUserID
	}
	result.RegisterTime = user.TimeUtcToCst(time.Now())
	return result
}

func (r *serviceDBRepository) Update(id string, s *user.ServicesModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.ServicesModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	result = newServicesModel(s, result)
	err = con.Update(bson.M{"id": id}, bson.M{"$set": result})
	return err
}

// NewServiceDBRepository returns a new instance of a in-memory cargo repository.
func NewServiceDBRepository(db string, collection string) user.ServiceRepository {
	return &serviceDBRepository{
		db:         db,
		collection: collection,
	}
}
