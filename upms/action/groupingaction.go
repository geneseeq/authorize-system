package action

import (
	"sync"
	"time"

	"github.com/geneseeq/authorize-system/db"
	"github.com/geneseeq/authorize-system/upms/user"
	"gopkg.in/mgo.v2/bson"
)

type groupDBRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *groupDBRepository) Find(id string) (*user.GroupModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.GroupModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, user.ErrUnknown
	}
	return &result, nil
}

func (r *groupDBRepository) FindGroupAll() []*user.GroupModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*user.GroupModel
	con.Find(nil).All(&result)
	return result
}

func (r *groupDBRepository) Store(g *user.GroupModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(g)
	return err
}

func (r *groupDBRepository) Remove(id string) error {
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

func newGroupModel(new *user.GroupModel, result user.GroupModel) user.GroupModel {
	if new.Type != 0 {
		result.Type = new.Type
	}

	if new.Parent != "" {
		result.Parent = new.Parent
	}

	if new.Name != "" {
		result.Name = new.Name
	}

	if new.Code != "" {
		result.Code = new.Code
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
	result.UpdateTime = user.TimeUtcToCst(time.Now())
	return result
}

func (r *groupDBRepository) Update(id string, g *user.GroupModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.GroupModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	result = newGroupModel(g, result)
	err = con.Update(bson.M{"id": id}, bson.M{"$set": result})
	return err
}

// NewGroupDBRepository returns a new instance of a in-memory cargo repository.
func NewGroupDBRepository(db string, collection string) user.GroupRepository {
	return &groupDBRepository{
		db:         db,
		collection: collection,
	}
}
