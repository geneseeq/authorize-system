package service

import (
	"sync"

	"github.com/geneseeq/authorize-system/db"
	"github.com/geneseeq/authorize-system/task/baseJob/model"
	"gopkg.in/mgo.v2/bson"
)

type baseDBRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *baseDBRepository) Distinct(id string, condition bson.M) ([]string, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []string
	err := con.Find(condition).Distinct(id, &result)
	if err != nil {
		return nil, model.ErrUnknown
	}
	return result, nil
}

func (r *baseDBRepository) Aggregate(pipeline *[]bson.M) ([]model.BaseInfoModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	var results []model.BaseInfoModel
	con := ds.GetConnect(r.db, r.collection)
	con.Pipe(pipeline).All(&results)
	return results, nil
}

// NewBaseDBRepository returns a new instance of a in-memory cargo repository.
func NewBaseDBRepository(db string, collection string) model.BaseRepository {
	return &baseDBRepository{
		db:         db,
		collection: collection,
	}
}
