// Package baseing provides the use-case of booking a cargo. Used by views
// facing an administrator.
package baseing

import (
	"errors"
	"time"

	"github.com/geneseeq/authorize-system/dataService/data"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
	ErrExceededMount   = errors.New("exceeded max mount")
	LimitMaxSum        = 50
)

// Service is the interface that provides booking methods.
type Service interface {
	GetBaseData(id string) (BaseData, error)
	PostBaseData(data []BaseData) ([]string, []string, error)
	GetAllBaseData() ([]BaseData, error)
	DeleteMutliLabel(labelid []data.LabelIDModel) ([]string, []string, error)
	DeleteMultiBaseData(listid []string) ([]string, []string, error)
	PutMultiBaseData(data []BaseData) ([]string, []string, error)
}

// BaseData is a basedata base info
type BaseData struct {
	ID           string    `json:"id"`
	SampleID     string    `json:"sample_id"`
	OrderID      string    `json:"order_id"`
	SaleID       string    `json:"sale_id"`
	Doctor       string    `json:"doctor"`
	Hospital     string    `json:"hospital"`
	HospitalDept string    `json:"hospital_dept"`
	School       string    `json:"school"`
	SchoolDept   string    `json:"school_dept"`
	Product      string    `json:"product"`
	Project      string    `json:"project"`
	LabelID      []string  `json:"label_id"`
	UpdateTime   time.Time `json:"update_time"`
	CreateTime   time.Time `json:"create_time"`
}

type service struct {
	datas data.DataRepository
}

// NewService creates a booking service with necessary dependencies.
func NewService(datas data.DataRepository) Service {
	return &service{
		datas: datas,
	}
}

func (s *service) PostBaseData(d []BaseData) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(d) < LimitMaxSum {
		for _, content := range d {
			curTime := data.TimeUtcToCst(time.Now())
			content.CreateTime = curTime
			content.UpdateTime = curTime
			err := s.datas.Store(baseDataToBaseDataModel(content))
			if err != nil {
				failed = append(failed, content.ID)
				continue
			}
			sucessed = append(sucessed, content.ID)
		}
		return sucessed, failed, nil
	}
	return sucessed, failed, ErrExceededMount
}

func (s *service) GetBaseData(id string) (BaseData, error) {
	if id == "" {
		return BaseData{}, ErrInvalidArgument
	}
	d, error := s.datas.Find(id)
	if error != nil {
		return BaseData{}, ErrNotFound
	}
	return baseDataModellToBaseData(d), nil
}

func (s *service) GetAllBaseData() ([]BaseData, error) {
	var result []BaseData
	for _, d := range s.datas.FindDataAll() {
		result = append(result, baseDataModellToBaseData(d))
	}
	return result, nil
}

func (s *service) PutMultiBaseData(d []BaseData) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(d) < LimitMaxSum {
		for _, data := range d {
			if len(data.ID) == 0 {
				failed = append(failed, data.ID)
				continue
			}
			_, err := s.GetBaseData(data.ID)
			if err != nil {
				failed = append(failed, data.ID)
				continue
			}
			err = s.datas.Update(data.ID, baseDataToBaseDataModel(data))
			if err != nil {
				failed = append(failed, data.ID)
				continue
			}
			sucessed = append(sucessed, data.ID)
		}
		return sucessed, failed, nil
	}
	return sucessed, failed, ErrExceededMount
}

func (s *service) DeleteMultiBaseData(listid []string) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(listid) == 0 {
		return sucessed, failed, ErrInvalidArgument
	}
	if len(listid) < LimitMaxSum {
		for _, id := range listid {
			error := s.datas.Remove(id)
			if error != nil {
				failed = append(failed, id)
				continue
			}
			sucessed = append(sucessed, id)
		}
		return sucessed, failed, nil
	}
	return sucessed, failed, ErrExceededMount
}

func (s *service) DeleteMutliLabel(labelid []data.LabelIDModel) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(labelid) < LimitMaxSum {
		for _, label := range labelid {
			error := s.datas.RemoveLabel(label.ID, label.LabelID)
			if error != nil {
				failed = append(failed, label.LabelID...)
				continue
			}
			sucessed = append(sucessed, label.LabelID...)
		}
		return sucessed, failed, nil
	}
	return sucessed, failed, ErrExceededMount
}

func baseDataToBaseDataModel(d BaseData) *data.BaseDataModel {

	return &data.BaseDataModel{
		UnionID:      d.ID,
		ID:           d.ID,
		SampleID:     d.SampleID,
		OrderID:      d.OrderID,
		SaleID:       d.SaleID,
		Doctor:       d.Doctor,
		Hospital:     d.Hospital,
		HospitalDept: d.HospitalDept,
		School:       d.School,
		SchoolDept:   d.SchoolDept,
		Product:      d.Product,
		Project:      d.Project,
		LabelID:      d.LabelID,
		CreateTime:   d.CreateTime,
		UpdateTime:   d.UpdateTime,
	}
}

func baseDataModellToBaseData(d *data.BaseDataModel) BaseData {
	return BaseData{
		ID:           d.ID,
		SampleID:     d.SampleID,
		OrderID:      d.OrderID,
		SaleID:       d.SaleID,
		Doctor:       d.Doctor,
		Hospital:     d.Hospital,
		HospitalDept: d.HospitalDept,
		School:       d.School,
		SchoolDept:   d.SchoolDept,
		Product:      d.Product,
		Project:      d.Project,
		LabelID:      d.LabelID,
		CreateTime:   d.CreateTime,
		UpdateTime:   d.UpdateTime,
	}
}
