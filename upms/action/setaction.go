package action

import (
	"sync"

	"github.com/geneseeq/authorize-system/upms/user"
	"github.com/geneseeq/authorize-system/db"
	"gopkg.in/mgo.v2/bson"
)

type setDBRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *setDBRepository) FindDataSet(id string) (*user.DataSetModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.DataSetModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, user.ErrUnknown
	}
	return &result, nil
}

func (r *setDBRepository) FindAllDataSet() []*user.DataSetModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*user.DataSetModel
	con.Find(nil).All(&result)
	return result
}

func (r *setDBRepository) Store(s *user.DataSetModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(s)
	return err
}

func (r *setDBRepository) Remove(id string) error {
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

func (r *setDBRepository) Update(id string, s *user.DataSetModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Update(bson.M{"id": id}, s)
	return err
}

// NewSetDBRepository returns a new instance of a in-memory cargo repository.
func NewSetDBRepository(db string, collection string) user.DataSetRepository {
	return &setDBRepository{
		db:         db,
		collection: collection,
	}
}
