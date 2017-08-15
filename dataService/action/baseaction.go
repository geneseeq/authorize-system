package action

import (
	"sync"
	"time"

	"github.com/geneseeq/authorize-system/dataService/data"
	"github.com/geneseeq/authorize-system/db"
	"gopkg.in/mgo.v2/bson"
)

type baseDataDBRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *baseDataDBRepository) Find(id string) (*data.BaseDataModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := data.BaseDataModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, data.ErrUnknown
	}
	return &result, nil
}

func (r *baseDataDBRepository) FindDataAll() []*data.BaseDataModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*data.BaseDataModel
	con.Find(nil).All(&result)
	return result
}

func (r *baseDataDBRepository) Store(d *data.BaseDataModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(d)
	return err
}

func (r *baseDataDBRepository) Remove(id string) error {
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

func (r *baseDataDBRepository) RemoveLabel(id string, label []string) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	// err := con.Remove(bson.M{"id": id})
	result := data.BaseDataModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return data.ErrUnknown
	}
	var newLabelID []string
	if result.LabelID != nil {
		tmpDict := map[string]string{}
		for _, v := range result.LabelID {
			tmpDict[v] = "true"
		}
		for _, v := range label {
			delete(tmpDict, v)
		}
		for key, _ := range tmpDict {
			newLabelID = append(newLabelID, key)
		}
		result.LabelID = newLabelID
	}
	err = con.Update(bson.M{"id": id}, result)
	return err
}

func newBaseDataModel(new *data.BaseDataModel, result data.BaseDataModel) data.BaseDataModel {
	if new.Doctor != "" {
		result.Doctor = new.Doctor
	}

	if new.Hospital != "" {
		result.Hospital = new.Hospital
	}

	if new.HospitalDept != "" {
		result.HospitalDept = new.HospitalDept
	}

	if new.OrderID != "" {
		result.OrderID = new.OrderID
	}

	if new.SaleID != "" {
		result.SaleID = new.SaleID
	}

	if new.SampleID != "" {
		result.SampleID = new.SampleID
	}

	if new.School != "" {
		result.School = new.School
	}

	if new.SchoolDept != "" {
		result.SchoolDept = new.SchoolDept
	}

	if new.Product != "" {
		result.Product = new.Product
	}

	if new.Project != "" {
		result.Project = new.Project
	}
	result.UpdateTime = data.TimeUtcToCst(time.Now())
	return result
}

func (r *baseDataDBRepository) Update(id string, d *data.BaseDataModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := data.BaseDataModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return data.ErrUnknown
	}
	result = newBaseDataModel(d, result)
	var newLabelID []string
	if result.LabelID != nil {
		result.LabelID = append(result.LabelID, d.LabelID...)
		tmpDict := map[string]string{}
		for _, v := range result.LabelID {
			if v != "" {
				tmpDict[v] = "true"
			}
		}
		for key, _ := range tmpDict {
			newLabelID = append(newLabelID, key)
		}
		result.LabelID = newLabelID
	}
	err = con.Update(bson.M{"id": id}, result)
	return err
}

// NewBaseDataDBRepository returns a new instance of a in-memory cargo repository.
func NewBaseDataDBRepository(db string, collection string) data.DataRepository {
	return &baseDataDBRepository{
		db:         db,
		collection: collection,
	}
}
