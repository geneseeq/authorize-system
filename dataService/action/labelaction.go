package action

import (
	"sync"

	"github.com/geneseeq/authorize-system/dataService/data"
	"github.com/geneseeq/authorize-system/db"
	"gopkg.in/mgo.v2/bson"
)

type labelDBRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *labelDBRepository) Find(id string) (*data.LabelModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := data.LabelModel{}
	err := con.Find(bson.M{"labelid": id}).One(&result)
	if err != nil {
		return nil, data.ErrUnknown
	}
	return &result, nil
}

func (r *labelDBRepository) FindLabelAll() []*data.LabelModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*data.LabelModel
	con.Find(nil).All(&result)
	return result
}

func (r *labelDBRepository) Store(d *data.LabelModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(d)
	return err
}

func (r *labelDBRepository) Remove(id string) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Remove(bson.M{"labelid": id})
	if err != nil {
		return data.ErrUnknown
	}
	return nil
}

func appendListID(destListid []string, srcListID []string) []string {
	var newLabelID []string
	if srcListID != nil {
		srcListID = append(srcListID, destListid...)
		tmpDict := map[string]string{}
		for _, v := range srcListID {
			if v != "" {
				tmpDict[v] = "true"
			}
		}
		for key, _ := range tmpDict {
			newLabelID = append(newLabelID, key)
		}
		return newLabelID
	}
	return srcListID
}

func (r *labelDBRepository) Update(id string, d *data.LabelModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := data.LabelModel{}
	err := con.Find(bson.M{"labelid": id}).One(&result)
	if err != nil {
		return data.ErrUnknown
	}
	result.SampleID = appendListID(d.SampleID, result.SampleID)
	result.OrderID = appendListID(d.OrderID, result.OrderID)
	err = con.Update(bson.M{"labelid": id}, result)
	return err
}

// NewLabelDBRepository returns a new instance of a in-memory cargo repository.
func NewLabelDBRepository(db string, collection string) data.LableRepository {
	return &labelDBRepository{
		db:         db,
		collection: collection,
	}
}
