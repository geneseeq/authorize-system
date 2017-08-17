// Package dataing provides the use-case of booking a cargo. Used by views
// facing an administrator.
package dataing

import (
	"errors"
	"time"

	"github.com/geneseeq/authorize-system/upms/user"
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
	GetDataSet(id string) (DataSet, error)
	PostDataSet(user []DataSet) ([]string, []string, error)
	GetAllDataSet() ([]DataSet, error)
	DeleteMultiDataSet(listid []string) ([]string, []string, error)
	PutMultiDataSet(set []DataSet) ([]string, []string, error)
}

// DataSet is a user base info
type DataSet struct {
	ID           string            `json:"id"`
	Rule         string            `json:"rule"`
	Name         string            `json:"name"`
	MatchField   []user.MatchField `json:"match_field"`
	Type         string            `json:"type"`
	Validity     bool              `json:"validity"`
	Buildin      bool              `json:"buildin"`
	CreateUserID string            `json:"create_user_id"`
	CreateTime   time.Time         `json:"create_time"`
}

type service struct {
	sets user.DataSetRepository
}

// NewService creates a booking service with necessary dependencies.
func NewService(sets user.DataSetRepository) Service {
	return &service{
		sets: sets,
	}
}

func (s *service) PostDataSet(set []DataSet) ([]string, []string, error) {
	var sucessedID []string
	var failedID []string
	if len(set) < LimitMaxSum {
		for _, data := range set {
			data.CreateTime = user.TimeUtcToCst(time.Now())
			err := s.sets.Store(setToSetModel(data))
			if err != nil {
				failedID = append(failedID, data.ID)
				continue
			}
			sucessedID = append(sucessedID, data.ID)
		}
		return sucessedID, failedID, nil
	}
	return sucessedID, failedID, ErrExceededMount
}

func (s *service) GetDataSet(id string) (DataSet, error) {
	if id == "" {
		return DataSet{}, ErrInvalidArgument
	}
	data, error := s.sets.FindDataSet(id)
	if error != nil {
		return DataSet{}, ErrNotFound
	}
	return setModelToSet(data), nil
}

func (s *service) GetAllDataSet() ([]DataSet, error) {
	var result []DataSet
	for _, data := range s.sets.FindAllDataSet() {
		result = append(result, setModelToSet(data))
	}
	return result, nil
}

func (s *service) PutMultiDataSet(set []DataSet) ([]string, []string, error) {
	var sucessedID []string
	var failedID []string
	if len(set) < LimitMaxSum {
		for _, data := range set {
			if len(data.ID) == 0 {
				failedID = append(failedID, data.ID)
				continue
			}
			_, err := s.GetDataSet(data.ID)
			if err != nil {
				failedID = append(failedID, data.ID)
				continue
			}
			err = s.sets.Update(data.ID, setToSetModel(data))
			if err != nil {
				failedID = append(failedID, data.ID)
				continue
			}
			sucessedID = append(sucessedID, data.ID)
		}
		return sucessedID, failedID, nil
	}
	return sucessedID, failedID, ErrExceededMount
}

func (s *service) DeleteMultiDataSet(listid []string) ([]string, []string, error) {
	var sucessedID []string
	var failedID []string
	if len(listid) < LimitMaxSum {
		for _, id := range listid {
			error := s.sets.Remove(id)
			if error != nil {
				failedID = append(failedID, id)
				continue
			}
			sucessedID = append(sucessedID, id)
		}
		return sucessedID, failedID, nil
	}
	return sucessedID, failedID, ErrExceededMount
}

func setToSetModel(s DataSet) *user.DataSetModel {

	return &user.DataSetModel{
		ID:           s.ID,
		Rule:         s.Rule,
		Name:         s.Name,
		MatchField:   s.MatchField,
		Type:         s.Type,
		Validity:     s.Validity,
		Buildin:      s.Buildin,
		CreateUserID: s.CreateUserID,
		CreateTime:   s.CreateTime,
	}
}

func setModelToSet(s *user.DataSetModel) DataSet {
	return DataSet{
		ID:           s.ID,
		Rule:         s.Rule,
		Name:         s.Name,
		MatchField:   s.MatchField,
		Type:         s.Type,
		Validity:     s.Validity,
		Buildin:      s.Buildin,
		CreateUserID: s.CreateUserID,
		CreateTime:   s.CreateTime,
	}
}
