package action

import (
	"sync"
	"time"

	"github.com/geneseeq/authorize-system/db"
	"github.com/geneseeq/authorize-system/upms/user"
	"gopkg.in/mgo.v2/bson"
)

type roleDBRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *roleDBRepository) Find(id string) (*user.RoleModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.RoleModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, user.ErrUnknown
	}
	return &result, nil
}

func (r *roleDBRepository) FindRoleAll() []*user.RoleModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*user.RoleModel
	con.Find(nil).All(&result)
	return result
}

func (r *roleDBRepository) Store(g *user.RoleModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(g)
	return err
}

func (r *roleDBRepository) Remove(id string) error {
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
func newRoleModel(new *user.RoleModel, result user.RoleModel) user.RoleModel {
	if new.Type != "" {
		result.Type = new.Type
	}

	if new.Name != "" {
		result.Name = new.Name
	}

	if new.Alias != "" {
		result.Alias = new.Alias
	}

	if new.Buildin != false {
		result.Buildin = new.Buildin
	}

	if new.CreateUserID != "" {
		result.CreateUserID = new.CreateUserID
	}

	if new.Validity != false {
		result.Validity = new.Validity
	}

	if new.UpdateUserID != "" {
		result.UpdateUserID = new.UpdateUserID
	}

	result.UpdateTime = user.TimeUtcToCst(time.Now())
	return result
}

func (r *roleDBRepository) Update(id string, role *user.RoleModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.RoleModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	result = newRoleModel(role, result)
	err = con.Update(bson.M{"id": id}, result)
	return err
}

// NewRoleDBRepository returns a new instance of a in-memory cargo repository.
func NewRoleDBRepository(db string, collection string) user.RoleRepository {
	return &roleDBRepository{
		db:         db,
		collection: collection,
	}
}
