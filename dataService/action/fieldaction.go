package action

import (
	"sync"
	"time"

	"github.com/geneseeq/authorize-system/dataService/data"
	"github.com/geneseeq/authorize-system/db"
	"gopkg.in/mgo.v2/bson"
)

type fieldDBRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *fieldDBRepository) Find(id string) (*data.FieldModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := data.FieldModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, data.ErrUnknown
	}
	return &result, nil
}

func (r *fieldDBRepository) FindFieldAll() []*data.FieldModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*data.FieldModel
	con.Find(nil).All(&result)
	return result
}

func (r *fieldDBRepository) Store(d *data.FieldModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(d)
	return err
}

func (r *fieldDBRepository) Remove(id string) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Remove(bson.M{"id": id})
	if err != nil {
		return data.ErrUnknown
	}
	return nil
}
func newBaseLabelModel(new *data.FieldModel, result data.FieldModel) data.FieldModel {
	if new.Field != "" {
		result.Field = new.Field
	}

	if new.Type != "" {
		result.Type = new.Type
	}

	if new.Comment != "" {
		result.Comment = new.Comment
	}
	result.UpdateTime = data.TimeUtcToCst(time.Now())
	return result
}

func (r *fieldDBRepository) Update(id string, d *data.FieldModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := data.FieldModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return data.ErrUnknown
	}
	result = newBaseLabelModel(d, result)
	err = con.Update(bson.M{"id": id}, bson.M{"$set": result})
	return err
}

// NewFieldDBRepository returns a new instance of a in-memory cargo repository.
func NewFieldDBRepository(db string, collection string) data.FieldRepository {
	return &fieldDBRepository{
		db:         db,
		collection: collection,
	}
}
