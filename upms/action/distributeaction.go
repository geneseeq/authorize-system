package action

import (
	"sync"
	"time"

	"github.com/geneseeq/authorize-system/db"
	"github.com/geneseeq/authorize-system/upms/user"
	"gopkg.in/mgo.v2/bson"
)

type distributeDBRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *distributeDBRepository) Find(id string) (*user.RoleDistributeModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.RoleDistributeModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, user.ErrUnknown
	}
	return &result, nil
}

func (r *distributeDBRepository) FindRoleDistributeAll() []*user.RoleDistributeModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*user.RoleDistributeModel
	con.Find(nil).All(&result)
	return result
}

func (r *distributeDBRepository) Store(g *user.RoleDistributeModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(g)
	return err
}

func (r *distributeDBRepository) Remove(id string) error {
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

func newRoleDistributeModel(new *user.RoleDistributeModel, result user.RoleDistributeModel) user.RoleDistributeModel {
	if new.GroupID != "" {
		result.GroupID = new.GroupID
	}

	if new.UserID != "" {
		result.UserID = new.UserID
	}

	if new.RoleID != "" {
		result.RoleID = new.RoleID
	}

	if new.Buildin != false {
		result.Buildin = new.Buildin
	}

	if new.CreateUserID != "" {
		result.CreateUserID = new.CreateUserID
	}

	if new.UpdateUserID != "" {
		result.UpdateUserID = new.UpdateUserID
	}

	result.UpdateTime = user.TimeUtcToCst(time.Now())
	return result
}

func (r *distributeDBRepository) Update(id string, g *user.RoleDistributeModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.RoleDistributeModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	result = newRoleDistributeModel(g, result)
	err = con.Update(bson.M{"id": id}, bson.M{"$set": result})
	return err
}

// NewnewRoleDistributeDBRepository returns a new instance of a in-memory cargo repository.
func NewnewRoleDistributeDBRepository(db string, collection string) user.RoleDistributeRepository {
	return &distributeDBRepository{
		db:         db,
		collection: collection,
	}
}
